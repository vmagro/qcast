package main

import (
	"flag"
	"github.com/golang/glog"
	"lib"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	server := lib.NewServer("amqp://guest:guest@rabbitmq-master.qcast:5672/")
	err := server.ListenAndServe()
	if err != nil {
		glog.Fatal(err)
	}
}
