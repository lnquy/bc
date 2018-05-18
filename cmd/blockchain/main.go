package main

import (
	"log"
	"strings"

	"github.com/lnquy/bc/cmd/blockchain/cmd"
	"github.com/lnquy/bc/config"
	"github.com/lnquy/bc/storage"
	"github.com/lnquy/bc/storage/bigcache"
	"github.com/lnquy/bc/storage/leveldb"
)

func main() {
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("main: failed to load configuration: %s", err)
	}

	var ledger storage.Ledger
	switch strings.ToLower(cfg.DBType) {
	case "leveldb":
		ledger, err = leveldb.NewStorage(cfg.LevelDB)
		if err != nil {
			log.Fatalf("main: failed to init LevelDB storage: %s", err)
		}
	case "bigcache":
		ledger, err = bigcache.NewCache()
		if err != nil {
			log.Fatalf("main: failed to init BigCache: %s", err)
		}
	default:
		log.Fatalf("main: invalid database type: %s", cfg.DBType)
	}
	defer ledger.Close()

	cmd.Execute(cfg, ledger)
}
