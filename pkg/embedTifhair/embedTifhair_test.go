package embedTifhair

import (
	"testing"
)

var (
	testData = MainData{
		Data: Data{
			Features: []Feature{
				{Type: "Feature", Properties: Properties{
					Nom: "Tete En L'hair Coiffure",
					Lat: 49.181161,
					Lng: -0.3767}},
				{Type: "Feature", Properties: Properties{
					Nom: "Imagina'tif",
					Lat: 47.268834,
					Lng: 5.010264}},
			},
		},
	}
)

func TestGetDataFromJSON(t *testing.T) {
	tcs := []struct {
		path       string
		expectErr  bool
		expectData MainData
	}{
		{"json/coiffeurs-test.json", false, testData},
		{"non-existent.json", true, EmptyResult},
	}

	for _, tc := range tcs {
		t.Run(tc.path, func(t *testing.T) {
			data, err := GetDataFromJSON(tc.path)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error, but none occurred")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if !tc.expectErr {
				dataLat := data.Data.Features[0].Properties.Lat
				expectDataLat := tc.expectData.Data.Features[0].Properties.Lat
				if dataLat != expectDataLat {
					t.Errorf("Data mismatch. Expected %v, got %v", dataLat, expectDataLat)
				}
			}
		})
	}

}
