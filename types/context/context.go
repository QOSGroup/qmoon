package context

import (
	"context"
	"database/sql"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/log"
)

type Context struct {
	context.Context

	db  *sql.DB
	log log.Logger
}

type contextKey int // local to the context module
const (
	contextKeyChainID contextKey = iota
	contextKeyLogger
	contextKeyDB
	contextKeyTendermintCli
	contextKeyQstarsCli
)

func NewContext(db *sql.DB, l log.Logger) Context {
	c := Context{
		Context: context.Background(),
		db:      db,
		log:     l,
	}

	return c
}

func (c Context) withValue(key interface{}, value interface{}) Context {
	return Context{
		Context: context.WithValue(c.Context, key, value),
		db:      c.db,
	}
}

func (c Context) WithValue(key interface{}, value interface{}) Context {
	return c.withValue(key, value)
}

func (c Context) WithLogger(logger log.Logger) Context { return c.withValue(contextKeyLogger, logger) }

func (c Context) WithChainID(chainID string) Context {
	return c.withValue(contextKeyChainID, chainID)
}

func (c Context) WithTendermintCli(tmcli *lib.TmClient) Context {
	return c.withValue(contextKeyTendermintCli, tmcli)
}

func (c Context) WithQstarsCli(qstarsCli *lib.QstarsClient) Context {
	return c.withValue(contextKeyQstarsCli, qstarsCli)
}

//
func (c Context) Logger() log.Logger { return c.log }
func (c Context) ChainID() string    { return c.Value(contextKeyChainID).(string) }
func (c Context) DB() *sql.DB        { return c.Value(contextKeyDB).(*sql.DB) }
func (c Context) TendermintCli() *lib.TmClient {
	return c.Value(contextKeyTendermintCli).(*lib.TmClient)
}
func (c Context) QstarsCli() *lib.QstarsClient {
	return c.Value(contextKeyQstarsCli).(*lib.QstarsClient)
}
