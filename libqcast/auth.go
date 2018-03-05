package libqcast

import (
	"strings"
	"sync"

	"github.com/golang/glog"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// TODO: store these securely and change when in dev/prod mode

const SpotifyRedirectURL = "qcast-spotify-login://callback"

const SpotifyClientID = "13eaa4a01afe41c09864ac1bcb51344e"

const SpotifySwapURL = "https://safe-gorge-78588.herokuapp.com/swap"

const SpotifyRefreshURL = "https://safe-gorge-78588.herokuapp.com/refresh"

type SpotifyAuthService struct {
	CurrentUser   *User
	SpotifyClient *spotify.Client
	callbacks     []AuthCallback
}

var authServiceMutex = sync.Mutex{}
var spotifyAuthService *SpotifyAuthService

// AuthService returns an instance of SpotifyAuthService - this instance is reused across calls and
// will always have the global CurrentUser and SpotifyClient
func AuthService() *SpotifyAuthService {
	authServiceMutex.Lock()
	defer authServiceMutex.Unlock()
	if spotifyAuthService == nil {
		spotifyAuthService = &SpotifyAuthService{
			CurrentUser:   nil,
			SpotifyClient: nil,
			callbacks:     make([]AuthCallback, 0),
		}
	}
	return spotifyAuthService
}

// SpotifyLogin logs in to Spotify with the given oauth token and sets the global SpotifyClient to
// the logged in instance
func (a *SpotifyAuthService) LoginWithToken(token string) error {
	glog.Infof("Logging in to Spotify")
	oauth := oauth2.Token{
		AccessToken: token,
	}
	authenticator := spotify.Authenticator{}
	client := authenticator.NewClient(&oauth)
	a.SpotifyClient = &client

	// load user info
	sptUser, err := a.SpotifyClient.CurrentUser()
	if err != nil {
		return err
	}
	a.CurrentUser = UserFromSpotify(&sptUser.User)
	a.CurrentUser.AccessToken = token
	glog.Infof("Logged in as %s", a.CurrentUser)

	// call the callbacks
	for _, cb := range a.callbacks {
		cb.LoginSucceeded()
	}
	return nil
}

// SpotifyScopes returns a comma-separated string of all the scopes required for regular operation
// of libqcast Spotify operations
func SpotifyScopes() string {
	return strings.Join([]string{
		spotify.ScopeUserReadEmail,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopeUserLibraryRead,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState,
		"streaming",
	}, ",")
}

// SpotifyIsLoggedIn checks if we have a valid access token for the user's Spotify account
func (a *SpotifyAuthService) LoggedIn() bool {
	return a.SpotifyClient != nil && a.CurrentUser != nil
}

type AuthCallback interface {
	LoginSucceeded()
	LoginFailed(err error)
}

func (a *SpotifyAuthService) WaitForLogin(callback AuthCallback) {
	a.callbacks = append(a.callbacks, callback)
	// if we're logged in already just call immediately
	if a.LoggedIn() {
		callback.LoginSucceeded()
		return
	}
}
