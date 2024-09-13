package template

import (
	"bytes"
	"html/template"
	"log"
)

func ProcessTemplate(templateFilename string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFilename)
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
