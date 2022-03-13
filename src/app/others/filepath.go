package others

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func IconFilePathCheck(user_id string) bool {
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
	for _, item := range resp.Contents {
		log.Println("Name:         ", *item.Key)
		if *item.Key == key {
			log.Println("file true")
			return true
		}
	}
	return false
}
