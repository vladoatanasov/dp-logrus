package middleware

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/deferpanic/deferclient/deferstats"
)

// DPHook delivers logs to an Defer Panic web service.
type DPHook struct {
	dps *deferstats.Client
}

// NewDPHook creates a hook using an initialized Defer Panic client.
func NewDPHook(dps *deferstats.Client) (*DPHook, error) {
	if dps == nil {
		return nil, errors.New("There is no Defer Panic client")
	}
	return &DPHook{dps}, nil
}

// Fire is called when a log event is fired.
func (hook *DPHook) Fire(entry *logrus.Entry) error {
	hook.dps.BaseClient.PrepSync(entry.Message, 0)

	return nil
}

// Available logging levels.
func (hook *DPHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}
