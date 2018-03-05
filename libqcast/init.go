package libqcast

import (
	"flag"

	"github.com/golang/glog"
)

//Init sets up some important things that are necessary for libqcast to function
func Init() {
	// this seems dumb but it keeps glog from bitching
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	glog.Info("libqcast Init()")
}
