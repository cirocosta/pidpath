package main

// GetExePathFromPid retrieves the absolute path to the executable of a running
// process.
func GetExePathFromPid(pid int) (path string, err error) {
	return getExePathFromPid(pid)
}
