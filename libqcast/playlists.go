package libqcast

import (
	"fmt"

	"github.com/golang/glog"
)

type PlaylistsResponse struct {
	Playlists []*Playlist
	Error     error
}

type PlaylistResponse struct {
	Playlist *Playlist
	Error    error
}

type PlaylistService interface {
	FetchPlaylists(notify chan<- *PlaylistsResponse)
	FetchPlaylistTracks(playlist *Playlist, notify chan<- *PlaylistResponse)
}

type spotifyPlaylists struct {
}

func (s *spotifyPlaylists) FetchPlaylists(notify chan<- *PlaylistsResponse) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				glog.Errorf("Recovered from panic %s", r)
				notify <- &PlaylistsResponse{Error: fmt.Errorf("%+v", r)}
			}
		}()

		glog.Info("Fetching playlists for current user")
		if AuthService().SpotifyClient == nil {
			glog.Errorf("Trying to get playlists when not logged in")
			notify <- &PlaylistsResponse{Error: fmt.Errorf("not logged in")}
			return
		}
		page, err := AuthService().SpotifyClient.CurrentUsersPlaylists()
		// TODO: make sure to load the next page if it exists
		if err != nil {
			glog.Errorf("Error fetching playlists: %s", err)
			// TODO: handle error
			return
		}
		playlists := make([]*Playlist, len(page.Playlists))
		for i, spotify := range page.Playlists {
			playlists[i] = PlaylistFromSpotify(&spotify)
		}
		glog.Infof("Loaded %d playlists", len(playlists))
		notify <- &PlaylistsResponse{Playlists: playlists}
	}()
}

func (s *spotifyPlaylists) FetchPlaylistTracks(playlist *Playlist, notify chan<- *PlaylistResponse) {
	// TODO: implement this
}

// NewSpotifyPlaylistService creates a PlaylistService that can load playlists from the Spotify API
func NewSpotifyPlaylistService() PlaylistService {
	return &spotifyPlaylists{}
}
