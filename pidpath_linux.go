package main

import (
	"os"
	"strconv"
)

func get(pid int) (path string, err error) {
	path, err = os.Readlink("/proc/" + strconv.Itoa(pid) + "/exe")
	return
}
