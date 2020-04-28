package api

import (
	"fmt"
	"ginp/base"
	"ginp/models"
	"ginp/xrom_mysql"
	"github.com/gin-gonic/gin"
)

//文章列表
func ArticleList(context *gin.Context) {
	pageNum := context.Query("page_num")
	pageSize := context.Query("page_size")
	pS,of := base.Offer(pageNum,pageSize)
	articles := make([]models.Article, 0)
	err := xrom_mysql.DataEngine.Cols("article_id", "article_sort", "article_browse", "article_keyword", "article_title", "article_icon",
		"created_at", "updated_at").Desc("created_at").Limit(pS,of).Find(&articles)
	articlesC := new(models.Article)
	total, err := xrom_mysql.DataEngine.Count(articlesC)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(articles)
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, total, articles))
}

//文章详情
func ArticleDetail(context *gin.Context) {
	articleId := context.Query("articleId")
	fmt.Println("articleId", articleId)
	article := new(models.Article)
	xrom_mysql.DataEngine.Where("article_id = ?", articleId).Get(article)
	article.ArticleBrowse = article.ArticleBrowse + 1
	xrom_mysql.DataEngine.Where("article_id = ?",article.ArticleId).Update(article)
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, 0, article))
}
