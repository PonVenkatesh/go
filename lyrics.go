package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const lyricsOvhAPI = "https://api.lyrics.ovh/v1"

type LyricsResponse struct {
	Lyrics string `json:"lyrics"`
}

func getLyricsByMBID(fmApiResponse *FmApiResponse, wg *sync.WaitGroup) error {
	defer wg.Done()
	// Check if there are tracks in the response
	if len(fmApiResponse.Tracks.Track) == 0 {
		fmt.Println("no tracks in the Last.fm API response")
		return fmt.Errorf("no tracks in the Last.fm API response")
	}

	// Construct the API URL
	firstTrack := fmApiResponse.Tracks.Track[0]
	artist := firstTrack.Artist.Name
	title := firstTrack.Name
	apiURL := fmt.Sprintf("%s/%s/%s", lyricsOvhAPI, artist, title)

	fmt.Println("Lyrics URL:", apiURL)

	// Make a GET request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch lyrics: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var lyricsResponse LyricsResponse
	err = json.Unmarshal(body, &lyricsResponse)
	if err != nil {
		return fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Update the track's lyrics in the FmApiResponse structure
	firstTrack.Lyrics = lyricsResponse.Lyrics
	fmApiResponse.Tracks.Track[0] = firstTrack
	return nil
}
