package bootstrap

import (
	"github.com/digisata/todo-gateway/pkg/memcached"
	"github.com/digisata/todo-gateway/pkg/mongo"
)

type Application struct {
	Cfg         *Config
	Mongo       mongo.Client
	MemcachedDB *memcached.Database
}

func App() (*Application, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	mongo, err := NewMongoDatabase(cfg)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Cfg:         cfg,
		Mongo:       mongo,
		MemcachedDB: memcached.NewDatabase(cfg.Memcached),
	}

	return app, nil
}

func (app Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
