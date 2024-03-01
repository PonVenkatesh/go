package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var lastFmAPIKey = "a368603a9aced6717cc16064d50e0d14"

// API Response structure
// We can extend this to add suggested track struct
type FmApiResponse struct {
	Tracks struct {
		Track []struct {
			Name      string `json:"name"`
			Duration  string `json:"duration"`
			Listeners string `json:"listeners"`
			MBID      string `json:"mbid"`
			Lyrics    string
			Artist    struct {
				Name  string `json:"name"`
				Id    string `json:"mbid"`
				Image string
			} `json:"artist"`
		} `json:"track"`
	} `json:"tracks"`
}

func getTopTrack(w http.ResponseWriter, r *http.Request) {

	region := r.URL.Query().Get("region")

	fmt.Println("=================================================")
	fmt.Printf("Fetching the top tracks from %s region", region)
	fmt.Println()
	fmt.Println("=================================================")

	// Call Last.fm API to get top track for the specified region
	apiURL := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=%s&api_key=%s&format=json", region, lastFmAPIKey)
	fmt.Println()
	fmt.Printf(" Hitting the lastFm API : %s", apiURL)
	response, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var fmApiResponse FmApiResponse
	var fmtArtistInfo LastFmArtistInfo
	if err := json.Unmarshal(body, &fmApiResponse); err != nil {
		http.Error(w, "Failed to unmarshal JSON response", http.StatusInternalServerError)
		return
	}

	// Extract the first track from the response
	if len(fmApiResponse.Tracks.Track) == 0 {
		http.Error(w, "No tracks found for the specified region", http.StatusNotFound)
		return
	}

	// getting lyrics and artistInfo using go routines
	// WaitGroup used to find the go routine completion
	var wg sync.WaitGroup
	wg.Add(1)
	go getLyricsByMBID(&fmApiResponse, &wg)
	wg.Add(1)
	go getArtistInfo(&fmtArtistInfo, &fmApiResponse, &wg)

	wg.Wait()

	// Create TrackInfo struct from Last.fm response
	firstTrack := fmApiResponse.Tracks.Track[0]
	fmt.Println(" === Name : ", firstTrack.Artist.Name)
	fmt.Println(" === Track : ", firstTrack.Name)
	fmt.Println(" === Lyrics : ", firstTrack.Lyrics)

	firstTrack.Artist.Image = fmtArtistInfo.Artist.Image[0].Text
	trackInfo := firstTrack
	// Convert the trackInfo struct to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trackInfo)
}

func main() {
	// Register the handler function for /track/
	fmt.Println("Starting the service..")
	http.HandleFunc("/api/v1/top_tracks", getTopTrack)

	// Start the server on port 8080
	fmt.Println("Listening the port 8080")
	http.ListenAndServe(":8080", nil)
}
