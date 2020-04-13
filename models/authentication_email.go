package models

import (
	"context"
	"errors"

	"github.com/hublabs/colleague-api/factory"
)

type EmailAuthentication struct{}

func (EmailAuthentication) Authentication(colleague Colleague, password string) (bool, error) {
	if colleague.Id == 0 {
		return false, errors.New("invalid email.")
	}
	if colleague.Password == "" || password == "" {
		return false, errors.New("Refuse to login with account password.")
	} else if colleague.Password != password {
		return false, errors.New("invalid password.")
	}

	return true, nil
}

func (EmailAuthentication) GetColleague(ctx context.Context, identiKey string) (Colleague, error) {
	var colleague Colleague
	if _, err := factory.DB(ctx).
		Where("colleague.email = ? ", identiKey).
		And("colleague.enable = ? ", true).
		Get(&colleague); err != nil {
		return Colleague{}, err
	}

	return colleague, nil
}
