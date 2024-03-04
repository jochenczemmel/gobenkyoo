package cfg

import (
	"fmt"
	"os/user"
	"path/filepath"
)

const subdir = ".gobenkyoo"

func UserDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return subdir, fmt.Errorf("determine current user: %w", err)
	}
	return filepath.Join(u.HomeDir, subdir), nil
}
