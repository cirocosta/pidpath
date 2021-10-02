package main

import (
	"fmt"
	"os"
	"strconv"
)

const help = `Usage:
  pidpath pid
`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, help)
		fmt.Fprintln(os.Stderr, "error: invalid number of arguments")
		os.Exit(1)
	}

	pidStr := os.Args[1]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: failed to parse pid - %v", err)
		os.Exit(1)
	}

	path, err := Get(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: failed to parse pid - %v", err)
		os.Exit(1)
	}

	fmt.Println(path)
}
