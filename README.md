# go
# Music Top Tracks Service

This service fetches the top track from Last.fm based on the specified region.

## Prerequisites

Before running the service, make sure you have the following:

- Go installed on your machine
- Last.fm API key

## Installation

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd <repository_directory>

2. Set your Last.fm API key:

Open main.go and replace the lastFmAPIKey variable value with your Last.fm API key.

3. Run the service using the following command:

    ```bash
    go run *.go

The service will be accessible at http://localhost:8080/api/v1/top_track.

## API Endpoint

### Get Top Tracks
**Endpoint:** /api/v1/top_tracks
**Method:** GET
**Query Parameter:**
    **region:** Specify the country for which you want to fetch the top track.
**Example:**
    `curl http://localhost:8080/api/v1/top_track?region=india`

## Response Structure

    ```json
    {
    "name": "Yellow",
    "duration": "267",
    "listeners": "2536975",
    "mbid": "8b5bf478-22f8-4902-a1c1-0db82261db58",
    "Lyrics": "Paroles de la chanson Yellow par Coldplay\r\n[Chris Martin]\nLook at the stars\nLook how they shine for you\nAnd everything you do\nYeah, they were all yellow\nI came along\nI wrote a song for you\nAnd all the things you do\nAnd it was called \"Yellow\"\nSo then I took my turn\nOh, what a thing to have done\nAnd it was all yellow\n\n[Chris, Jonny & Will]\n(Aah) Your skin, oh yeah, your skin and bones\n(Ooh) Turn into something beautiful\n\n(Aah) You know, you know I love you so\nYou know I love you so\n\n[Chris Martin]\nI swam across\nI jumped across for you\nOh, what a thing to do\n'Cause you were all yellow\nI drew a line\nI drew a line for you\nOh, what a thing to do\nAnd it was all yellow\n\n[Chris, Jonny & Will]\n(Aah) Your skin, oh yeah, your skin and bones\n(Ooh) Turn into something beautiful\n(Aah) And you know\nFor you, I'd bleed myself dry\n\nFor you, I'd bleed myself dry\n\n[ Chris Martin]\nIt's true, look how they shine for you\nLook how they shine for you\nLook how they shine for\nLook how they shine for you\nLook how they shine for you\nLook how they shine\n\n[ Chris Martin]\nLook at the stars\nLook how they shine for you\nAnd all the things that you do",
    "artist": {
        "name": "Coldplay",
        "mbid": "cc197bad-dc9c-440d-a5b5-d52ba2e14234",
        "Image": "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"
    }
}