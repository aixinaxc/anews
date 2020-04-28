package api

import (
	"encoding/json"
	"fmt"
	"ginp/base"
	"ginp/models"
	"ginp/vo"
	"ginp/xrom_mysql"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func Login(context *gin.Context)  {
	code  := context.Query("code")
	if code == ""{
		context.JSON(200, base.RetunMsgFunc(base.CodeDataError, 0, nil))
		return
	}
	url := `https://api.weixin.qq.com/sns/jscode2session?appid=` + base.AppId + `&secret=` + base.AppSecret + `&js_code=` + code + `&grant_type=authorization_code`
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("code请求：",err,resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var wx vo.WX
	json.Unmarshal(body,&wx)
	var user models.User
	if wx.Openid != "" {
		xrom_mysql.DataEngine.Where("wx_openid = ?", wx.Openid).Get(&user)
		fmt.Println(user)
		if user.UserId == "" {
			user.UserId = base.UniqueId()
			user.WxOpenid = wx.Openid
			user.UserState = 1
			user.CreatedAt = int(time.Now().Unix())

			_,err := xrom_mysql.DataEngine.Insert(user)
			if err != nil {
				fmt.Println("sort_save:",err)
			}
		}
	}
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, 0, user))
}

func UserSave(context *gin.Context)  {
	var user models.User
	context.ShouldBind(&user)
	_,err := xrom_mysql.DataEngine.Where("user_id = ?",user.UserId).Update(user)
	if err != nil {
		fmt.Println("sort_save:",err)
	}
	context.JSON(200, base.RetunMsgFunc(base.CodeDataSuccess, 0, user))
}