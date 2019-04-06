package makeweb

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Page struct {
	Vars    map[string]interface{}
	Content string
	Path    string
}

func plugPageInputFs() ([]Page, error) {
	err := os.Chdir("input")
	if err != nil {
		return nil, err
	}

	Pages := make([]Page, 0)
	files, err := recursiveLs(".")
	if err != nil {
		return nil, err
	}
	for _, d := range files {
		if strings.HasSuffix(d, ".html") {

			text, err := ioutil.ReadFile(d)
			if err != nil {
				return nil, err
			}

			varsstr, content := splitVarsContent(string(text))

			var vars map[string]interface{}
			err = json.Unmarshal([]byte(varsstr), &vars)
			if err != nil {
				return nil, err
			}

			Pages = append(Pages, Page{vars, content, d})
		} else {
			return nil, fmt.Errorf("extension of file %s is not currently supported in \"input\", try .html or put it in \"static\"", d)
		}
	}
	err = os.Chdir("..")
	if err != nil {
		return nil, err
	}
	return Pages, nil
}
func plugPageInputJSON() ([]Page, error) {
	Pages := make([]Page, 0)

	jsonstr, err := ioutil.ReadFile("/home/michal/temp/pages.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonstr, &Pages)
	if err != nil {
		return nil, err
	}

	return Pages, nil
}

// Load variables from filesystem
// = from json files in "vars" directory
func plugCollectVarsFs() (map[string]interface{}, error) {
	var values map[string]interface{}

	files, err := recursiveLs("vars")
	if err != nil {
		return nil, err
	}
	for _, d := range files {
		jsontxt, err := ioutil.ReadFile(d)
		if err != nil {
			return nil, err
		}

		var tmpvalues map[string]interface{}
		err = json.Unmarshal(jsontxt, &tmpvalues)
		if err != nil {
			return nil, err
		}
		values, err = joinmaps(values, tmpvalues)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
func plugCollectVarsJSON() (map[string]interface{}, error) {
	var values map[string]interface{}
	jsontxt, err := ioutil.ReadFile("vars.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsontxt, &values)
	if err != nil {
		return nil, err
	}

	return values, err
}

// collect templates
func collectTemplates() (*template.Template, error) {
	// returns a *template.Template containing all templates it found

	ok, err := exists("templates")
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("templates directory not found")
	}

	err = os.Chdir("templates")
	if err != nil {
		return nil, err
	}

	files, err := recursiveLs(".")
	if err != nil {
		return nil, err
	}

	// root := template.New(path.Base(files[0]))
	root := template.New("root")
	// root.Funcs(fm)
	templates, err := root.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	err = os.Chdir("..")
	if err != nil {
		return nil, err
	}

	return templates, nil
}

// collect global variables
func collectVars() (map[string]interface{}, error) {
	// TODO: collect from sqlite, mysql, mongodb, yaml, xml?...
	// values, err := plugCollectVarsJSON()
	values, err := plugCollectVarsFs()
	return values, err
}

func collectPages() ([]Page, error) {
	Pages, err := plugPageInputFs()
	// Pages, err := plugPageInputJSON()
	if err != nil {
		return nil, err
	}

	return Pages, nil
}
