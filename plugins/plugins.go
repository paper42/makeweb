package plugins

import (
	"fmt"
	"github.com/PaperMountainStudio/makeweb/gallery"
	"os"
	"strings"
)

type Plugins struct {
	Gallery    bool
	PrettyPath bool
}

var plugins Plugins

func Init(plugs []string) error {
	for _, plugin := range plugs {
		switch strings.ToLower(plugin) {
		case "gallery":
			plugins.Gallery = true
		case "prettypath":
			plugins.PrettyPath = true
		default:
			return fmt.Errorf("Plugin %s not valid", plugin)
		}
	}
	return nil
}

// events
func EventModifyOutPath(outPath string) string {
	if plugins.PrettyPath == true {
		return PrettyOutPath(outPath)
	}
	return outPath
}

func EventIndependentAfter() error {
	if plugins.Gallery == true {
		_, err := os.Stat("gallery")
		if err == nil {
			gallery.Execute()
		} else if os.IsNotExist(err) {
			// pass
		} else {
			return err
		}
	}
	return nil
}
