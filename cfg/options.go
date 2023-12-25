// Package cfg provides static and default configuration values.
package cfg

import (
	"path/filepath"
)

// static defaults
const DefaultDbType = "jsonfile"
const DefaultGuiType = "qt6"

// defaults using user information
var (
	DefaultCfgFile = filepath.Join(defaultBaseDir, "cfg.yaml")
	DefaultDbPath  = filepath.Join(defaultBaseDir, "db")
)
