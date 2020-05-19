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

func (store *StoreJsonView) GetBrands(ctx context.Context, tenantCode string) ([]tenants.Brand, error) {
	if store.Brands == nil {
		if err := store.loadBrands(ctx, tenantCode); err != nil {
			return nil, err
		}
	}

	return store.Brands, nil
}

func (store *StoreJsonView) loadBrands(ctx context.Context, tenantCode string) error {
	var brands []tenants.Brand
	if err := factory.DB(ctx).Select("brand.*").
		Table("store_brand").
		Join("INNER", "brand", "store_brand.brand_code = brand.code").
		Join("INNER", "store", "store.id = store_brand.store_id").
		Where("store.code = ? ", store.Code).
		And("store.tenant_code = ? ", tenantCode).
		And("brand.tenant_code = ? ", tenantCode).
		And("store_brand.enable = ? ", true).
		And("brand.enable = ? ", true).
		Find(&brands); err != nil {
		return err
	}

	store.Brands = brands

	return nil
}
