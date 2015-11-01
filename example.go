// Example of using Logrus framework with Defer Panic client library
package main

import (
	"github.com/deferpanic/deferclient/deferstats"
	//	"github.com/deferpanic/dpgorilla/middleware"
	"github.com/Sirupsen/logrus"
)

// How to use Logrus properly with Defer Panic client library
func main() {
	dps := deferstats.NewClient("z57z3xsEfpqxpr0dSte0auTBItWBYa1c")

	go dps.CaptureStats()

	log := logrus.New()
	var err error
	if err == nil {
		//       log.Hooks.Add(hooker)
	}

	log.Println("Hello, world!")
}
