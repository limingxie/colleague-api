package tenants

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Seed(xormEngine *xorm.Engine) error {

	var (
		tenants = []Tenant{
			{Code: "hublabs", Name: "hublabs", Enable: true},
		}

		brands = []Brand{
			{Code: "NK", Name: "Nike", TenantCode: "hublabs", Enable: true},
			{Code: "AD", Name: "Adidas", TenantCode: "hublabs", Enable: true},
		}
	)

	for _, u := range tenants {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range brands {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	return nil
}
