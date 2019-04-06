package makeweb

import (
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
)

func toHTML(page Page) (Page, error) {
	format, ok := page.Vars["format"].(string)
	if ok {
		if format == "markdown" || format == "md" {
			page.Content = md2html(page.Content)
		} else if format == "html" {
			page.Content = html2html(page.Content)
		} else {
			return Page{}, fmt.Errorf("unknown format: %v", format)
		}
	} else {
		// undefined format, use html
		page.Vars["format"] = "html"
		page.Content = html2html(page.Content)
	}
	return page, nil
}

func md2html(text string) string {
	content := string(blackfriday.Run([]byte(text)))
	return content
}

func html2html(text string) string {
	return text
}
