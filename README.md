# go-fstab

A small Go package for reading Linux `fstab` files.

It reads a file like `/etc/fstab` and gives you a list of entries with the device, mount point, filesystem type, options, dump value, and fsck value.

## Install

```sh
go get github.com/loissascha/go-fstab
```

## Usage

```go
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

	for _, entry := range entries {
		fmt.Println(entry.Device, "is mounted at", entry.Mountpoint)
		fmt.Println("type:", entry.FsType)
		fmt.Println("options:", entry.Options)
	}
}
```

## What you get

Each line becomes an `FsEntry`:

```go
type FsEntry struct {
	Device     string
	Mountpoint string
	FsType     string
	Options    []string
	Dump       string
	Fsck       string
}
```

Comments and empty lines are skipped.
