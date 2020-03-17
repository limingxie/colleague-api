package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Seed(xormEngine *xorm.Engine) error {

	var (
		colleagues = []Colleague{
			{Id: 1, Name: "xiao_ming", Enable: true},
			{Id: 2, Name: "xiao_zhang", Enable: true},
			{Id: 3, Name: "lao_li", Enable: true},
			{Id: 4, Name: "lao_wang", Enable: false},
			{Id: 5, Name: "lao_zhang", Enable: true},
		}
	)
	for _, u := range colleagues {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}
	return nil
}
