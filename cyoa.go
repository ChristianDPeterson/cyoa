package cyoa

import (
	"encoding/json"
)

func ParseJSON(file []byte) Adventure {
	var adventure Adventure
	if err := json.Unmarshal(file, &adventure); err != nil {
		panic(err)
	}
	return adventure
}

type Scene struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Adventure map[string]Scene
