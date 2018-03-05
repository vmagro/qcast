package libqcast

import "github.com/golang/glog"

type Logger interface {
	Log(event string, data interface{}) error
}

type glogLogger struct{}

func (l *glogLogger) Log(event string, data interface{}) error {
	glog.Infof("%s: %+v", event, data)
	return nil
}
