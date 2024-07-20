package config

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
    "os"
)

type AWSConfig struct {
    Region          string
    AccessKeyID      string
    SecretAccessKey  string
    S3BucketName     string
}

func LoadAWSConfig() *AWSConfig {
    return &AWSConfig{
        Region:          getEnv("AWS_REGION", "us-west-2"),
        AccessKeyID:      getEnv("AWS_ACCESS_KEY_ID", ""),
        SecretAccessKey:  getEnv("AWS_SECRET_ACCESS_KEY", ""),
        S3BucketName:     getEnv("S3_BUCKET_NAME", "my-bucket"),
    }
}

func NewS3Session(cfg *AWSConfig) *s3.S3 {
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(cfg.Region),
        Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
    })
    if err != nil {
        log.Fatal("Failed to create AWS session: ", err)
    }
    return s3.New(sess)
}
