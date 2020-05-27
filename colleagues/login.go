package colleagues

import (
	"context"
	"errors"

	"github.com/hublabs/colleague-api/factory"
)

type Login struct{}

func (Login) GetColleagueByLoginModeAndCheckPasswork(ctx context.Context, mode string, identiKey string, password string) (Colleague, error) {
	switch mode {
	case "email":
		return Login{}.GetColleagueByEmail(ctx, identiKey, password)
	case "mobile":
		return Login{}.GetColleagueByMobile(ctx, identiKey, password)
	case "wechat":
		return Login{}.GetColleagueByUnionid(ctx, identiKey)
	default:
		return Colleague{}, nil
	}
}

func (Login) GetTokenDetail(ctx context.Context, mode string, identiKey string, password string) (map[string]interface{}, error) {
	colleague, err := Login{}.GetColleagueByLoginModeAndCheckPasswork(ctx, mode, identiKey, password)
	if err != nil {
		return nil, err
	}

	if colleague.Id == 0 {
		return nil, errors.New("login failed.")
	}

	tokenDetail := make(map[string]interface{})

	tokenDetail["colleagueId"] = colleague.Id

	return tokenDetail, nil
}

func (Login) GetColleagueByEmail(ctx context.Context, identiKey string, password string) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).
		Where("colleague.email = ? ", identiKey).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}

	if colleague.Id == 0 {
		return Colleague{}, errors.New("invalid email.")
	}
	if colleague.Password == "" || password == "" {
		return Colleague{}, errors.New("Refuse to login with account password.")
	} else if colleague.Password != password {
		return Colleague{}, errors.New("invalid password.")
	}

	return colleague, nil
}

func (Login) GetColleagueByMobile(ctx context.Context, identiKey string, password string) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).
		Where("colleague.mobile = ? ", identiKey).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}

	if colleague.Password == "" || password == "" {
		return Colleague{}, errors.New("Refuse to login with account password.")
	} else if colleague.Password != password {
		return Colleague{}, errors.New("invalid password.")
	}

	if colleague.Id == 0 {
		return Colleague{}, errors.New("invalid mobile.")
	}

	return colleague, nil
}

func (Login) GetColleagueByUnionid(ctx context.Context, identiKey string) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).Select("colleague.*").
		Table("colleague").
		Join("INNER", "wechat_userinfo", "wechat_userinfo.colleague_id = colleague.id").
		Where("wechat_userinfo.unionid = ? ", identiKey).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}

	if colleague.Id == 0 {
		return Colleague{}, errors.New("invalid wechat account.")
	}

	return colleague, nil
}

func (Login) GetColleagueInfos(ctx context.Context, tenantCode string, colleagueId int64) (map[string]interface{}, error) {
	colleague, err := Colleague{}.GetById(ctx, colleagueId)
	if err != nil {
		return nil, err
	}

	if colleague.Id == 0 {
		return nil, err
	}

	stores, err := colleague.GetStoreAndRoles(ctx, tenantCode)
	if err != nil {
		return nil, err
	}

	apps, err := colleague.GetAppAndRoles(ctx, tenantCode)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["tenantCode"] = tenantCode
	result["id"] = colleague.Id
	result["name"] = colleague.Name
	result["username"] = colleague.Username
	result["stores"] = stores
	result["apps"] = apps

	return result, nil

}
