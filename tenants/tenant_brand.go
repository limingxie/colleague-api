package tenants

import (
	"time"
)

type Brand struct {
	TenantCode string    `json:"tenantCode"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Enable     bool      `json:"-"`
	CreatedAt  time.Time `json:"-" xorm:"created"`
	UpdatedAt  time.Time `json:"-" xorm:"updated"`
}
