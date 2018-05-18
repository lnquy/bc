package main

import (
	"log"
	"strings"

	"github.com/lnquy/blockchain/cmd/blockchain/cmd"
	"github.com/lnquy/blockchain/config"
	"github.com/lnquy/blockchain/ledger"
	"github.com/lnquy/blockchain/ledger/bigcache"
	"github.com/lnquy/blockchain/ledger/leveldb"
)

func main() {
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("main: failed to load configuration: %s", err)
	}

	var myLedger ledger.Ledger
	switch strings.ToLower(cfg.DBType) {
	case "leveldb":
		myLedger, err = leveldb.NewStorage(cfg.LevelDB)
		if err != nil {
			log.Fatalf("main: failed to init LevelDB storage: %s", err)
		}
	case "bigcache":
		myLedger, err = bigcache.NewCache()
		if err != nil {
			log.Fatalf("main: failed to init BigCache: %s", err)
		}
	default:
		log.Fatalf("main: invalid database type: %s", cfg.DBType)
	}
	defer myLedger.Close()

	cmd.Execute(cfg, myLedger)
}
