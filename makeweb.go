package makeweb

import (
	"github.com/PaperMountainStudio/makeweb/plugins"
	"html/template"
	"io"
	"log"
	"os"
)

func render(writer io.Writer, templates *template.Template, vars map[string]interface{}) {
	// render "template" using "vars" and write output to "writer"
	template, ok := vars["template"].(string)
	if !ok {
		log.Println("WARNING: variable template is not valid, using template \"default\"")
		vars["template"] = "default"
		template = "default"
	}
	_, ok = vars["title"].(string)
	if !ok {
		log.Println("WARNING: variable title is not valid, using \"\"")
		vars["title"] = ""
	}
	err := templates.ExecuteTemplate(writer, template, vars)
	if err != nil {
		panic(err)
	}
}

func Execute() {
	err := stageLoadPlugins()
	if err != nil {
		panic(err)
	}
	// delete output directory
	// TODO: delete only files that are being overwritten
	os.RemoveAll("output")
	err = os.Mkdir("output", os.ModePerm)
	if err != nil {
		panic(err)
	}
	// TODO: write only files that need to be written

	// get global vars
	var varsGlobal map[string]interface{}
	log.Println("Collecting variables")
	ok, err := exists("vars")
	if err != nil {
		panic(err)
	}
	if ok {
		varsGlobal, err = collectVars()
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("WARNING: vars directory not found")
	}

	// get templates
	log.Println("Collecting templates")
	templates, err := collectTemplates()
	if err != nil {
		panic(err)
	}
	templates = templates.Option("missingkey=error") // throw error if key is not found

	// get pages
	log.Println("Collecting pages")
	pages, err := collectPages()
	if err != nil {
		panic(err)
	}

	// render
	log.Println("Render:")
	err = stageRender(pages, varsGlobal, templates)
	if err != nil {
		panic(err)
	}

	// hardlink files in static directory
	err = stageLink()
	if err != nil {
		panic(err)
	}

	err = plugins.EventIndependentAfter()
	if err != nil {
		panic(err)
	}
}
