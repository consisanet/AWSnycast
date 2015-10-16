package main

import (
	"flag"
	"github.com/bobtfish/AWSnycast/daemon"
	"os"
)

var (
	debug   = flag.Bool("debug", false, "Enable debugging")
	f       = flag.String("f", "/etc/awsnycast.yaml", "point configration file, default /etc/awsnycast.yaml")
	oneshot = flag.Bool("oneshot", false, "Run route table manipulation exactly once, ignoring healthchecks, then exit")
	noop    = flag.Bool("noop", false, "Don't actually *do* anything, just print what would be done")
)

func main() {
	flag.Parse()
	d := new(daemon.Daemon)
	d.Debug = *debug
	d.ConfigFile = *f
	os.Exit(d.Run(*oneshot, *noop))
}
