package main

import (
	"flag"
	"github.com/golang/glog"
	"lib"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	server := lib.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		glog.Fatal(err)
	}
}
