// Example of using Logrus framework with Defer Panic client library
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/deferpanic/dp-logrus/middleware"
)

// How to use Logrus properly with Defer Panic client library
func main() {
	dps := deferstats.NewClient("z57z3xsEfpqxpr0dSte0auTBItWBYa1c")

	go dps.CaptureStats()

	log := middleware.NewWrapper(logrus.New(), dps)

	log.Panicln("There is no need for panic!")

	log.Panic("There is no need for panic!")

	log.Panicf("%v", "There is no need for panic!")
}
