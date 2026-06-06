package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	path, err := readmePath()
	if err != nil {
		return err
	}
	_, err = loadReadme(path)
	return err
}
