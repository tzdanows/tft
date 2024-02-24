package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func datafetcher() {

	// API Key Load
	err := godotenv.Load("../.env")

	// Endpoint for a provided matchID
	// Note will alter this to grab batches of matchIDs in the future
	apiKey := os.Getenv("RIOT_API_KEY")
	matchID := "NA1_4913073543"
	url := fmt.Sprintf("https://americas.api.riotgames.com/tft/match/v1/matches/%s?api_key=%s", matchID, apiKey)

	fmt.Println(url)

	// HTTP request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	// Fetched response body
	body, err := io.ReadAll(response.Body)
	dataDir := "../sample-tft-data"

	// Reponse saves to json file
	filePath := filepath.Join(dataDir, fmt.Sprintf("%s.json", matchID))
	err = os.WriteFile(filePath, body, 0644)

	fmt.Printf("Match details saved to %s\n", filePath)
}
