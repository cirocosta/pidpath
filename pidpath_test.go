package pidpath_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cirocosta/pidpath"
)

func TestGet(t *testing.T) {
	pid := int32(os.Getpid())
	if pid == 0 {
		t.Fatalf("Failed to retrieve current PID")
	}

	realPath, err := os.Executable()
	if err == nil {
		// os.Executable() may return a symlink'd path.
		realPath, err = filepath.EvalSymlinks(realPath)
	}
	if err != nil {
		t.Fatalf("Failed to retrieve current executable: %v", err)
	}

	gotPath, err := pidpath.Get(pid)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if gotPath != realPath {
		t.Fatalf(
			"Expected no get executable path.\nwant: %s\ngot: %s",
			realPath,
			gotPath,
		)
	}
}
