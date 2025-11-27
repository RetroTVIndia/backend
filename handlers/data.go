package handlers

import (
	"encoding/json"
	"log"
	"os"
)

type Video struct {
	Title       string   `json:"title"`
	Years       string   `json:"years"`
	YoutubeURLs []string `json:"youtube_urls"`
}

type Category struct {
	Name   string  `json:"name"`
	Videos []Video `json:"videos"`
}

type Era struct {
	Era        string     `json:"era"`
	Categories []Category `json:"categories"`
}

// global variable holding the era data
var EraData Era

// InitEra reads JSON file and loads EraData
func InitEra(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	if err := json.Unmarshal(file, &EraData); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
}
