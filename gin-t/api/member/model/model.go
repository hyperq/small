package model

import (
	"time"
)

// Member 用户struct
type Member struct {
	ID            int64     `gorm:"primary_key" json:"id"`
	Monicker      string    `gorm:"column:monicker" json:"monicker"`
	Name          string    `gorm:"column:name" json:"name"`
	Mobile        string    `gorm:"column:mobile" json:"mobile"`
	Salt          string    `gorm:"column:salt" json:"salt"`
	Password      string    `gorm:"column:password" json:"password"`
	GID           int64     `gorm:"column:gid" json:"gid"`
	UnionID       string    `gorm:"column:unionid" json:"unionid"`
	OpenidPublic  string    `gorm:"column:openid_public" json:"openid_public"`
	OpenidMiniapp string    `gorm:"column:openid_miniapp" json:"openid_miniapp"`
	OpenidApp     string    `gorm:"column:openid_app" json:"openid_app"`
	OpenidWeb     string    `gorm:"column:openid_web" json:"openid_web"`
	Sex           int64     `gorm:"column:sex" json:"sex"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	Bid           int64     `gorm:"column:bid" json:"bid"`
	Number        string    `gorm:"column:number" json:"number"`
	Addtime       time.Time `gorm:"column:addtime" json:"addtime"`
	Addip         string    `gorm:"column:addip" json:"addip"`
	Logintime     time.Time `gorm:"column:logintime" json:"logintime"`
	Loginip       string    `gorm:"column:loginip" json:"loginip"`
	Loginnum      int64     `gorm:"column:loginnum" json:"loginnum"`
	Attention     int64     `gorm:"column:attention" json:"attention"`
	Jointime      string    `gorm:"column:jointime" json:"jointime"`
	Order         string    `gorm:"column:order" json:"order"`
	Province      int64     `gorm:"column:province" json:"province"`
	City          int64     `gorm:"column:city" json:"city"`
	Area          int64     `gorm:"column:area" json:"area"`
	ProvinceCn    string    `gorm:"column:province_cn" json:"province_cn"`
	CityCn        string    `gorm:"column:city_cn" json:"city_cn"`
	AreaCn        string    `gorm:"column:area_cn" json:"area_cn"`
	Position      int64     `gorm:"column:position" json:"position"`
}
