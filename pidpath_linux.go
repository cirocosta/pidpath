package pidpath

import (
	"fmt"
	"os"
)

func get(pid int32) (path string, err error) {
	return os.Readlink("/proc/" + fmt.Sprint(pid) + "/exe")
}
