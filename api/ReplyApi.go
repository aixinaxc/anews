package api

import (
	"fmt"
	"ginp/base"
	"ginp/models"
	"ginp/vo"
	"ginp/xrom_mysql"
	"github.com/gin-gonic/gin"
	"time"
)
//评论列表
func ReplyList(context *gin.Context) {
	articleId := context.Query("articleId")
	replyList := make([]vo.ReplyListVo, 0)
	sql := "SELECT * FROM reply AS r LEFT JOIN user AS u ON r.user_id = u.user_id WHERE r.article_id = ? ORDER BY r.created_at DESC"
	xrom_mysql.DataEngine.SQL(sql,articleId).Find(&replyList)
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, 0, replyList))
}

//评论保存
func ReplySave(context *gin.Context) {
	var reply models.Reply
	_ = context.ShouldBind(&reply)
	reply.ReplyId = base.UniqueId()
	reply.ReplyState = 1
	reply.CreatedAt = int(time.Now().Unix())
	_,err := xrom_mysql.DataEngine.Insert(reply)
	if err != nil {
		fmt.Println("sort_save:",err)
	}
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, 0, nil))
}

