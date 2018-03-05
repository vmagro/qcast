package libqcast

import (
	"sync"

	"github.com/golang/glog"
	r "gopkg.in/gorethink/gorethink.v3"
)

var sessionMutex = sync.Mutex{}
var session *r.Session
var sessionErr error

// RethinkSession returns a session to connect to rethinkdb
func RethinkSession() *r.Session {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	if session == nil {
		// also allow rethink to read json tags on Go structs
		r.SetTags("gorethink", "json")
		var err error
		glog.Info("Attempting to connect to RethinkDB")
		session, err = r.Connect(r.ConnectOpts{
			// Address:  "vmagro-dev.lan",
			Address: "rethink.infra.qca.st",
			// Database: "qcast_dev",
			// Username: "qcast_dev",
			Database: "qcast",
			Username: "qcast",
		})
		sessionErr = err
		if err != nil {
			// TODO: do something smarter
			glog.Errorf("Error connecting to database: %s", err)
		}
	}
	return session
}

func CanConnectToDB() bool {
	return sessionErr == nil
}
