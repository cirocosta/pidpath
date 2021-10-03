package pidpath

import (
	"github.com/shirou/gopsutil/process"
)

func get(pid int32) (path string, err error) {
	// Until a more performant method has been implemented for Windows, fall back
	// to gopsutil/process.
	nps, err := process.NewProcess(pid)
	if err != nil {
		return
	}
	return nps.Exe()
}
