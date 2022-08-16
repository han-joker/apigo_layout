package tool

import (
	"encoding/json"
	"os"
)

type Conf struct {
	SQL ConfSQL `json:"sql"`
}

type ConfSQL struct {
	Driver string `json:"driver,omitempty"`
	DSN    string `json:"dsn,omitempty"`
}

var defaultC = Conf{
	SQL: ConfSQL{
		Driver: "",
		DSN:    "",
	},
}

var C = defaultC

func init() {
	confFile := "conf.json"
	content, err := os.ReadFile(confFile)
	if err != nil {
		return
	}

	if err := json.Unmarshal(content, &C); err != nil {
		return
	}
}
