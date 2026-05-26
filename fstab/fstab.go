package fstab

import (
	"fmt"
	"os"
	"strings"
)

type LineParsingError struct {
	Len  int
	Line string
}

func (e LineParsingError) Error() string {
	return fmt.Sprintf("invalid line len(%d): %s", e.Len, e.Line)
}

var _ error = LineParsingError{}

type FsEntry struct {
	Device     string
	Mountpoint string
	FsType     string
	Options    []string
	Dump       string
	Fsck       string
}

func ReadFile(path string) ([]FsEntry, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseStr(string(f))
}

func AddEntry(path string, entry FsEntry) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	line, err := entryToLine(entry)
	if err != nil {
		return err
	}
	newContent := fmt.Sprintf("%s\n%s", string(f), line)
	return os.WriteFile(path, []byte(newContent), 0644)
}

func parseStr(f string) ([]FsEntry, error) {
	split := strings.SplitSeq(f, "\n")
	res := make([]FsEntry, 0)
	for line := range split {
		line = clearComments(line)
		if strings.TrimSpace(line) == "" {
			continue
		}
		entry, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		res = append(res, entry)
	}

	return res, nil
}

func clearComments(line string) string {
	split := strings.Split(line, "#")
	return split[0]
}

func parseLine(line string) (FsEntry, error) {
	split := strings.Fields(line)
	if len(split) < 6 {
		return FsEntry{}, LineParsingError{
			Line: line,
			Len:  len(line),
		}
	}
	return FsEntry{
		Device:     split[0],
		Mountpoint: split[1],
		FsType:     split[2],
		Options:    strings.Split(split[3], ","),
		Dump:       split[4],
		Fsck:       split[5],
	}, nil
}

func entryToLine(entry FsEntry) (string, error) {
	var sb strings.Builder
	first := true
	for _, o := range entry.Options {
		if !first {
			_, err := sb.WriteString(",")
			if err != nil {
				return "", err
			}
		}
		first = false
		_, err := sb.WriteString(o)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s %s %s %s %s %s", entry.Device, entry.Mountpoint, entry.FsType, sb.String(), entry.Dump, entry.Fsck), nil
}
