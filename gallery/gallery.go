package gallery

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Gallery struct {
	Name string
	Path string
}

func generateIndex(galleries []Gallery) {
	f, err := os.OpenFile("output/gallery/index.html", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t, err := template.ParseFiles("templates/gallery_index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(f, galleries)
}
func generateGallery(basepath string) {
	files, err := ioutil.ReadDir("gallery/" + basepath) // paths to all files in directory
	if err != nil {
		panic(err)
	}

	images := make([]string, 0) // paths to images only - not html or json files

	for _, d := range files {
		if !strings.HasSuffix(d.Name(), ".html") && !strings.HasSuffix(d.Name(), ".json") {
			images = append(images, d.Name())
			err = os.Link(path.Join("gallery", basepath, d.Name()), path.Join("output", "gallery", basepath, d.Name()))
			if err != nil {
				panic(err)
			}
		}
	}

	f, err := os.OpenFile(path.Join("output", "gallery", basepath, "index.html"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t, err := template.ParseFiles("templates/gallery_gallery.html")
	if err != nil {
		panic(err)
	}
	vars := make(map[string]interface{}, 0)
	vars["basepath"] = basepath
	vars["images"] = images
	t.Execute(f, vars)
}
func Execute() {
	files, err := ioutil.ReadDir("gallery")
	if err != nil {
		panic(err)
	}
	galleries := make([]Gallery, 0)
	for _, d := range files {
		if d.IsDir() {
			galleries = append(galleries, Gallery{d.Name(), d.Name()})
			err := os.MkdirAll(path.Join("output", "gallery", d.Name()), os.ModePerm)
			if err != nil {
				panic(err)
			}
			generateGallery(d.Name())
		}
	}

	generateIndex(galleries)
}
