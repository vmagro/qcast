package libqcast

import (
	"sort"
	"time"

	"github.com/golang/glog"
	r "gopkg.in/gorethink/gorethink.v3"
)

type RethinkQueue struct {
	party     *Party
	session   *r.Session
	listeners []chan QueueEvent
	tracks    []*Track
	trackMap  map[string]*Track
}

type SafeQueue interface {
	Tracks() []*Track
	AddTrack(t *Track) error
	TrackInQueue(t *Track) bool
	RemoveCurrent() error
	CurrentTrack() *Track
}

type queue interface {
	SafeQueue
	Notify(chan QueueEvent)
}

// QueueEventType indicates which kind of an event a QueueEvent is.
type QueueEventType byte

const (
	// QueueReady is fired when the initial queue contents have downloaded.
	QueueReady QueueEventType = iota
	// TrackAdded is fired when a new track is received to be added to the queue.
	TrackAdded
	// TrackRemoved is fired when a track is removed from the queue.
	TrackRemoved
)

// QueueEvent is an event fired when the queue is updated.
type QueueEvent struct {
	Type  QueueEventType
	Track *Track
}

// NewQueue constructs a new Queue for the given Party.
// The returned Queue will lazily load data from Firebase, waiting until something registers an
// update listener or calls GetTracks()
func NewRethinkQueue(party *Party) *RethinkQueue {
	queue := &RethinkQueue{
		session:  RethinkSession(),
		party:    party,
		trackMap: make(map[string]*Track),
	}
	go queue.startUpdating()
	return queue
}

func (q *RethinkQueue) Tracks() []*Track {
	return q.tracks
}

func (q *RethinkQueue) AddTrack(t *Track) error {
	glog.Infof("Inserting %s (%s) in the queue", t.Title, t.ID)
	// set the queue information for a track that is in a queue
	t.Visible = true
	t.AddedBy = AuthService().CurrentUser
	t.AddedAt = r.Now()
	t.PartyID = q.party.ID
	res, err := r.Table("queued_tracks").Insert(t).RunWrite(q.session)
	glog.Infof("result: %+v", res)
	if err != nil {
		glog.Errorf("Error adding track to queue: %s", err)
		return err
	}
	return err
}

func (q *RethinkQueue) TrackInQueue(t *Track) bool {
	for _, track := range q.tracks {
		if t.ID == track.ID {
			return true
		}
	}
	return false
}

// RemoveCurrent deletes the currently playing track from the queue
func (q *RethinkQueue) RemoveCurrent() error {
	if len(q.tracks) == 0 {
		return nil
	}
	current := q.tracks[0]
	// don't really delete the track from the db, just make it invisible
	_, err := r.Table("queued_tracks").Get(current.QueueID).Update(map[string]interface{}{
		"visible": false,
	}).RunWrite(RethinkSession())
	return err
}

func (q *RethinkQueue) CurrentTrack() *Track {
	if len(q.tracks) > 0 {
		return q.tracks[0]
	}
	return nil
}

func (q *RethinkQueue) Notify(notifications chan QueueEvent) {
	q.listeners = append(q.listeners, notifications)
}

func (q *RethinkQueue) sortTracksFromMap() {
	// get all the tracks in the queue as a slice
	newTracks := make([]*Track, 0, len(q.trackMap))
	for _, value := range q.trackMap {
		newTracks = append(newTracks, value)
	}

	// TODO: round robin sort
	sort.SliceStable(newTracks, func(i, j int) bool {
		iTime, iOk := newTracks[i].AddedAt.(time.Time)
		jTime, jOk := newTracks[j].AddedAt.(time.Time)
		if !iOk || !jOk {
			glog.Errorf("Wasn't able to convert %s or %s into a time", newTracks[i].AddedAt, newTracks[j].AddedAt)
			return i < j
		}
		return iTime.Before(jTime)
	})
	q.tracks = newTracks
}

func (q *RethinkQueue) startUpdating() {
	glog.Infof("Watching for queue changes for %s", q.party.ID)
	// watch for changes to any parties that have the given geohash
	cursor, err := r.Table("queued_tracks").Filter(map[string]interface{}{
		"party_id": q.party.ID,
		"visible":  true,
	}).Changes(r.ChangesOpts{IncludeInitial: true, IncludeStates: true}).Run(q.session)
	// include initial and include states above so that we can get all the tracks at the beginning
	// and tell when the resulting query is ready

	if err != nil {
		// TODO: pass this to client somehow
		glog.Errorf("Error watching changefeed for queue: %s", err)
		return
	}
	defer cursor.Close()

	var change r.ChangeResponse
	for cursor.Next(&change) {
		if change.State == "ready" {
			q.sendEvent(QueueEvent{
				Type: QueueReady,
			})
			continue
		}
		if change.State != "" {
			// if we got any other state, there's nothing that we can do
			continue
		}
		if change.NewValue != nil {
			var track Track
			err := DecodeMap(change.NewValue, &track)
			if err != nil {
				glog.Warningf("Object in queue didn't seem to be a Track: %s", err)
				continue
			}
			glog.Infof("Adding %+v", track)
			// add to the queue map
			if _, ok := q.trackMap[track.ID]; !ok {
				q.trackMap[track.ID] = &track
				q.sortTracksFromMap()
				q.sendEvent(QueueEvent{
					Type:  TrackAdded,
					Track: &track,
				})
			} else {
				// TODO: handle this error somehow
			}
		}
		if change.OldValue != nil {
			var track Track
			err := DecodeMap(change.OldValue, &track)
			if err != nil {
				glog.Warningf("Object in queue didn't seem to be a Track: %s", err)
				continue
			}
			glog.Infof("Removing %+v", track)
			// remove from the queue track map
			if _, ok := q.trackMap[track.ID]; ok {
				delete(q.trackMap, track.ID)
				q.sortTracksFromMap()
				q.sendEvent(QueueEvent{
					Type:  TrackAdded,
					Track: &track,
				})
			} else {
				// TODO: handle this error somehow
				glog.Warningf("Got removed event for %+v but wasn't in the queue", track)
			}
		}
		glog.Infof("Loaded %d tracks in queue", len(q.tracks))
	}
}

func (q *RethinkQueue) sendEvent(event QueueEvent) {
	glog.Infof("Sending event on %d registered channels", len(q.listeners))
	for _, cb := range q.listeners {
		cb <- event
	}
}
