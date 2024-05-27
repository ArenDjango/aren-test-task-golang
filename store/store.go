package store

import (
	"context"
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger/log"

	"github.com/ArenDjango/golang-test-task/store/pg"
	"github.com/pkg/errors"

	"time"
)

type Store struct {
	Pg    *pg.DB
	Rates RatesRepo
}

func New(ctx context.Context) (*Store, error) {
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	if pgDB != nil {
		log.Infof("Running PostgreSQL migrations...")
		if err := runPgMigrations(); err != nil {
			return nil, errors.Wrap(err, "runPgMigrations failed")
		}
	}

	var store Store

	if pgDB != nil {
		store.Pg = pgDB
		go store.KeepAlivePg()
		store.Rates = pg.NewRatesRepo(pgDB)
	}

	return &store, nil
}

const KeepAlivePollPeriod = 3

func (store *Store) KeepAlivePg() {
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.Pg == nil {
			lostConnect = true
		} else if _, err = store.Pg.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		log.Warnf("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		store.Pg, err = pg.Dial()
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("[store.KeepAlivePg] PostgreSQL reconnected")
	}
}
