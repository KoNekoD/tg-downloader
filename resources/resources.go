package resources

import (
	"os"
	"path"
	"runtime"
)

// init: use this package import for cd to project root(allow read resources)
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
