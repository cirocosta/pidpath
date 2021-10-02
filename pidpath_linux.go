package main

import (
	"fmt"
	"os"
)

func get(pid int32) (path string, err error) {
	path, err = os.Readlink("/proc/" + fmt.Sprint(pid) + "/exe")
	return
}
