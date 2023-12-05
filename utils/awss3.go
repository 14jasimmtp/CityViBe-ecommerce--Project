package utils

// func CreateSession() *session.Session {
// 	sess := session.Must(session.NewSession(
// 		&aws.Config{
// 			Region: aws.String(cfg.Region),
// 			Credentials: credentials.NewStaticCredentials(
// 				cfg.AccessKeyID,
// 				cfg.AccessKeySecret,
// 				"",
// 			),
// 		},
// 	))
// 	return sess
// }

// func CreateS3Session(sess *session.Session) *s3.S3 {
// 	s3Session := s3.New(sess)
// 	return s3Session
// }

// func UploadImageToS3(file *multipart.FileHeader, sess *session.Session) (string, error) {

// 	image, err := file.Open()
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	defer image.Close()

// 	uploader := s3manager.NewUploader(sess)
// 	upload, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String("mobile-mart-image"),
// 		Key:    aws.String(file.Filename),
// 		Body:   image,
// 		ACL:    aws.String("public-read"),
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	return upload.Location, nil
// }
