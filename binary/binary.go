// Package binary defines local executable dependencies
package binary

type BinaryPath struct {
	Linux   string
	Windows string
}

type Binary struct {
	Name          string
	Command       string
	Path          BinaryPath
	OverwriteFlag string
}
