dp-logrus
=========
[![GoDoc](https://godoc.org/github.com/deferpanic/dp-logrus?status.svg)](https://godoc.org/github.com/deferpanic/dp-logrus)

Example of using Logrus framework with Defer Panic client library.

__example__

```go
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/deferpanic/deferclient/deferstats"
	"github.com/deferpanic/dp-logrus/middleware"
)

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
```
