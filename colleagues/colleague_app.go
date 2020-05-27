package colleagues

import (
	"context"
	"time"

	"github.com/hublabs/colleague-api/factory"
)

type ColleagueApp struct {
	Id          int64     `json:"id"`
	AppId       int64     `json:"storeId"`
	ColleagueId int64     `json:"colleagueId"`
	Role        string    `json:"role"`
	Enable      bool      `json:"enable"`
	CreatedAt   time.Time `json:"-" xorm:"created"`
	UpdatedAt   time.Time `json:"-" xorm:"updated"`
}

func (colleague *Colleague) GetAppAndRoles(ctx context.Context, tenantCode string) ([]AppJsonView, error) {
	var appJsonViews []AppJsonView
	if err := factory.DB(ctx).Select("app.id, app.code, app.name, colleague_app.role").
		Table("colleague_app").
		Join("INNER", "app", "app.id = colleague_app.app_id").
		Where("colleague_app.colleague_id = ? ", colleague.Id).
		And("colleague_app.enable = ? ", true).
		And("app.tenant_code = ? ", tenantCode).
		And("app.enable = ? ", true).
		Find(&appJsonViews); err != nil {
		return nil, err
	}

	return appJsonViews, nil
}
