package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Config struct {
	DBHost string `mapstructure:"db_host"`
	DBPort int    `mapstructure:"db_port"`
}

type Database struct {
	Mc *memcache.Client
}

func NewDatabase(cfg Config) *Database {
	return &Database{
		Mc: memcache.New(fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort)),
	}
}

func (db Database) Set(req *memcache.Item) error {
	err := db.Mc.Set(req)
	if err != nil {
		return err
	}

	return nil
}

func (db Database) Get(key string) (*memcache.Item, error) {
	data, err := db.Mc.Get(key)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db Database) Delete(key string) error {
	err := db.Mc.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
