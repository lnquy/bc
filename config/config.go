package config

import "github.com/kelseyhightower/envconfig"

const (
	APP_PREFIX = "BC"
	BLOCKCHAIN_POW_DIFFICULTY = 4 // Hash of block must start with 4 leading zeros
)

type (
	Config struct {
		DBType string `envconfig:"DBTYPE" default:"leveldb"`
		LevelDB LevelDB `envconfig:"LEVELDB"`
	}

	LevelDB struct {
		DBFile string `envconfig:"DBFILE" default:"ledger.db"`
	}
)

func LoadEnvConfig() (*Config, error) {
	conf := Config{}
	err := envconfig.Process(APP_PREFIX, &conf)
	return &conf, err
}
