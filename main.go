package main

import (
	"github.com/shihtzu-systems/bright/cmd"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {

	log.SetLevel(log.InfoLevel)

	// setup signal catching
	signals := make(chan os.Signal, 1)
	// catch all signals since not explicitly listing
	signal.Notify(signals)
	// method invoked upon seeing signal
	go func() {
		s := <-signals
		log.Tracef("%s received", s.String())
		switch s.String() {
		case "window size changes":
			log.Trace("ok")

		case "trace/breakpoint trap":
			fallthrough
		case "interrupt":
			fallthrough
		default:
			log.Fatalf("stopping!")
		}
	}()
	// run
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
