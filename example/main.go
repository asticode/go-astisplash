package main

import (
	"flag"
	"time"

	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astisplash"
	"github.com/pkg/errors"
)

func main() {
	// Init logger
	flag.Parse()
	astilog.SetLogger(astilog.New(astilog.FlagConfig()))

	// Build splasher
	astilog.Debug("Building splasher...")
	var s *astisplash.Splasher
	var err error
	if s, err = astisplash.New(); err != nil {
		astilog.Fatal(errors.Wrap(err, "building Splasher failed"))
	}
	defer s.Close()

	// Splash
	astilog.Debug("Splashing...")
	var sp *astisplash.Splash
	if sp, err = s.Splash("./example/splash.png", 484, 316); err != nil {
		astilog.Fatal(errors.Wrap(err, "splashing failed"))
	}
	defer sp.Close()

	// Sleeping
	astilog.Debug("Sleeping...")
	time.Sleep(5 * time.Second)
}
