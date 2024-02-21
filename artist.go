package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type LastFmArtistInfo struct {
	Artist struct {
		Name  string `json:"name"`
		Image []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
	} `json:"artist"`
}

func getArtistInfo(fmArtistInfo *LastFmArtistInfo, fmApiResponse *FmApiResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	firstTrack := fmApiResponse.Tracks.Track[0]
	artistName := firstTrack.Artist.Name

	apiURL := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=%s&api_key=%s&format=json", url.QueryEscape(artistName), lastFmAPIKey)
	response, err := http.Get(apiURL)
	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {

		return
	}

	if err := json.Unmarshal(body, &fmArtistInfo); err != nil {

		fmt.Println(err)
		return
	}
	return
}
