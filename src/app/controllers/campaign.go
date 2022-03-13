package controllers

import (
	"log"
	"main/app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func campaign(c *gin.Context) {
	Number := c.Param("number")
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	log.Println(Number)
	if Number == "1" {
		c.HTML(200, "campaign1", gin.H{
			"title":    "簡単出品ガイド",
			"username": UserInfo.UserName,
			"login":    loginbool,
		})
	} else if Number == "2" {
		c.HTML(200, "campaign2", gin.H{
			"title":    "artmktとは？",
			"username": UserInfo.UserName,
			"login":    loginbool,
		})
	} else if Number == "3" {
		c.HTML(200, "campaign3", gin.H{
			"title":    "売上キャンペーン",
			"username": UserInfo.UserName,
			"login":    loginbool,
		})
	}
	c.Redirect(302, "/")
}
