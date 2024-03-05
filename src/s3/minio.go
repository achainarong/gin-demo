package s3

import (
	"gin-demo/src/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// TODO: Use minio to create files and do stuff so i can check the performance metrics with pyroscope
func GetMinioClient(settings *config.Settings) *s3.S3 {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(settings.MinioAccessKey, settings.MinioAccessSecret, ""),
		Endpoint:         aws.String(settings.MinioHost),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		panic(err)
	}
	return s3.New(newSession)
}
