package backend

import (
	"os"
	"path/filepath"
	"strings"
)

var workDir string
var seqStr = string(filepath.Separator)
var binDir = Concat(seqStr, "bin", seqStr)

// GetWorkDir is
func GetWorkDir() string {
	if workDir == "" {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		workDir = dir
	}
	return workDir
}

// GetExeFileName is
func GetExeFileName() string {
	return filepath.Base(os.Args[0])
}

// FullPath is
func FullPath(path string) string {
	isAbs := filepath.IsAbs(path)
	if !isAbs {
		tmp, _ := filepath.Abs(path)
		if strings.Contains(tmp, binDir) {
			tmp = strings.ReplaceAll(tmp, binDir, seqStr)
		}
		if tmp != "" {
			return tmp
		}
	}
	return path
}

func Concat(path ...string) string {
	var dir string
	for _, p := range path {
		dir = filepath.Join(dir, p)
	}
	return dir
}

// CreateNotExitstDir is
func CreateNotExitstDir(dir string) {
	path := FullPath(dir)
	os.MkdirAll(path, os.ModePerm)
}
