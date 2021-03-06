package controllers

import (
	"main/app/models"
	"main/config"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	csrf "github.com/utrack/gin-csrf"
)

func BuyerInformation(c *gin.Context) {
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.StripeAccount = session.Get("StripeAccount")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)

	//userid := models.GetUserID(UserInfo.UserId)
	if loginbool == true && models.CheckStripeAccountId(UserInfo.StripeAccount) == true {
		stripe.Key = config.Config.StripeKey
		productid := c.PostForm("productid")
		buyer_userid := models.GetUserIdWithPHT(productid)
		int_product_userid, _ := strconv.Atoi(buyer_userid)
		personal := models.GetPersonal(int_product_userid)
		c.HTML(200, "buyerInfo", gin.H{
			"title":          "BuyerInformation",
			"login":          true,
			"buyerUserId":    buyer_userid,
			"csrfToken":      csrf.GetToken(c),
			"kanji_f_name":   personal[2],
			"kanji_l_name":   personal[3],
			"kana_f_name":    personal[4],
			"kana_l_name":    personal[5],
			"postal_code":    personal[6],
			"address_level1": personal[7],
			"address_level2": personal[8],
			"address_line1":  personal[9],
			"address_line2":  personal[10],
			"organization":   personal[11],
			"username":       UserInfo.UserName,
			"productid":      productid,
		})
	} else {
		c.Redirect(302, "/loginform")
	}
}
