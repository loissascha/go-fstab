package main

import (
	"fmt"

	"github.com/loissascha/go-fstab/fstab"
)

func main() {
	fmt.Println("main")

	_, err := fstab.ReadFile("/etc/fstab")
	if err != nil {
		panic(err)
	}
}
