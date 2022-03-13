package controllers

import (
	"log"
	"main/app/models"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//マイページ
func mypage(c *gin.Context) {
	uname := c.Param("username")
	sessiond := sessions.Default(c)
	UserInfo.UserName = sessiond.Get("UserName")
	UserInfo.logintoken = sessiond.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)

	if UserInfo.UserName == uname && loginbool == true {
		userid := models.GetUserID(UserInfo.UserName)
		Self_Introduction := models.GetSelfIntroduction(userid)
		email := models.GetUserEmail(userid)
		struserid := strconv.Itoa(userid)

		//filepathbool := others.IconFilePathCheck(struserid)

		key := "icon/" + "userid" + struserid + "icon.jpg"

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
				log.Println("file true")
			}
		}

		c.HTML(200, "mypage", gin.H{
			"title":            "Art-Mkt｜マイページ",
			"login":            true,
			"username":         UserInfo.UserName,
			"userid":           userid,
			"email":            email,
			"csrfToken":        csrf.GetToken(c),
			"SelfIntroduction": Self_Introduction,
			"filepath":         filepathbool,
		})

	} else {
		c.Redirect(302, "/")
	}
}

func mypageDetail(c *gin.Context) {
	sessiond := sessions.Default(c)
	UserInfo.UserName = sessiond.Get("UserName")
	UserInfo.logintoken = sessiond.Get("logintoken")
	loginbool := models.LoginTokenCheck(UserInfo.UserName, UserInfo.logintoken)

	if loginbool == true {
		userid := models.GetUserID(UserInfo.UserName)
		self_introduction := c.PostForm("self-introduction")
		models.SelfIntroductionRegistration(userid, self_introduction)
		struserid := strconv.Itoa(userid)
		username := models.GetUserName(struserid)
		redirecturl := "/mypage/" + username

		file, header, err := c.Request.FormFile("image")
		if err != nil {
			log.Println("c.Request.FormFile err = ", err)
			c.Redirect(302, redirecturl)
			return
		}

		filename := header.Filename
		log.Println("filename =", filename)
		pos := strings.LastIndex(filename, ".")
		log.Println("pos =", filename[pos:])
		/*
			create_path := "app/static/img/icon/userid" + struserid + "/" + filename
			mkdir_path := "app/static/img/icon/userid" + struserid

			// mkdir
			iconpath := others.IconFilePathCheck(struserid)
			if iconpath == false {
				err = os.Mkdir(mkdir_path, 0755)
				if err != nil {
					log.Println("os.Mkdir err = ", err)
				}
			}

			out, err := os.Create(create_path)
			if err != nil {
				log.Println("os.create err = ", err)
			}

			//file copy
			_, err = io.Copy(out, file)
			if err != nil {
				log.Println("io.copy err = ", err)
			}

			defer out.Close()

			newpathjpg := "app/static/img/icon/userid" + struserid + "/" + "userid" + struserid + "icon.jpg"

			if filename[pos:] == ".png" {
				if err := os.Rename(create_path, newpathjpg); err != nil {
					log.Println(err)
				}
			} else {
				if err := os.Rename(create_path, newpathjpg); err != nil {
					log.Println(err)
				}
			}

			imagefile, err := os.Open(newpathjpg)
			if err != nil {
				log.Println("os.Open err = ", err)
			}
			defer imagefile.Close()
		*/
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-northeast-1")},
		)
		if err != nil {
			log.Println("session.NewSession err = ", err)
		}
		svc := s3.New(sess)

		key := "icon/" + "userid" + struserid + "icon.jpg"

		resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String("artmkt")})
		if err != nil {
			log.Println("svc.ListObjects err = ", err)
		}
		for _, item := range resp.Contents {
			log.Println("Name:         ", *item.Key)
			log.Println("Last modified:", *item.LastModified)
		}

		uploader := s3manager.NewUploader(sess)
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("artmkt"),
			Key:    aws.String(key),
			Body:   file,
		})
		if err != nil {
			log.Println("uploader err =", err)
		}
		/*
			err = os.Remove(mkdir_path)
			if err != nil {
				log.Println("remove err =", err)
			}
		*/
		c.Redirect(302, redirecturl)

	} else {
		c.Redirect(302, "/")
	}
}
