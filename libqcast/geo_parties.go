package libqcast

import (
	"github.com/golang/glog"
	r "gopkg.in/gorethink/gorethink.v3"
)

// PartyGeoEventType indicates what event type an event is.
type PartyGeoEventType byte

const (
	// PartyEntered is triggered when a new Party appears in the watched region.
	PartyEntered PartyGeoEventType = iota
	// PartyExited is triggered when a Party leaves the watched region.
	PartyExited
	// GeoPartiesReady indicates that any parties present in the database will have alrady been
	// loaded.
	GeoPartiesReady
)

// PartyGeoEvent is an event fired by the PartiesGeoService when a nearby party appears, disappears,
// etc.
type PartyGeoEvent struct {
	Type  PartyGeoEventType
	Party *Party
}

type PartiesGeoService interface {
	WatchLocation(loc *Location, notifications chan PartyGeoEvent)
	AddParty(party *Party) error
}

type rethinkPartiesGeo struct {
	session *r.Session
}

func NewPartiesGeo() PartiesGeoService {
	glog.Infof("Creating PartiesGeoService")
	return &rethinkPartiesGeo{
		session: RethinkSession(),
	}
}

func (g *rethinkPartiesGeo) WatchLocation(loc *Location, notifications chan PartyGeoEvent) {
	// spawn a goroutine so we can watch in the background
	go func() {
		// watch for changes to any parties that within the given range
		point := r.Point(loc.Lng, loc.Lat)
		cursor, err := r.Table("parties").Filter(func(party r.Term) r.Term {
			// see if the party is within 2 miles of the current location
			return party.Field("location").Intersects(r.Circle(point, 2, r.CircleOpts{Unit: "mi"}))
		}).Changes(r.ChangesOpts{IncludeInitial: true, IncludeStates: true}).Run(g.session)
		// include initial and include states above so that we can get all the parties at the beginning
		// and tell when the resulting query is ready

		if err != nil {
			// TODO: pass this to client somehow
			glog.Errorf("Error watching changefeed for parties: %s", err)
			return
		}
		defer cursor.Close()

		var change r.ChangeResponse
		for cursor.Next(&change) {
			glog.Infof("Got change: %+v", change)
			if change.State == "ready" {
				notifications <- PartyGeoEvent{Type: GeoPartiesReady}
				continue
			}
			if change.State != "" {
				// if we got any other state, there's nothing that we can do
				continue
			}
			var party Party
			event := PartyGeoEvent{}
			if change.NewValue != nil {
				// the party has been added
				event.Type = PartyEntered
				DecodeMap(change.NewValue, &party)
				glog.Infof("%+v -> %+v", change.NewValue, party)
			} else if change.OldValue != nil {
				// the party has been deleted
				event.Type = PartyExited
				DecodeMap(change.OldValue, &party)
			}
			event.Party = &party
			notifications <- event
		}
	}()
}

func (g *rethinkPartiesGeo) AddParty(party *Party) error {
	// glog.Infof("Adding location information to party: %+v", party)
	point := r.Point(party.location.Lng, party.location.Lat)
	_, err := r.Table("parties").Get(party.ID).Update(map[string]interface{}{
		"location": point,
	}).RunWrite(g.session)
	if err != nil {
		glog.Errorf("Error adding location to party: %s", err)
		return err
	}
	return nil
}
