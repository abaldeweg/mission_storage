package create

import (
	"encoding/json"
	"time"
)

type Logfile struct {
	Notes        []string          `json:"notes"`
	Replacements map[string]string `json:"replacements"`
	Missions     []Mission         `json:"missions"`
}

type Mission struct {
	Date      string `json:"date"`
	Situation string `json:"situation"`
	Unit      string `json:"unit"`
	Location  string `json:"location"`
	Private   bool   `json:"private"`
}

func Create(data []byte) (string, error) {
	create := Mission{
		Date: time.Now().Format("2006-01-02"),
	}

	var t Logfile
	if err := json.Unmarshal(data, &t); err != nil {
		return "", err
	}

	t.Missions = append(t.Missions, create)

	d, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
