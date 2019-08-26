package model

import "time"

// User 用户列表
type User struct {
	ID              int       `gorm:"primary_key" json:"id"` //
	Attention       int       `json:"attention"`             //关注公众号
	Nickname        string    `json:"nickname"`              //昵称
	Avatar          string    `json:"avatar"`                //头像
	PromotionID     int       `json:"promotion_id"`          //推广人id
	WxOpenidPublic  string    `json:"wx_openid_public"`      //微信公众号openid
	WxUnionid       string    `json:"wx_unionid"`            //微信unionid
	Sex             int       `json:"sex"`                   //性别  1时是男性，值为2时是女性，值为0时是未知
	WxProvince      string    `json:"wx_province"`           //微信  省份
	WxCity          string    `json:"wx_city"`               //微信  市
	WxCountry       string    `json:"wx_country"`            //微信 国家
	Beans           int       `json:"beans"`                 //豆子
	Phonenumber     string    `json:"phonenumber"`           //手机号码
	Password        string    `json:"password"`              //密码
	Salt            string    `json:"salt"`                  //
	WxOpenidMiniapp string    `json:"wx_openid_miniapp"`     //微信小程序openid
	WxOpenidApp     string    `json:"wx_openid_app"`         //微信appopenid
	WxOpenidWeb     string    `json:"wx_openid_web"`         //微信网页端openid
	Addtime         time.Time `json:"addtime"`               //注册时间
	Addip           string    `json:"addip"`                 //注册ip
	Logintime       time.Time `json:"logintime"`             //上次登入时间
	Loginip         string    `json:"loginip"`               //上次登录ip
	Logintimes      int       `json:"logintimes"`            //登录次数
}
