package main

// Get retrieves the absolute path to the executable of a running
// process.
func Get(pid int32) (path string, err error) {
	return get(pid)
}
