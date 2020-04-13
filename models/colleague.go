package models

import (
	"context"
	"time"
)

type Colleague struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"password"`
	Enable    bool      `json:"enable"`
	CreatedAt time.Time `json:"-" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	Version   int       `json:"version" xorm:"version"`
}

func (Colleague) GetColleagueById(ctx context.Context, id int64) (Colleague, error) {
	return Colleague{Id: 1, Name: "xiao_ming", Enable: true}, nil
}
