package template

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func ProcessTemplate(templateFilename, templateVersion string, data interface{}) (string, error) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}
	t, err := template.New(fmt.Sprintf("template-%s.html", templateVersion)).Funcs(funcMap).ParseFiles(templateFilename)
	if err != nil {
		log.Printf("Error while parse template file:%+v\n", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println("Error while execute template:", err)
		return "", err

	}

	return buf.String(), nil
}
