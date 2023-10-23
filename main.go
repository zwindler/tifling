package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zwindler/tifling/pkg/embedjson"
)

func main() {
	// Read the JSON file
	data, err := embedjson.GetDataFromJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	// Initialize the random number generator
	rand.NewSource(time.Now().UnixNano())

	// Select a random entry from the list of features
	randomIndex := rand.Intn(len(data.Data.Features))
	randomEntry := data.Data.Features[randomIndex]

	// Print the random entry
	fmt.Printf("Random Entry:\n")
	fmt.Printf("Name: %s\n", randomEntry.Properties.Nom)
	fmt.Printf("Latitude: %f\n", randomEntry.Properties.Lat)
	fmt.Printf("Longitude: %f\n", randomEntry.Properties.Lng)
	// Add more fields as needed

}
