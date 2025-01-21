package data

import (
	"context"
	"pet/configs"
	"pet/internal/data/ent"
	"pet/internal/data/ent/migrate"
	"pet/pkg/entutil"

	"entgo.io/ent/dialect"
)

type Data struct {
	opt *Option
	db  *ent.Client
}

func newEntClient(conf *configs.Config) (*ent.Client, error) {
	var (
		drv dialect.Driver
		err error
	)
	drv, err = entutil.NewDriver(&entutil.Config{
		Dialect:     conf.MasterDB.Dialect,
		DSN:         conf.MasterDB.DSN,
		MaxIdle:     conf.MasterDB.MaxIdle,
		MaxActive:   conf.MasterDB.MaxActive,
		MaxLifetime: conf.MasterDB.MaxLifetime,
	})

	if err != nil {
		return nil, err
	}
	if conf.IsDebugMode() {
		drv = entutil.Debug(drv)
	}

	client := ent.NewClient(ent.Driver(drv))
	if conf.MasterDB.AutoMigrate {
		err = client.Schema.Create(
			context.Background(),
			migrate.WithForeignKeys(false),
			migrate.WithDropIndex(true),
		)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func NewData(opt *Option) (*Data, error) {
	db, err := newEntClient(opt.Configs)
	if err != nil {
		return nil, err
	}

	return &Data{
		opt: opt,
		db:  db,
	}, nil
}
