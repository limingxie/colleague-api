package colleagues

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/colleague-api/tenants"
)

type Colleague struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Mobile    string          `json:"mobile"`
	Password  string          `json:"password"`
	Enable    bool            `json:"enable"`
	CreatedAt time.Time       `json:"-" xorm:"created"`
	UpdatedAt time.Time       `json:"-" xorm:"updated"`
	Stores    []StoreJsonView `json:"-" xorm:"-"`
}

type ColleagueStoreJsonView struct {
	Id     int64           `json:"id"`
	Code   string          `json:"code"`
	Name   string          `json:"name"`
	Role   string          `json:"role"`
	Brands []tenants.Brand `json:"brands"`
}

func (Colleague) GetById(ctx context.Context, colleagueId int64) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).
		Where("colleague.id = ? ", colleagueId).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}
	return colleague, nil
}
