package colleagues

import (
	"time"
)

type Store struct {
	Id         int64     `json:"id"`
	TenantCode string    `json:"tenantCode"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Enable     bool      `json:"enable"`
	CreatedAt  time.Time `json:"-" xorm:"created"`
	UpdatedAt  time.Time `json:"-" xorm:"updated"`
}

type StoreJsonView struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
