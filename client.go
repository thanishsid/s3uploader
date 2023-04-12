package s3uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Client struct {
	session *session.Session

	bucket  string
	baseDir string
}

type ClientConfig struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsRegion          string

	Bucket        string
	BaseDirectory string
}

func NewClient(c ClientConfig) (*Client, error) {
	conf := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AwsAccessKeyID, c.AwsSecretAccessKey, ""),
	}

	if c.AwsRegion != "" {
		conf.Region = aws.String(c.AwsRegion)
	}

	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}

	return &Client{
		session: sess,
		bucket:  c.Bucket,
		baseDir: c.BaseDirectory,
	}, nil
}
