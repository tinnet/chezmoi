// +build !windows

package chezmoi

import (
	"os"
	"path"
	"syscall"
)

func init() {
	Umask = os.FileMode(syscall.Umask(0))
	syscall.Umask(int(Umask))
}

// PathJoin returns a clean, absolute path. If file is not an absolute path then
// it is joined on to dir.
func PathJoin(dir, file string) string {
	if !path.IsAbs(file) {
		file = path.Join(dir, file)
	}
	return path.Clean(file)
}

// PathsToSlashes returns paths.
func PathsToSlashes(paths []string) []string {
	return paths
}
