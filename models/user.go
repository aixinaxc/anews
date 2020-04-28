package models

type User struct {
	CreatedAt    int    `xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt    int    `xorm:"default 0 comment('更新时间') INT(13)"`
	UserIcon     string `xorm:"comment('用户头像') VARCHAR(255)"`
	UserId       string `xorm:"not null pk comment('用户id') VARCHAR(64)"`
	UserNickname string `xorm:"comment('用户昵称') VARCHAR(255)"`
	UserState    int    `xorm:"default 1 comment('状态（1正常，2冻结，3删除）') INT(1)"`
	WxOpenid     string `xorm:"comment('微信openid') VARCHAR(64)"`
}
