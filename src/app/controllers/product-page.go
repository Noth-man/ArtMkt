package controllers

import (
	"log"
	"main/app/models"
	"math"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//-------------------------------------------------- ProductPage --------------------------------------------------
func ProductPage(c *gin.Context) {
	productNumber := c.Param("number")
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")
	UserInfo.logintoken = session.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	product := models.GetProduct(productNumber)
	log.Println(product)
	username := models.GetUserName(product[1])
	f, err := strconv.ParseFloat(product[6], 64)
	if err != nil {
		log.Println(err)
	}
	f = f * 1.1
	taxamount := int(math.Round(f))
	log.Println(taxamount)
	title := "Art-Mkt｜" + product[4]
	if loginbool == false {
		c.HTML(200, "product", gin.H{
			"title":           title,
			"login":           false,
			"csrfToken":       csrf.GetToken(c),
			"ProductId":       product[0],
			"ProductUsername": username,
			"UserId":          product[1],
			"StripeProductId": product[2],
			"StripePriceId":   product[3],
			"ItemName":        product[4],
			"Description":     product[5],
			"username":        UserInfo.UserName,
			"Amount":          taxamount,
			"SoldOut":         product[7],
			"Category":        product[8],
		})
	} else {
		userid := models.GetUserID(UserInfo.UserName)
		if models.CheckCart(userid, product[0]) == true {
			c.HTML(200, "product", gin.H{
				"title":           title,
				"login":           true,
				"username":        UserInfo.UserName,
				"csrfToken":       csrf.GetToken(c),
				"ProductId":       product[0],
				"ProductUsername": username,
				"UserId":          product[1],
				"StripeProductId": product[2],
				"StripePriceId":   product[3],
				"ItemName":        product[4],
				"Description":     product[5],
				"Amount":          taxamount,
				"cart":            true,
				"SoldOut":         product[7],
				"Category":        product[8],
			})
		} else {
			c.HTML(200, "product", gin.H{
				"title":           title,
				"login":           true,
				"username":        UserInfo.UserName,
				"csrfToken":       csrf.GetToken(c),
				"ProductId":       product[0],
				"ProductUsername": username,
				"UserId":          product[1],
				"StripeProductId": product[2],
				"StripePriceId":   product[3],
				"ItemName":        product[4],
				"Description":     product[5],
				"Amount":          taxamount,
				"cart":            false,
				"SoldOut":         product[7],
				"Category":        product[8],
			})
		}

	}
}

func ProductImage(c *gin.Context) {
	productNumber := c.Param("number")
	session := sessions.Default(c)
	UserInfo.UserName = session.Get("UserName")

	c.HTML(200, "image", gin.H{
		"productid": productNumber,
	})
}

func UserProductPage(c *gin.Context) {
	user_id := c.Param("number")
	int_user_id, _ := strconv.Atoi(user_id)
	sessiond := sessions.Default(c)
	UserInfo.UserName = sessiond.Get("UserName")
	UserInfo.logintoken = sessiond.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)
	products := models.GetAllProductOfUserId(user_id)
	product_username := models.GetUserName(user_id)
	title := product_username + "さんのアイテム"
	count_follower := models.CountFollower(int_user_id)
	my_user_id := models.GetUserID(UserInfo.UserName)
	count_product := models.CountProduct(int_user_id)
	//filepathbool := others.IconFilePathCheck(user_id)
	self_introduction := models.GetSelfIntroduction(int_user_id)

	key := "icon/" + "userid" + user_id + "icon.jpg"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		log.Println("session.NewSession err = ", err)
	}
	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("artmkt")})
	if err != nil {
		log.Println("svc.ListObjects err = ", err)
	}
	var filepathbool bool
	for _, item := range resp.Contents {
		log.Println("Name:         ", *item.Key)
		if *item.Key == key {
			filepathbool = true
		}
	}

	log.Println(filepathbool)
	if loginbool == false {
		/*
			c.HTML(200, "UserProductPage", gin.H{
				"title":            title,
				"login":            false,
				"csrfToken":        csrf.GetToken(c),
				"products":         products,
				"productUserId":    user_id,
				"productUsername":  product_username,
				"countFollower":    count_follower,
				"countProduct":     count_product,
				"filepath":         filepathbool,
				"SelfIntroduction": self_introduction,
			})*/
		c.Redirect(302, "/loginform")
	} else {
		if models.CheckFollow(my_user_id, int_user_id) == "なし" {
			c.HTML(200, "UserProductPage", gin.H{
				"title":            title,
				"login":            true,
				"username":         UserInfo.UserName,
				"csrfToken":        csrf.GetToken(c),
				"products":         products,
				"productUsername":  product_username,
				"productUserId":    user_id,
				"countFollower":    count_follower,
				"follow":           false,
				"countProduct":     count_product,
				"filepath":         filepathbool,
				"SelfIntroduction": self_introduction,
			})
		} else {
			c.HTML(200, "UserProductPage", gin.H{
				"title":            title,
				"login":            true,
				"username":         UserInfo.UserName,
				"csrfToken":        csrf.GetToken(c),
				"products":         products,
				"productUsername":  product_username,
				"productUserId":    user_id,
				"countFollower":    count_follower,
				"countProduct":     count_product,
				"follow":           true,
				"filepath":         filepathbool,
				"SelfIntroduction": self_introduction,
			})
		}

	}

}
