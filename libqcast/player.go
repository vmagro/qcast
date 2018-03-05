package libqcast

import (
	"github.com/golang/glog"
)

// TODO: this could have a cleaner separation of the gomobile jank

type NativePlayerWatcher interface {
	OnPlayPause(playing bool)
	OnTrackStart(track *Track)
	OnTrackEnd()
	OnError(err error)
}

type NativePlayerInterface interface {
	SetTrack(track *Track) error
	Play() error
	Pause() error
	Notify(watcher NativePlayerWatcher)
}

type PlayerEventType byte

const (
	PlayerPaused PlayerEventType = iota
)

type PlayerEvent struct {
	Type PlayerEventType
}

type NativePlayerWrapper struct {
	backing NativePlayerInterface
	Events  chan PlayerEvent
}

func (w *NativePlayerWrapper) OnPlayPause(playing bool) {
	// TODO: update rethink here?
}

func (w *NativePlayerWrapper) OnTrackEnd() {
	// TODO: this is gross to chain this every time
	party := RethinkPartyService().CurrentParty()
	queue := party.Queue()
	queue.RemoveCurrent()
	// start playing the next Track
	if len(queue.Tracks()) > 0 {
		track := queue.Tracks()[0]
		glog.Infof("Telling player to play %s", track)
		w.backing.SetTrack(track)
	} else {
		glog.Infof("No more tracks in the queue")
	}
}

func (w *NativePlayerWrapper) OnError(err error) {
	// TODO: do something
}

func (w *NativePlayerWrapper) OnTrackStart(track *Track) {
	// this is safe to just ignore - it's only used in the mobile app
}

func (w *NativePlayerWrapper) watchQueue() {
	glog.Info("Player starting to watch queue")
	notifications := make(chan QueueEvent)
	// TODO: this chain is kinda gross
	queue := RethinkPartyService().CurrentParty().Queue()
	queue.Notify(notifications)

	currentTrack := queue.CurrentTrack()
	for _ = range notifications {
		glog.Infof("Got queue event")
		// TODO: use the actual event to determine this?
		// check if the current track has changed
		newCurrent := queue.CurrentTrack()
		if newCurrent == nil {
			glog.Infof("Current track is nil, pausing")
			currentTrack = nil
			w.backing.Pause()
		} else if newCurrent != currentTrack {
			glog.Infof("Changing current track to %+v", newCurrent)
			currentTrack = newCurrent
			w.backing.SetTrack(newCurrent)
		}
	}
}

var NativePlayer *NativePlayerWrapper

func RegisterNativePlayer(native NativePlayerInterface) {
	if NativePlayer != nil {
		glog.Fatalf("NativePlayer already registered")
	}
	NativePlayer = &NativePlayerWrapper{
		backing: native,
		Events:  make(chan PlayerEvent),
	}
	NativePlayer.backing.Notify(NativePlayer)
}
