package vo

import "ginp/models"

type ReplyListVo struct {
	models.Reply `xorm:"extends"`
	CreatedAt    int    `xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt    int    `xorm:"default 0 comment('更新时间') INT(13)"`
	models.User `xorm:"extends"`
}
