package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Trace("-> main()")
	err := mozHTML()
	if err != nil {
		log.Error(err)
	}
	log.Trace("main() ->")
}
