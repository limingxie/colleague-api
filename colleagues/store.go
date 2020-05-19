package colleagues

import (
	"time"

	"github.com/hublabs/colleague-api/tenants"
)

type Store struct {
	Id         int64           `json:"id"`
	TenantCode string          `json:"tenantCode"`
	Code       string          `json:"code"`
	Name       string          `json:"name"`
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
