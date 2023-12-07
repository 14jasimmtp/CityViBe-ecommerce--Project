package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateSession() *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("AwsRegion")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AwsAccessKey"),
				os.Getenv("AwsSecretKey"),
				"",
			),
		},
	))
	return sess
}

func UploadImageToS3(file *multipart.FileHeader, sess *session.Session) (string, error) {
	image, err := file.Open()
	if err != nil {
		return "", err
	}
	// fmt.Println("**", sess)
	defer image.Close()
	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("vibecity1/product_images/"),
		Key:    aws.String(file.Filename),
		Body:   image,
		ACL:    aws.String("private"),
	})
	if err != nil {
		fmt.Println("eror")
		return "", err
	}
	return upload.Location, nil
}
