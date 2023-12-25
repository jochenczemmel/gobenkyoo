package cfg

import (
	"os"
	"path/filepath"
)

// determine default base directory on windows
var defaultBaseDir = filepath.Join(os.Getenv("USERPROFILE"), "_benkyoo")
