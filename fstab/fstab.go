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
	for line := range split {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fmt.Println(line)
	}
	return nil, nil
}
