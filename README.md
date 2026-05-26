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

## Possible errors

`ReadFile` can return normal file errors, for example when the file does not exist or cannot be read.

It can also return a `LineParsingError` when a line is not a valid fstab line. A valid line needs these 6 fields:

```txt
<device> <mountpoint> <type> <options> <dump> <fsck>
```

Example of a broken line:

```txt
UUID=1234 / ext4 defaults
```
