package mobile

import "libqcast"

// PartiesGeoCallback is a mobile-friendly way to show nearby parties changes in the UI.
type PartiesGeoCallback interface {
	PartyEntered(*libqcast.Party)
	PartyExited(*libqcast.Party)
	PartiesReady()
}

// PartiesGeoService is the mobile-friendly wrapper around libqcast.PartiesGeoService.
type PartiesGeoService interface {
	WatchParties(loc *libqcast.Location, callback PartiesGeoCallback)
	AddParty(party *libqcast.Party) error
}

type partiesGeoWrapper struct {
	real libqcast.PartiesGeoService
}

func (w *partiesGeoWrapper) WatchParties(loc *libqcast.Location, callback PartiesGeoCallback) {
	notifications := make(chan libqcast.PartyGeoEvent)
	w.real.WatchLocation(loc, notifications)
	go func() {
		for ev := range notifications {
			party := ev.Party
			if ev.Type == libqcast.PartyEntered {
				callback.PartyEntered(party)
			}
			if ev.Type == libqcast.PartyExited {
				callback.PartyExited(party)
			}
			if ev.Type == libqcast.GeoPartiesReady {
				callback.PartiesReady()
			}
		}
	}()
}

func (w *partiesGeoWrapper) AddParty(party *libqcast.Party) error {
	return w.real.AddParty(party)
}

// NewPartiesGeo creates a mobile-friendly instance of PartiesGeoService
func NewPartiesGeo() PartiesGeoService {
	return &partiesGeoWrapper{
		real: libqcast.NewPartiesGeo(),
	}
}
