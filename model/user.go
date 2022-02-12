package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	// User Information
	Id       int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// Salted random string
	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`
	Online int    `xorm:"int(10)" form:"online" json:"online"`
	// Front-end token
	Token string `xorm:"varchar(40)" form:"token" json:"token"`
	Memo  string `xorm:"varchar(140)" form:"memo" json:"memo"`
	// calc user amount
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
