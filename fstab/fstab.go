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
	Options    string
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
		if strings.TrimSpace(line) == "" {
			continue
		}
		entry, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		res = append(res, entry)
	}
	return nil, nil
}

func parseLine(line string) (FsEntry, error) {
	fmt.Println("parse line:", line)
	return FsEntry{}, nil
}
