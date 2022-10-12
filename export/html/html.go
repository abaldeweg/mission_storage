package html

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"time"

	"github.com/abaldeweg/mission_storage/mission/create"
)

var filename = "mission.json"

const tpl = `<ul>
    {{- range .Missions -}}
    {{- if eq .Private false -}}
    <li>{{ formatDate .Date }} {{ getUnit .Unit }}: {{ .Situation }}, {{ .Location }}</li>
    {{- end -}}
    {{- end -}}
</ul>`

var store create.Logfile

func Export(content []byte) (string, error) {
	var b bytes.Buffer

	if err := json.Unmarshal([]byte(content), &store); err != nil {
		return "", err
	}

	t, err := template.New("export").Funcs(template.FuncMap{
		"formatDate": formatDate,
		"getUnit":    getUnit,
	}).Parse(tpl)
	if err != nil {
		return "", err
	}

	if err = t.Execute(&b, store); err != nil {
		return "", err
	}

	return b.String(), nil
}

func formatDate(val string) string {
	t, err := time.Parse("2006-01-02", val)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return t.Format("02.01.2006")
}

func getUnit(val string) string {
	return store.Replacements[val]
}
