package tenants

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
)

type Brand struct {
	TenantCode string    `json:"tenantCode"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Enable     bool      `json:"-"`
	CreatedAt  time.Time `json:"-" xorm:"created"`
	UpdatedAt  time.Time `json:"-" xorm:"updated"`
}

func (Brand) GetBrandsByTenantCodeAndCodes(ctx context.Context, tenantCode string, brandCodes []string) ([]Brand, error) {
	var brands []Brand
	if err := factory.DB(ctx).Select("brand.tenant_code, brand.code, brand.name").
		Table("brand").
		Where("brand.tenant_code = ? ", tenantCode).
		In("brand.code", brandCodes).
		And("brand.enable = ? ", true).
		Find(&brands); err != nil {
		return nil, err
	}

	return brands, nil
}
