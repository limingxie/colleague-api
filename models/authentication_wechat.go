package models

import (
	"context"
	"errors"

	"github.com/hublabs/colleague-api/factory"
)

type WechatAuthentication struct{}

func (WechatAuthentication) Authentication(colleague Colleague, password string) (bool, error) {
	if colleague.Id == 0 {
		return false, errors.New("invalid wechat account.")
	}

	return true, nil
}

func (WechatAuthentication) GetColleague(ctx context.Context, identiKey string) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).Select("colleague.*").
		Table("colleague").
		Join("INNER", "wechat_userinfo", "wechat_userinfo.colleague_id = colleague.id").
		Where("wechat_userinfo.unionid = ? ", identiKey).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}
	return colleague, nil
}
