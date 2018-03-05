package mobile

import (
	"fmt"
	"libqcast"

	"github.com/golang/glog"
)

type QueueCallback interface {
	QueueUpdated(queue Queue)
}

type Queue interface {
	AddTrack(track *libqcast.Track) error
	RemoveCurrent() error
	NumTracks() int
	TrackAt(i int) (*libqcast.Track, error)
	TrackInQueue(t *libqcast.Track) bool
	CurrentTrack() *libqcast.Track
	RegisterUpdateListener(cb QueueCallback)
	UnregisterUpdateListener(cb QueueCallback)
}

type queueWrapper struct {
	real          *libqcast.RethinkQueue
	notifications chan libqcast.QueueEvent
	listeners     []QueueCallback
}

// WrapQueue creates a mobile-friendly Queue that is backed by a real queue that is go-friendly.
// NOTE: this funtion requires a pointer to a RethinkQueue because gomobile doesn't allow any
// exported interfaces to refer to go-only functions like channels.
func WrapQueue(queue *libqcast.RethinkQueue) Queue {
	notifications := make(chan libqcast.QueueEvent, 10)
	// cast it back to the "unsafe" (by gomobile standards) queue
	q := queueWrapper{
		real:          queue,
		notifications: notifications,
		listeners:     make([]QueueCallback, 0),
	}
	q.real.Notify(notifications)

	go q.watch()

	return &q
}

func (q *queueWrapper) AddTrack(track *libqcast.Track) error {
	return q.real.AddTrack(track)
}

func (q *queueWrapper) RemoveCurrent() error {
	return q.real.RemoveCurrent()
}

func (q *queueWrapper) NumTracks() int {
	return len(q.real.Tracks())
}

func (q *queueWrapper) TrackAt(index int) (*libqcast.Track, error) {
	tracks := q.real.Tracks()
	if index >= len(tracks) {
		return nil, fmt.Errorf("Index %d out of range (%d)", index, len(tracks))
	}
	return tracks[index], nil
}

func (q *queueWrapper) TrackInQueue(track *libqcast.Track) bool {
	return q.real.TrackInQueue(track)
}

func (q *queueWrapper) CurrentTrack() *libqcast.Track {
	tracks := q.real.Tracks()
	if len(tracks) == 0 {
		return nil
	}
	return tracks[0]
}

func (q *queueWrapper) RegisterUpdateListener(cb QueueCallback) {
	q.listeners = append(q.listeners, cb)
	// also call the callback immediately so that it doesn't have to wait for an QueueEventType
	cb.QueueUpdated(q)
}

func (q *queueWrapper) UnregisterUpdateListener(cb QueueCallback) {
	// TODO
	// it doesn't seem that the argument we get here will be equal to any of the callbacks - figure
	// this out
}

func (q *queueWrapper) watch() {
	glog.Infof("Watching for queue updates")
	for _ = range q.notifications {
		// TODO: more granular callbacks based on event type
		glog.Infof("Forwarding queue event to %d listeners", len(q.listeners))
		for _, cb := range q.listeners {
			cb.QueueUpdated(q)
		}
	}
}
