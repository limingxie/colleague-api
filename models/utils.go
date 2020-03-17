package models

import (
	"github.com/go-xorm/xorm"
)

func SetXormEngineSync(xormEngine *xorm.Engine) {
	//xormEngine.ShowSQL(true)
	xormEngine.Sync(new(Colleague))
}
