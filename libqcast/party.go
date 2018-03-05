package libqcast

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/glog"
	r "gopkg.in/gorethink/gorethink.v3"
)

// Party is a currently in-progress party - that is it exists and it has a queue
type Party struct {
	ID                  string `gorethink:"id,omitempty"`
	Code                string `json:"code"`
	Name                string `json:"name"`
	MinSupportedVersion int    `json:"min_supported_version"`
	Host                *User  `json:"host"`
	location            *Location
	queue               *RethinkQueue
}

// TODO: this is some gross tight coupling returning a RethinkQueue - this should be refactored
func (p *Party) Queue() *RethinkQueue {
	if p.queue == nil {
		glog.Infof("Party %s didn't have a queue, creating one now", p.Code)
		p.queue = NewRethinkQueue(p)
	}
	return p.queue
}

func (p *Party) AmIHost() bool {
	return p.Host.ID == AuthService().CurrentUser.ID
}

func (p *Party) BecomePlayer() error {
	if NativePlayer == nil {
		return errors.New("NativePlayer is nil (not registered)")
	}
	go NativePlayer.watchQueue()
	return nil
}

// CurrentParty is the global party that we are currently in (may be nil)
// var CurrentParty Party

// PartyService provides functionality to get nearby parties, join a party and create new parties.
type PartyService interface {
	// LoadParty loads information about a Party from the backend
	LoadParty(code string) (*Party, error)
	// StartParty takes information from the host and creates a new party in the backend
	StartParty(name string, location *Location, moodPlaylist *Playlist) (*Party, error)
	// JoinParty joins the given Party
	JoinParty(party *Party) error
	// CurrentParty returns the party that the user is currently part of, or nil otherwise
	CurrentParty() *Party
	/// LeaveParty leaves the current party
	LeaveParty() error
}

var globalPartyService *rethinkPartyService

var globalPartyServiceMutex sync.Mutex

// RethinkPartyService returns an implementation of PartyService based on rethinkdb, it is created
// on the first call and will be the same instance on all subsequent calls
func RethinkPartyService() PartyService {
	globalPartyServiceMutex.Lock()
	defer globalPartyServiceMutex.Unlock()
	// if it doesn't exist create it now
	if globalPartyService == nil {
		globalPartyService = newRethinkPartyService()
	}
	return globalPartyService
}

type rethinkPartyService struct {
	session      *r.Session
	currentParty *Party
}

func newRethinkPartyService() *rethinkPartyService {
	return &rethinkPartyService{
		session:      RethinkSession(),
		currentParty: nil,
	}
}

func (f *rethinkPartyService) LoadParty(code string) (*Party, error) {
	glog.Infof("Loading Party %s", code)

	// TODO: find the party that has the given shortcode
	// if err != nil {
	// 	glog.Errorf("Error retreiving party from db: %+v", err)
	// }
	// return &party, err
	return nil, nil
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateCode generates a random 5 digit party code
func GenerateCode() string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (p *rethinkPartyService) StartParty(name string, location *Location, moodPlaylist *Playlist) (*Party, error) {
	glog.Infof("Starting new party %s %s %s", name, location, moodPlaylist)
	// TODO generate a code randomly
	code := GenerateCode()
	party := Party{
		Code:                code,
		Name:                name,
		Host:                AuthService().CurrentUser,
		location:            location,
		MinSupportedVersion: APIVersion,
	}
	// save the party to rethink
	response, err := r.Table("parties").Insert(party).RunWrite(p.session)
	if err != nil {
		glog.Errorf("Error saving party in db %+v", err)
		return nil, err
	}
	// get the generated party id
	party.ID = response.GeneratedKeys[0]
	glog.Infof("Created party with id=%s", party.ID)

	err = NewPartiesGeo().AddParty(&party)
	if err != nil {
		glog.Errorf("Error saving geolocation information in db %+v", err)
		return nil, err
	}
	return &party, nil
}

func (p *rethinkPartyService) JoinParty(party *Party) error {
	glog.Infof("Joining")
	if party == nil {
		glog.Errorf("Tried to a join a nil Party")
		return errors.New("party can't be nil")
	}
	glog.Infof("Joining party %+v", party)
	if p.currentParty != nil && p.currentParty != party {
		err := fmt.Errorf("Attempted to Join party %s but already in party %s", party.ID, p.currentParty.ID)
		glog.Error(err)
		return err
	}
	if party.MinSupportedVersion > APIVersion {
		glog.Errorf("Party is beyond our API version (%d > %d)", party.MinSupportedVersion, APIVersion)
		return errors.New("api-version-fail")
	}
	p.currentParty = party
	return nil
}

func (p *rethinkPartyService) CurrentParty() *Party {
	return p.currentParty
}

func (p *rethinkPartyService) LeaveParty() error {
	if p.currentParty == nil {
		glog.Error("Attempted to Leave party but not in one")
		return errors.New("Attempted to Leave party but not in one")
	}
	p.currentParty = nil
	return nil
}
