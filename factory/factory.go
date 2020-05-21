package factory

import (
	"context"
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/sirupsen/logrus"
)

var (
	xormEngine *xorm.Engine
	once       sync.Once
)

func InitDB(e *xorm.Engine) {
	once.Do(func() {
		xormEngine = e
	})
}

func DB(ctx context.Context) *xorm.Session {
	v := ctx.Value(echomiddleware.ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	if db, ok := v.(*xorm.Engine); ok {
		return db.NewSession()
	}
	panic("DB is not exist")
}

func XormEngine() *xorm.Engine {
	return xormEngine
}

func Logger(ctx context.Context) *logrus.Entry {
	v := ctx.Value(echomiddleware.ContextLoggerName)
	if v == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	if logger, ok := v.(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithFields(logrus.Fields{})
}
