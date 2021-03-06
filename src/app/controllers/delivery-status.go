package controllers

import (
	"log"
	"main/app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SippingSuccess(c *gin.Context) {
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.StripeAccount = session.Get("StripeAccount")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	productid := c.PostForm("productid")
	if loginbool == true && models.CheckDeliveryStatusProductId(productid) == "なし" {
		models.InsertSipping(productid)
		log.Println("なし")
		c.Redirect(302, "/registered-items")
	} else if loginbool == true && models.CheckDeliveryStatusProductId(productid) == "あり" {
		log.Println("あり")
		c.Redirect(302, "/registered-items")
	} else {
		c.Redirect(302, "/loginform")
	}
}

func ArrivalSuccess(c *gin.Context) {
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.StripeAccount = session.Get("StripeAccount")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	productid := c.PostForm("productid")
	if loginbool == true && models.CheckDeliveryStatusProductId(productid) == "あり" && models.CheckArrives(productid) != "1" {
		log.Println("あり")
		models.UpdateArrives(productid)
		/*
			transferGroup := models.GetTransferGroup(productid)
			product := models.GetProduct(productid)
			amount, _ := strconv.ParseFloat(product[6], 64)
			payamount := int64(math.Round(amount * 0.78))
			log.Println(payamount)
			productuserid, _ := strconv.Atoi(product[1])
			stripe_account, _ := models.GetStripeAccountId(productuserid)

			stripe.Key = config.Config.StripeKey

			transferParams := &stripe.TransferParams{
				Amount:        stripe.Int64(payamount),
				Currency:      stripe.String(string(stripe.CurrencyJPY)),
				Destination:   stripe.String(stripe_account),
				TransferGroup: stripe.String(transferGroup),
			}
			_, err := transfer.New(transferParams)
			if err != nil {
				log.Println(err)
			}
		*/
		c.Redirect(302, "/purchase-history")
	} else if loginbool == true && models.CheckDeliveryStatusProductId(productid) == "なし" {
		c.Redirect(302, "/purchase-history")

	} else {
		c.Redirect(302, "/loginform")

	}
}
