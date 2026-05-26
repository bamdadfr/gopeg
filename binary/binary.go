// Package binary defines local executable dependencies
package binary

import (
	"os"
	"os/exec"
)

var resolved = map[string]string{} // name → winning path

type Binary struct {
	Name          string
	OverwriteFlag string
	Candidates    []string // bare commands or absolute paths
}

func Init() {
	binaries := []Binary{
		Ffmpeg(),
		Video2x(),
	}

	for _, b := range binaries {
		if path := b.probe(); path != "" {
			resolved[b.Name] = path
		}
	}
}

func IsAvailable(name string) bool {
	return resolved[name] != ""
}

func ResolvedPath(name string) string {
	return resolved[name]
}

func (b Binary) probe() string {
	for _, candidate := range b.Candidates {
		if path, err := exec.LookPath(candidate); err == nil {
			return path
		}

		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return ""
}
