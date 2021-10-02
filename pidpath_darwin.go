package pidpath

// #include <libproc.h>
// #include <stdlib.h>
// #include <errno.h>
import "C"

import (
	"fmt"
	"unsafe"
)

const bufSize = C.PROC_PIDPATHINFO_MAXSIZE

func get(pid int32) (path string, err error) {
	buf := C.CString(string(make([]byte, bufSize)))
	defer C.free(unsafe.Pointer(buf))

	ret, err := C.proc_pidpath(C.int(pid), unsafe.Pointer(buf), bufSize)
	if ret <= 0 {
		err = fmt.Errorf("failed to retrieve pid path: %v", err)
		return
	}

	path = C.GoString(buf)
	return
}
