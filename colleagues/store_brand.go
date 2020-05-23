package colleagues

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/colleague-api/tenants"
)

type StoreBrand struct {
	Id        int64     `json:"id"`
	StoreId   int64     `json:"storeId"`
	BrandCode string    `json:"brandCode"`
	Enable    bool      `json:"enable"`
	CreatedAt time.Time `json:"-" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
}

func (Store) GetBrandsByStoreId(ctx context.Context, storeId int64) ([]tenants.Brand, error) {
	var brands []tenants.Brand
	if err := factory.DB(ctx).Select("brand.*").
		Table("store_brand").
		Join("INNER", "brand", "store_brand.brand_code = brand.code").
		Join("INNER", "store", "store.id = store_brand.store_id").
		Where("store.id = ? ", storeId).
		And("store.tenant_code = brand.tenant_code ").
		And("store_brand.enable = ? ", true).
		And("brand.enable = ? ", true).
		Find(&brands); err != nil {
		return nil, err
	}
	return brands, nil
}
