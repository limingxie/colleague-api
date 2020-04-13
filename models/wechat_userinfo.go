package models

import (
	"time"
)

type WechatUserinfo struct {
	Id         int64  `json:"id"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"   xorm:"index"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	// Privilege       []string  `json:"privilege" xorm:"json"`
	ColleagueId     int64     `json:"colleagueId"   xorm:"index"`
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	CreatedAt       time.Time `json:"-" xorm:"created"`
	UpdatedAt       time.Time `json:"-" xorm:"updated"`
}
