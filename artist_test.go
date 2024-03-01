package main

import (
	"sync"
	"testing"
)

func Test_getArtistInfo(t *testing.T) {
	type args struct {
		fmArtistInfo  *LastFmArtistInfo
		fmApiResponse *FmApiResponse
		wg            *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getArtistInfo(tt.args.fmArtistInfo, tt.args.fmApiResponse, tt.args.wg)
		})
	}
}
