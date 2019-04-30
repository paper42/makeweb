package makeweb

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// returns union of input maps
func joinmaps(maps ...map[string]interface{}) (map[string]interface{}, error) {
	// TODO: warn if something gets overwritten
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result, nil
}

// split variables from content by \n---\n
func splitVarsContent(text string) (string, string) {
	temp := strings.SplitN(text, "\n---\n", 2)
	if len(temp) < 2 {
		log.Println("WARNING: no \"---\" found, assuming this is all content")
		return "{}", text
	}
	vars := temp[0]
	content := temp[1]
	return vars, content
}

// recursivelly scan folder and return slice of paths (files only)
func recursiveLs(folder string) ([]string, error) {
	//TODO: Readdirnames? maybe?
	var result []string

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, d := range files {
		if d.IsDir() {
			tmp, err := recursiveLs(path.Join(folder, d.Name()))
			if err != nil {
				return nil, err
			}
			result = append(result, tmp...)
		} else {
			result = append(result, path.Join(folder, d.Name()))
		}
	}
	return result, nil
}

// https://stackoverflow.com/a/10510783
// return true if file exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
