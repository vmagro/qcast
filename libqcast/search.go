package libqcast

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/zmb3/spotify"
)

// SearchResults is a struct that holds search results from the Spotify API after being massaged into
// QCast datatypes
type SearchResults struct {
	Tracks  []*Track
	Albums  []*Album
	Artists []*Artist
	// there might have been an error grabbing the search results
	Error error
}

// SearchService provides an interface to search for Tracks, Artists and Albums from the Spotify API
type SearchService interface {
	// Search will hit the Spotify API to retrieve search results for the given string
	Search(query string, notify chan<- *SearchResults)
}

type spotifySearch struct {
}

func NewSpotifySearch() SearchService {
	return &spotifySearch{}
}

func (s *spotifySearch) Search(query string, notify chan<- *SearchResults) {
	defer func() {
		if r := recover(); r != nil {
			glog.Warningf("Recovered from panic in doSearch %s", r)
			notify <- &SearchResults{
				Error: fmt.Errorf("%+v", r),
			}
		}
	}()

	glog.Infof("Querying for %s", query)

	// construct a SearchResults struct we'll use to combine all the search results
	results := &SearchResults{}

	// if the query is empty, we can just return an empty result set now
	if query == "" {
		notify <- results
		return
	}

	// make sure that only tracks playable by this user show up in the results
	country := "from_token"
	spotifyResults, err := AuthService().SpotifyClient.SearchOpt(query,
		spotify.SearchTypeTrack|spotify.SearchTypeAlbum|spotify.SearchTypeArtist,
		&spotify.Options{
			Country: &country,
		})
	// TODO: check error better
	if err != nil {
		glog.Errorf("error: %s\n", err)
		results.Error = err
		notify <- results
		return
	}
	tracks := spotifyResults.Tracks.Tracks
	results.Tracks = make([]*Track, len(tracks))
	for i := 0; i < len(tracks); i++ {
		results.Tracks[i] = TrackFromSpotify(&tracks[i])
	}
	// TODO: process results for artists and albums too

	glog.Infof("Processed results for %s", query)
	notify <- results
}
