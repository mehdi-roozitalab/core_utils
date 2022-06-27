package core_utils

import (
	"os"
	"path/filepath"
	"syscall"
)

func Exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil || !os.IsNotExist(err)
}
func AbsolutePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		path, err := filepath.Abs(path)
		if err == nil {
			return filepath.Clean(path), nil
		}
		return path, err
	} else {
		return filepath.Clean(path), nil
	}
}
func IsSymlinkMode(mode os.FileMode) bool { return (mode & os.ModeSymlink) != 0 }
func IsSymlink(fi os.FileInfo) bool       { return IsSymlinkMode(fi.Mode()) }
func GetOwner(fi os.FileInfo) (uid, gid int) {
	if sys, ok := fi.Sys().(*syscall.Stat_t); ok {
		return int(sys.Uid), int(sys.Gid)
	} else {
		return -1, -1
	}
}
