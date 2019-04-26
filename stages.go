package makeweb

import (
	"bytes"
	"encoding/json"
	"github.com/PaperMountainStudio/makeweb/plugins"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func stageLoadPlugins() error {
	var plugs []string

	ok, err := exists("plugins.json")
	if err != nil {
		return err
	}
	if ok {
		jsonstr, err := ioutil.ReadFile("plugins.json")
		if err != nil {
			return err
		}

		err = json.Unmarshal(jsonstr, &plugs)
		if err != nil {
			return err
		}
	}
	err = plugins.Init(plugs)
	if err != nil {
		return err
	}
	return nil
}

func stageLink() error {
	ok, err := exists("static")
	if err != nil {
		return err
	}
	if !ok {
		log.Println("WARNING: static directory not found")
		return nil
	}

	err = os.Chdir("static")
	if err != nil {
		return err
	}
	files, err := recursiveLs(".")
	if err != nil {
		return err
	}
	log.Println("Link:")
	for _, d := range files {
		outPath := path.Join("../output/", d)
		err := os.MkdirAll(path.Dir(outPath), os.ModePerm)
		if err != nil {
			return err
		}
		log.Println("- " + d)
		if _, err = os.Stat(outPath); !os.IsNotExist(err) {
			log.Println("already there")
			continue
		}
		err = os.Link(d, outPath)
		if err != nil {
			return err
		}
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

func stageRender(pages []Page, varsGlobal map[string]interface{}, templates *template.Template) error {
	for _, page := range pages {
		outPath := path.Join("output/", page.Path)
		outPath = plugins.EventModifyOutPath(outPath)
		err := os.MkdirAll(path.Dir(outPath), os.ModePerm)
		if err != nil {
			return err
		}

		log.Println("- " + page.Path)

		// convert .text to html if needed
		page, err = toHTML(page)
		if err != nil {
			return err
		}

		vars, err := joinmaps(varsGlobal, page.Vars)
		if err != nil {
			return err
		}

		temporaryVars := vars
		contentwriter := bytes.NewBufferString("")
		contentTemplate := template.New("default")
		_, err = contentTemplate.Parse(page.Content)
		if err != nil {
			return err
		}
		temporaryVars["template"] = "default"
		err = render(contentwriter, contentTemplate, temporaryVars)
		if err != nil {
			return err
		}
		page.Content = contentwriter.String()
		vars["text"] = template.HTML(page.Content)

		// write
		f, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		err = render(f, templates, vars)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
