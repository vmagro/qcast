package libqcast

import (
	"fmt"
	"strings"

	"github.com/zmb3/spotify"
)

type Track struct {
	ID          string    `json:"spotify_id"`
	Title       string    `json:"title"`
	Artists     []*Artist `json:"artists"`
	Album       *Album    `json:"album"`
	PlayableURL string    `json:"playback_url"`
	Duration    float64   `json:"duration"`

	// These fields are only set for tracks in the Queue
	// TODO: refactor this out into another struct
	// Visible allows a track to remain in the queue (so we can eventually go back and look at
	// historical plays) but be visually and logically deleted from the queue
	Visible bool   `json:"visible"`
	QueueID string `json:"id,omitempty"`
	PartyID string `json:"party_id"`
	// AddedAt is an interface{} because when we create it it comes from r.Now() with is r.Term, but
	// in normal operation it will be time.Time
	AddedAt interface{} `json:"added_at"`
	AddedBy *User       `json:"added_by"`
}

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image"`
}

type Playlist struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ImageURL  string   `json:"image"`
	NumTracks int      `json:"num_tracks"`
	Tracks    []*Track `json:"tracks"`
}

type User struct {
	ID          string `json:"id"`
	AccessToken string `json:"-"`
	Name        string `json:"name"`
	ImageURL    string `json:"image"`
}

// gomobile is annoying and doesn't let us pass slices, so we have to make array accessor functions
func (t *Track) NumArtists() int {
	return len(t.Artists)
}
func (t *Track) ArtistAt(i int) (*Artist, error) {
	if i >= len(t.Artists) {
		return nil, fmt.Errorf("Index %d out of range (%d)", i, len(t.Artists))
	}
	return t.Artists[i], nil
}

// ArtistDisplay returns a string that has concatenated the artists of the track
func (t *Track) ArtistDisplay() string {
	// arbitrarily limit at 2 so it doesn't get out of hand
	artistNames := make([]string, min(2, len(t.Artists)))
	for i := 0; i < 2; i++ {
		if i < len(t.Artists) {
			artistNames[i] = t.Artists[i].Name
		}
	}
	return strings.Join(artistNames, ", ")
}

func (t *Track) String() string {
	return fmt.Sprintf("Track: %v by %v", t.Title, t.ArtistDisplay())
}

func TrackFromSpotify(spotify *spotify.FullTrack) *Track {
	t := Track{
		ID:          spotify.ID.String(),
		Title:       spotify.Name,
		Album:       AlbumFromSpotify(&spotify.Album),
		Artists:     make([]*Artist, len(spotify.Artists)),
		PlayableURL: string(spotify.URI),
		Duration:    spotify.TimeDuration().Seconds(),
	}
	for i := 0; i < len(spotify.Artists); i++ {
		t.Artists[i] = ArtistFromSpotify(&spotify.Artists[i])
	}
	return &t
}

func ArtistFromSpotify(spotify *spotify.SimpleArtist) *Artist {
	a := Artist{}
	a.ID = spotify.ID.String()
	a.Name = spotify.Name
	return &a
}

func AlbumFromSpotify(spotify *spotify.SimpleAlbum) *Album {
	var imageURL string
	if len(spotify.Images) > 0 {
		// the biggest image is first, in the future we might want to give smaller sizes as well, but for
		// now we will always eventually want the largest image
		imageURL = spotify.Images[0].URL
	}
	a := Album{
		ID:       spotify.ID.String(),
		Name:     spotify.Name,
		ImageURL: imageURL,
	}
	return &a
}

func PlaylistFromSpotify(spotify *spotify.SimplePlaylist) *Playlist {
	var imageURL string
	if len(spotify.Images) > 0 {
		// the biggest image is first, in the future we might want to give smaller sizes as well, but for
		// now we will always eventually want the largest image
		imageURL = spotify.Images[0].URL
	}
	p := Playlist{
		ID:        spotify.ID.String(),
		Name:      spotify.Name,
		ImageURL:  imageURL,
		NumTracks: int(spotify.Tracks.Total),
	}
	return &p
}

func UserFromSpotify(spotify *spotify.User) *User {
	var imageURL string
	if len(spotify.Images) > 0 {
		// the biggest image is first, in the future we might want to give smaller sizes as well, but for
		// now we will always eventually want the largest image
		imageURL = spotify.Images[0].URL
	}
	return &User{
		ID:       spotify.ID,
		Name:     spotify.DisplayName,
		ImageURL: imageURL,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s %s", u.ID, u.Name)
}
