package fstab

import (
	"fmt"
	"os"
	"strings"
)

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
	split := strings.SplitSeq(string(f), "\n")
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
		return FsEntry{}, fmt.Errorf("invalid line len(%d): %s", len(split), line)
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
