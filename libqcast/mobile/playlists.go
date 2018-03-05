package mobile

import (
	"fmt"
	"libqcast"
)

type PlaylistsList struct {
	Playlists []*libqcast.Playlist
	Error     error
}

func (p *PlaylistsList) NumPlaylists() int {
	return len(p.Playlists)
}

func (p *PlaylistsList) PlaylistAt(i int) (*libqcast.Playlist, error) {
	if i < len(p.Playlists) {
		return p.Playlists[i], nil
	}
	return nil, fmt.Errorf("Playlist %d out of range, length %d", i, len(p.Playlists))
}

type PlaylistsCallback interface {
	PlaylistsReceived(playlists *PlaylistsList)
}

type PlaylistCallback interface {
	PlaylistTracksReceived(playlist *libqcast.Playlist)
}

type PlaylistService interface {
	FetchPlaylists(callback PlaylistsCallback)
	FetchPlaylistTracks(playlist *libqcast.Playlist, callback PlaylistCallback)
}

type playlistWrapper struct {
	real libqcast.PlaylistService
}

func NewSpotifyPlaylistService() PlaylistService {
	return &playlistWrapper{
		real: libqcast.NewSpotifyPlaylistService(),
	}
}

func (w *playlistWrapper) FetchPlaylists(callback PlaylistsCallback) {
	go func() {
		notify := make(chan *libqcast.PlaylistsResponse)
		w.real.FetchPlaylists(notify)
		response := <-notify
		callback.PlaylistsReceived(&PlaylistsList{
			Playlists: response.Playlists,
			Error:     response.Error,
		})
	}()
}

func (w *playlistWrapper) FetchPlaylistTracks(playlist *libqcast.Playlist, callback PlaylistCallback) {
	// TODO: implement this
}
