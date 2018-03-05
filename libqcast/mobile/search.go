package mobile

import (
	"fmt"
	"libqcast"
)

// SearchResults is a struct that holds search results from the Spotify API after being massaged into
// QCast datatypes - this is the gomobile friendly version.
type SearchResults struct {
	tracks  []*libqcast.Track
	albums  []*libqcast.Album
	artists []*libqcast.Artist
	// there might have been an error grabbing the search results
	Error error
}

// We need some helper methods since gomobile currently doesn't support passing slices of structs
// across the language barrier

// NumTracks returns how many tracks were in the results.
func (s *SearchResults) NumTracks() int {
	return len(s.tracks)
}

// TrackAt is a way to index into the search results since gomobile doesn't allow passing structs.
func (s *SearchResults) TrackAt(i int) (*libqcast.Track, error) {
	if i >= len(s.tracks) {
		return nil, fmt.Errorf("Index %d out of range (%d)", i, len(s.tracks))
	}
	return s.tracks[i], nil
}

// SearchCallback is a gomobile-friendly way for the Searcher to return results to the client.
type SearchCallback interface {
	SearchResultsReceived(results *SearchResults)
}

// SearchService is an interface for a service that can search for tracks/artists/albums.
type SearchService interface {
	Search(query string, callback SearchCallback)
}

type queuedSearch struct {
	query    string
	callback SearchCallback
}

type searchWrapper struct {
	real  libqcast.SearchService
	queue chan queuedSearch
}

// Search proxies the search request to the real SearchService and then calls the callbacks.
// Additionally, this function implements throttling/debouncing preventing a rampant number of
// in-flight search requests.
func (s *searchWrapper) Search(query string, callback SearchCallback) {
	s.queue <- queuedSearch{
		query:    query,
		callback: callback,
	}
}

func (s *searchWrapper) debounce() {
	// TODO: debounce/rate limit
	for search := range s.queue {
		s.doSearch(search)
	}
}

func (s *searchWrapper) doSearch(search queuedSearch) {
	notify := make(chan *libqcast.SearchResults)
	go s.real.Search(search.query, notify)
	results := <-notify
	search.callback.SearchResultsReceived(&SearchResults{
		tracks:  results.Tracks,
		albums:  results.Albums,
		artists: results.Artists,
	})
}

// SpotifySearchService returns a mobile wrapper around the go SpotifySearchService
func SpotifySearchService() SearchService {
	s := searchWrapper{
		real:  libqcast.NewSpotifySearch(),
		queue: make(chan queuedSearch, 10),
	}
	go s.debounce()
	return &s
}
