package s3

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "gin-demo/src/config"
)

func NewMinioClient(settings *config.Settings) *s3.S3 {
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

