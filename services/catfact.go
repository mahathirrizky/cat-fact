package services

import (
	"encoding/json"
	"net/http"
)

// Struct to represent the response from the external API
type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// Function to call the cat fact API
func GetCatFact() (*CatFactResponse, error) {
	// Make HTTP request to the cat fact API
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response body into the struct
	var response CatFactResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
