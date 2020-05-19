package tenants

import (
	"time"
)

type Tenant struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Enable    bool      `json:"-"`
	CreatedAt time.Time `json:"-" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
}
