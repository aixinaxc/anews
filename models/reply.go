package models

type Reply struct {
	ReplyId      string `xorm:"not null pk comment('回复id') VARCHAR(64)"`
	UserId       string `xorm:"comment('用户id') VARCHAR(64)"`
	ArticleId    string `xorm:"comment('文章id') VARCHAR(64)"`
	ReplyContent string `xorm:"comment('回复内容') TEXT"`
	ReplyState   int    `xorm:"default 1 comment('回复状态（1正常，2冻结，3删除）') INT(1)"`
	CreatedAt    int    `xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt    int    `xorm:"default 0 comment('更新时间') INT(13)"`
}
