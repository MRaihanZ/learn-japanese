package main

import (
	"encoding/json"
	"os"
)

// Type for json key
type QAData struct {
	Kana   []string `json:"kana"`
	Romaji []string `json:"romaji"`
}

func openJson(path string) ([]string, []string) {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode JSON into struct
	var data QAData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		panic(err)
	}

	// Check length consistency
	if len(data.Kana) != len(data.Romaji) {
		panic("Mismatch: number of questions and answers must be the same")
	}
	return data.Kana, data.Romaji
}
