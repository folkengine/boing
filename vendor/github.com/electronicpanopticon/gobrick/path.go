package gobrick

import (
	"fmt"
	"go/build"
	"os"
)

/// GetGOPATH returns the value of the GOPATH variable if set, and if not returns the default GOPATH.
func GetGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

/// FilePathInGoPath returns a full path from the system's GOPATH of the passed in relative path value.
/// TODO: Support for multiple GOPATHs
func FilePathInGoPath(relativePath string) string {
	return fmt.Sprintf("%s/%s", GetGOPATH(), relativePath)
}

/// FilePathInGoPath returns a file from the full path from the system's GOPATH of the passed in relative path value.
func FileInGoPath(relativePath string) (*os.File, error) {
	fp := FilePathInGoPath(relativePath)
	return os.Open(fp)
}