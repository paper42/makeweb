package plugins

import (
	"path"
	"path/filepath"
	"strings"
)

func PrettyOutPath(outPath string) string {
	if path.Base(outPath) != "index.html" {
		dir, file := path.Split(outPath)
		dir, err := filepath.Abs(dir)
		if err != nil {
			panic(err)
		}
		filefolder := strings.Split(file, ".")[0]
		newOutPath := path.Join(dir, filefolder, "index.html")
		return newOutPath
	}
	return outPath
}
