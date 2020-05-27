package colleagues

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Init(xormEngine *xorm.Engine) error {
	return xormEngine.Sync(
		new(Colleague),
		new(Store),
		new(StoreBrand),
		new(StoreColleague),
		new(App),
		new(ColleagueApp),
	)
	//xormEngine.ShowSQL(true)

}

func DropTables(xormEngine *xorm.Engine) error {
	return xormEngine.DropTables(
		new(Colleague),
		new(Store),
		new(StoreBrand),
		new(StoreColleague),
		new(App),
		new(ColleagueApp),
	)
}

func Seed(xormEngine *xorm.Engine) error {
	var (
		colleagues = []Colleague{
			{Id: 1, Username: "system", Name: "系统管理员", Email: "system@email.com", Password: "1111", Enable: true},
			{Id: 2, Name: "Nike店长", Email: "nike@email.com", Password: "nike", Enable: true},
			{Id: 3, Name: "Nike店员", Email: "nike01@email.com", Password: "nike01", Enable: true},
			{Id: 4, Name: "Nike小时工", Email: "nike02@email.com", Password: "nike02", Enable: true},
			{Id: 5, Name: "Adidas店长", Email: "adidas@email.com", Password: "adidas", Enable: true},
			{Id: 6, Name: "上海专卖店长", Email: "shanghai@email.com", Password: "shanghai", Enable: true},
		}

		stores = []Store{
			{Id: 1, Code: "C001", Name: "北京Nike店", TenantCode: "hublabs", AddrProvince: "北京市", AddrCity: "北京市", AddrDistrict: "朝阳区", AddrDetail: "酒仙桥中路恒通商务园B37", Mobile: "13331056672", Tel: "1111111110", Enable: true},
			{Id: 2, Code: "C002", Name: "北京Adidas店", TenantCode: "hublabs", AddrProvince: "北京市", AddrCity: "北京市", AddrDistrict: "朝阳区", AddrDetail: "酒仙桥中路恒通商务园B37", Mobile: "17031056672", Tel: "2222222220", Enable: true},
			{Id: 3, Code: "C003", Name: "上海运动专卖店", TenantCode: "hublabs", AddrProvince: "上海市", AddrCity: "上海市", AddrDistrict: "浦东区", AddrDetail: "漕宝路23号", Mobile: "15031056672", Tel: "3333333330", Enable: true},
		}

		storeBrands = []StoreBrand{
			{Id: 1, StoreId: 1, BrandCode: "NK", Enable: true},
			{Id: 2, StoreId: 2, BrandCode: "AD", Enable: true},
			{Id: 3, StoreId: 3, BrandCode: "NK", Enable: true},
			{Id: 4, StoreId: 3, BrandCode: "AD", Enable: true},
		}

		storeColleagues = []StoreColleague{
			{Id: 1, ColleagueId: 1, StoreId: 1, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 2, ColleagueId: 1, StoreId: 2, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 3, ColleagueId: 1, StoreId: 3, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 4, ColleagueId: 2, StoreId: 1, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 5, ColleagueId: 3, StoreId: 1, StartDate: "", EndDate: "", Role: "member", Enable: true},
			{Id: 6, ColleagueId: 4, StoreId: 1, StartDate: "", EndDate: "", Role: "guest", Enable: true},
			{Id: 7, ColleagueId: 5, StoreId: 2, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 8, ColleagueId: 6, StoreId: 3, StartDate: "", EndDate: "", Role: "admin", Enable: true},
		}

		apps = []App{
			{Id: 1, Code: "O2O", Name: "在线抢单", TenantCode: "hublabs", Enable: true},
			{Id: 2, Code: "OHUB", Name: "在线结算", TenantCode: "hublabs", Enable: true},
		}

		colleagueApps = []ColleagueApp{
			{Id: 1, ColleagueId: 1, AppId: 1, Role: "admin", Enable: true},
			{Id: 2, ColleagueId: 1, AppId: 2, Role: "admin", Enable: true},
			{Id: 3, ColleagueId: 2, AppId: 1, Role: "admin", Enable: true},
			{Id: 4, ColleagueId: 2, AppId: 2, Role: "admin", Enable: true},
			{Id: 5, ColleagueId: 3, AppId: 1, Role: "member", Enable: true},
			{Id: 6, ColleagueId: 4, AppId: 1, Role: "guest", Enable: true},
			{Id: 7, ColleagueId: 5, AppId: 1, Role: "admin", Enable: true},
			{Id: 8, ColleagueId: 5, AppId: 2, Role: "admin", Enable: true},
			{Id: 9, ColleagueId: 6, AppId: 1, Role: "admin", Enable: true},
			{Id: 10, ColleagueId: 6, AppId: 2, Role: "admin", Enable: true},
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

	for _, u := range apps {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range colleagueApps {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	return nil
}
