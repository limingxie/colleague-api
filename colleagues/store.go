package colleagues

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/colleague-api/tenants"
)

type Store struct {
	Id         int64           `json:"id"`
	TenantCode string          `json:"tenantCode"`
	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Province   string          `json:"province"`
	City       string          `json:"city"`
	District   string          `json:"district"`
	Detail     string          `json:"detail"`
	Enable     bool            `json:"enable"`
	CreatedAt  time.Time       `json:"-" xorm:"created"`
	UpdatedAt  time.Time       `json:"-" xorm:"updated"`
	Brands     []tenants.Brand `json:"-" xorm:"-"`
}

type StoreJsonView struct {
	Code   string          `json:"code"`
	Name   string          `json:"name"`
	Role   string          `json:"role"`
	Brands []tenants.Brand `json:"brands"`
}

func (Store) GetStoreAddressById(ctx context.Context, storeId int64) (map[string]interface{}, error) {
	var store Store
	if _, err := factory.DB(ctx).
		Where("store.id = ? ", storeId).
		And("store.enable = ? ", true).
		Get(&store); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":       store.Id,
		"province": store.Province,
		"city":     store.City,
		"district": store.District,
		"detail":   store.Detail,
	}, nil
}
