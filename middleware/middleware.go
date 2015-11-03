package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/deferpanic/deferclient/deferstats"
)

type Wrapper struct {
	*logrus.Logger
	dps *deferstats.Client
}

// NewDeferPanic creates a wrapper for instance of logger.
func NewWrapper(log *logrus.Logger, dps *deferstats.Client) *Wrapper {
	return &Wrapper{log, dps}
}

func (wrapper *Wrapper) Panicln(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			wrapper.dps.BaseClient.PrepSync(err, 0)
		}
	}()

	wrapper.Logger.Panicln(args)
}

func (wrapper *Wrapper) Panic(args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			wrapper.dps.BaseClient.PrepSync(err, 0)
		}
	}()

	wrapper.Logger.Panic(args)
}

func (wrapper *Wrapper) Panicf(format string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			wrapper.dps.BaseClient.PrepSync(err, 0)
		}
	}()

	wrapper.Logger.Panicf(format, args)
}
