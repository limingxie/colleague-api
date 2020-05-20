package colleagues

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Seed(xormEngine *xorm.Engine) error {

	var (
		colleagues = []Colleague{
			{Id: 1, Username: "system", Name: "系统管理员", Email: "system@email.com", Password: "1111", Enable: true},
			{Id: 2, Name: "xiao_zhang", Enable: true},
			{Id: 3, Name: "lao_li", Enable: true},
			{Id: 4, Name: "lao_wang", Enable: false},
			{Id: 5, Name: "lao_zhang", Enable: true},
		}

		stores = []Store{
			{Id: 1, Code: "C001", Name: "北京朝阳门店", TenantCode: "hublabs", Province: "北京市", City: "北京市", District: "朝阳区", Detail: "酒仙桥中路恒通商务园B37", Enable: true},
			{Id: 2, Code: "C002", Name: "北京新世界百货店", TenantCode: "hublabs", Enable: true},
			{Id: 3, Code: "C003", Name: "上海西单店", TenantCode: "hublabs", Enable: true},
		}

		storeBrands = []StoreBrand{
			{Id: 1, StoreId: 1, BrandCode: "NK", Enable: true},
			{Id: 2, StoreId: 1, BrandCode: "AD", Enable: true},
			{Id: 3, StoreId: 2, BrandCode: "NK", Enable: true},
			{Id: 4, StoreId: 3, BrandCode: "AD", Enable: true},
		}

		storeColleagues = []StoreColleague{
			{Id: 1, StoreId: 1, ColleagueId: 1, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 2, StoreId: 2, ColleagueId: 1, StartDate: "", EndDate: "", Role: "member", Enable: true},
			{Id: 3, StoreId: 3, ColleagueId: 1, StartDate: "", EndDate: "", Role: "guest", Enable: true},
			{Id: 4, StoreId: 1, ColleagueId: 2, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 5, StoreId: 1, ColleagueId: 3, StartDate: "", EndDate: "", Role: "admin", Enable: true},
		}
	)

	for _, u := range colleagues {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range stores {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range storeBrands {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range storeColleagues {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	return nil
}
