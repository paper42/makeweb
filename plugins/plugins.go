package plugins

import (
	"github.com/PaperMountainStudio/makeweb/gallery"
	"os"
)

// events
func EventModifyOutPath(outPath string) string {
	return PrettyOutPath(outPath)
	// return outPath
}

func EventIndependentAfter() {
	_, err := os.Stat("gallery")
	if err == nil {
		gallery.Execute()
	} else if os.IsNotExist(err) {
		// pass
	} else {
		panic(err)
	}
}
