package helper

import (
	"bytes"
	"log"
	"text/template"
	"time"
)

func ParseTemplate(templateFilename string, data interface{}) (string, error) {
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

func TimeHostNow(tz string) time.Time {
	// you can change Asia/Jakarta with your own location.
	// check on this https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}
