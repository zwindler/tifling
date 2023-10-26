package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zwindler/tifling/pkg/embedTifhair"
)

func main() {
	data, err := embedTifhair.GetDataFromJSON()
	if err != nil {
		panic(err)
	}

	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(data.Data.Features))
	randomEntry := data.Data.Features[randomIndex]

	// Print the random entry
	fmt.Printf("Random Entry:\n")
	fmt.Printf("- Name: %s\n", randomEntry.Properties.Nom)
	fmt.Printf("- Latitude/Longitude: %f %f\n", randomEntry.Properties.Lat, randomEntry.Properties.Lng)
}
