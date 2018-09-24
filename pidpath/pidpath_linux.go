// +build linux
package pidpath

import (
	"os"
	"strconv"
)

func GetExePathFromPid(pid int) (path string, err error) {
	path, err = os.Readlink("/proc/" + strconv.Itoa(pid) + "/exe")
	return
}
