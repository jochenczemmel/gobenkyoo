//go:build !windows

package cfg

import (
	"os"
	"path/filepath"
)

// determine default base directory on non-windows
var defaultBaseDir = filepath.Join(os.Getenv("HOME"), ".gobenkyoo")
