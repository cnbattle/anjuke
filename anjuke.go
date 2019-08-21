package main

import (
	"github.com/cnbattle/anjuke/cmd"
	"github.com/cnbattle/anjuke/config"
)

func main() {
	if config.V.IsAll {
		cmd.GrabAll()
	} else {
		cmd.Grab()
	}
}
