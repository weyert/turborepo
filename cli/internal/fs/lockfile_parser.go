package fs

import (
	"fmt"
	"path/filepath"
)

func parseLockfile(file string) (*YarnLockfile, error) {
	// use the base name of the given file (path)
	filename := filepath.Base(file)

	switch filename {
	case "yarn.lock":
		return parseYarnLockfile(file)
	default:
		return nil, fmt.Errorf("unsupported lock file type: %s", filename)
	}
}
