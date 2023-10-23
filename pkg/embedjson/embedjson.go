package embedjson

import (
	"embed"
	"encoding/json"
	"fmt"
)

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		Nom             string  `json:"nom"`
		Lat             float64 `json:"lat"`
		Lng             float64 `json:"lng"`
		Num             string  `json:"num"`
		Voie            string  `json:"voie"`
		Ville           string  `json:"ville"`
		CodePostal      string  `json:"codepostal"`
		MarkerInnerHTML string  `json:"markerinnerhtml"`
		LiInnerHTML     string  `json:"liinnerhtml"`
		Addresse        string  `json:"addresse"`
	} `json:"properties"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

type Data struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Features []Feature `json:"features"`
}

type MainData struct {
	Data Data `json:"data"`
}

var (
	//go:embed json/*
	EmbeddedJSON embed.FS

	EmptyResult = MainData{}
)

// GetDataFromJSON extracts data from json/coiffeurs.json and put it in a data struct
func GetDataFromJSON() (data MainData, err error) {
	// Read the embedded JSON file
	file, err := EmbeddedJSON.Open("json/coiffeurs-small.json")
	if err != nil {
		return EmptyResult, fmt.Errorf("error opening embedded JSON: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error closing embedded JSON: %w", cerr)
		}
	}()

	// Decode the JSON from the file into the matrix variable
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return EmptyResult, fmt.Errorf("error decoding embedded JSON: %w", err)
	}

	return data, nil
}
