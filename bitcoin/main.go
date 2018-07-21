package main

import (
	"github.com/outsideris/learning-blockchain/bitcoin/pkg/store"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("initializing...")

	blockStore, err := store.NewBlockStore("./database")
	if err != nil {
		log.Fatal(err)
	}

	log.Info(blockStore)
}
