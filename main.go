package main

import (
	"fmt"

	"github.com/loissascha/go-fstab/fstab"
)

func main() {
	entries, err := fstab.ReadFile("/etc/fstab")
	if err != nil {
		panic(err)
	}

	for _, r := range entries {
		fmt.Println("_____ FsEntry ____")
		fmt.Println("Device:", r.Device)
		fmt.Println("Mountpoint:", r.Mountpoint)
		fmt.Println("FsType:", r.FsType)
		fmt.Println("Options:", r.Options)
		fmt.Println("Dump:", r.Dump)
		fmt.Println("Fsck:", r.Fsck)
		fmt.Println("__________________")
	}
}
