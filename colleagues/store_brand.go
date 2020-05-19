package colleagues

import (
	"time"
)

type StoreBrand struct {
	Id        int64     `json:"id"`
	StoreId   int64     `json:"storeId"`
	BrandCode string    `json:"brandCode"`
	Enable    bool      `json:"enable"`
	CreatedAt time.Time `json:"-" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
}
