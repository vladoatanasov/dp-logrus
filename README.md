dp-logrus
=========
[![GoDoc](https://godoc.org/github.com/deferpanic/dp-logrus?status.svg)](https://godoc.org/github.com/deferpanic/dp-logrus)

Example of using Logrus framework with Defer Panic client library.
It allows to catch all the panics that should be issued by calling Logrus methods.

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

	log := middleware.NewWrapper(logrus.New(), dps)

	log.Panicln("There is no need for panic!")

	log.Panic("There is no need for panic!")

	log.Panicf("%v", "There is no need for panic!")
}
```
