package models

import (
	"context"
	"errors"
)

type Authenticator interface {
	Authentication(colleague Colleague, password string) (bool, error)
	GetColleague(ctx context.Context, identiKey string) (Colleague, error)
}

func GetAuthenticator(mode string) Authenticator {
	switch mode {
	case "email":
		return &EmailAuthentication{}
	case "mobile":
		return &MobileAuthentication{}
	case "wechat":
		return &WechatAuthentication{}
	default:
		return nil
	}
}

func Authentication(ctx context.Context, mode string, identiKey string, password string) (map[string]interface{}, error) {
	authenticator := GetAuthenticator(mode)

	colleague, err := authenticator.GetColleague(ctx, identiKey)
	if err != nil {
		return nil, err
	}

	if isPassed, err := authenticator.Authentication(colleague, password); err != nil {
		return nil, err
	} else if !isPassed {
		return nil, errors.New(`Authentication failed.`)
	}

	tokenDetail := make(map[string]interface{})

	tokenDetail["colleagueId"] = colleague.Id

	return tokenDetail, nil
}
