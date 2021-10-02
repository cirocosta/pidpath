package main

import (
	"os"
	"strconv"
)

func getExePathFromPid(pid int) (path string, err error) {
	path, err = os.Readlink("/proc/" + strconv.Itoa(pid) + "/exe")
	return
}
