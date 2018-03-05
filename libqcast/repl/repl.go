package main

import (
	"flag"

	// "libqcast"

	"github.com/golang/glog"
	"github.com/ivahaev/timer"
	"time"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	glog.Infof("Hello world")

	// session, err := r.Connect(r.ConnectOpts{
	// 	Address:  "vmagro-dev.lan",
	// 	Database: "qcast_dev",
	// })
	// if err != nil {
	// 	glog.Fatalf("Error connecting: %s", err)
	// }
	// track := libqcast.Track{
	// 	ID:    "someid",
	// 	Title: "Brother",
	// 	Artists: []*libqcast.Artist{&libqcast.Artist{
	// 		Name: "Kodaline",
	// 	}},
	// }
	// track := map[string]string{
	// 	"ID":     "someId",
	// 	"Artist": "Kodaline",
	// 	"Name":   "Brother",
	// }
	// resp, err := r.Table("queues").Insert(track).RunWrite(session)
	// if err != nil {
	// 	glog.Fatal(err)
	// }
	// glog.Infof("%d row(s) inserted", resp.Inserted)
	//
	// glog.Infof("Watching for queue updates")
	// cursor, err := r.Table("queues").Changes().Run(session)
	// if err != nil {
	// 	glog.Fatal(err)
	// }
	// defer cursor.Close()
	// var response interface{}
	// for cursor.Next(&response) {
	// 	glog.Infof("object from cursor: %+v", response)
	// }
	//

	// notifications := make(chan libqcast.PartyGeoEvent)
	// loc := &libqcast.Location{
	// 	Lat: 34.0263983,
	// 	Lng: -118.2824107,
	// }
	// libqcast.NewPartiesGeo().WatchLocation(loc, notifications)
	// for ev := range notifications {
	// 	glog.Infof("event: %+v", ev)
	// 	if ev.Party != nil {
	// 		glog.Infof("got party '%s'", ev.Party.Name)
	// 	}
	// }

	// party := libqcast.Party{
	// 	ID: "75290160-4d8e-456a-80ea-00c07702b93d",
	// }
	// queue := libqcast.NewQueue(&party)
	// glog.Info("%+v", queue)
	// for {

	// }

	t1 := timer.NewTimer(3 * time.Second)
	t1.Start()

	glog.Infof("Started 3s timer, now pausing it")
	t1.Pause()

	go func() {
		t2 := timer.NewTimer(5 * time.Second)
		t2.Start()
		glog.Infof("sleeping 5s")
		<-t2.C
		t1.Start()
		glog.Infof("restarted main")
	}()
	<-t1.C
	glog.Infof("Main timer finished")
}
