package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cirocosta/pidpath"
)

const help = `Usage:
  pidpath pid
`

func main() {
	if len(os.Args) != 2 {
		fmtFatal("%s\nerror: invalid number of arguments", help)
	}

	pidStr := os.Args[1]
	pid64, err := strconv.ParseInt(pidStr, 10, 32)
	if err != nil {
		fmtFatal("error: failed to parse pid - %v\n", err)
	}
	pid := int32(pid64)

	path, err := pidpath.Get(pid)
	if err != nil {
		fmtFatal("error: failed to parse pid - %v\n", err)
	}

	fmt.Println(path)
}

func fmtFatal(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}
