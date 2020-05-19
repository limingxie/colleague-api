package colleagues

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
)

type StoreColleague struct {
	Id          int64     `json:"id"`
	StoreId     int64     `json:"storeId"`
	ColleagueId int64     `json:"colleagueId"`
	StartDate   string    `json:"startDate"`
	EndDate     string    `json:"endDate"`
	Role        string    `json:"role"`
	Enable      bool      `json:"enable"`
	CreatedAt   time.Time `json:"-" xorm:"created"`
	UpdatedAt   time.Time `json:"-" xorm:"updated"`
}

func (colleague *Colleague) GetStoreAndRoles(ctx context.Context, tenantCode string) ([]StoreJsonView, error) {
	if colleague.Stores == nil {
		if err := colleague.loadStoreAndRoles(ctx, tenantCode); err != nil {
			return nil, err
		}
	}

	return colleague.Stores, nil
}

func (colleague *Colleague) loadStoreAndRoles(ctx context.Context, tenantCode string) error {
	var storeJsonViews []StoreJsonView
	if err := factory.DB(ctx).Select("store.id, store.code, store.name, store_colleague.role").
		Table("store_colleague").
		Join("INNER", "store", "store.id = store_colleague.store_id").
		Where("store_colleague.colleague_id = ? ", colleague.Id).
		And("store_colleague.enable = ? ", true).
		And("store.tenant_code = ? ", tenantCode).
		And("store.enable = ? ", true).
		Find(&storeJsonViews); err != nil {
		return err
	}

	if len(storeJsonViews) > 0 {
		for i := range storeJsonViews {
			brands, err := storeJsonViews[i].GetBrands(ctx, tenantCode)
			if err != nil {
				return err
			}
			storeJsonViews[i].Brands = brands
		}
	}

	colleague.Stores = storeJsonViews

	return nil
}
