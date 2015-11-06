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

	log := logrus.New()

	hook, err := middleware.NewDPHook(dps)
	if err == nil {
		log.Hooks.Add(hook)
	}

	log.Error("Error is here!")

	log.Fatal("Fatal is here!")

	log.Panic("Panic is here!")
}
