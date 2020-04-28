package models

type Article struct {
	ArticleId          string `xorm:"not null pk comment('文章id') VARCHAR(64)"`
	UserId             string `xorm:"comment('用户id') VARCHAR(64)"`
	ArticleTitle       string `xorm:"comment('文章标题') VARCHAR(255)"`
	ArticleBrowse      int    `xorm:"default 0 comment('浏览数') INT(13)"`
	ArticleIcon        string `xorm:"comment('文章图标') VARCHAR(255)"`
	ArticleSort        string `xorm:"comment('文章类型') VARCHAR(255)"`
	ArticleKeyword     string `xorm:"comment('文章关键词') VARCHAR(255)"`
	ArticleEditContent string `xorm:"comment('文章编辑内容') TEXT"`
	ArticleShowContent string `xorm:"comment('文章展示内容') TEXT"`
	ArticleState       int    `xorm:"default 1 comment('文章状态（1正常）') INT(1)"`
	CreatedAt          int    `xorm:"default 0 comment('创建时间') INT(13)"`
	UpdatedAt          int    `xorm:"default 0 comment('修改时间') INT(13)"`
}
