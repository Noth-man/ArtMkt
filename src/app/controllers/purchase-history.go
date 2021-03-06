package controllers

import (
	"main/app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//購入履歴
func purchaseHistory(c *gin.Context) {
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	if loginbool == false {
		c.Redirect(302, "/loginform")
	} else {
		userid := models.GetUserID(UserInfo.UserName)
		products := models.GetProductIdFromPaymentHistory(userid)
		c.HTML(200, "purchaseHistory", gin.H{
			"title":     "Art-Mkt｜購入履歴",
			"login":     true,
			"username":  UserInfo.UserName,
			"csrfToken": csrf.GetToken(c),
			"products":  products,
		})
	}
}
