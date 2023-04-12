package s3uploader

import (
	"fmt"
	"io"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type UploadPayload struct {
	File        io.Reader
	Filename    string
	ContentType string
}

// Uploads data from provided upload payload and then returns the object key.
func (c *Client) Upload(p UploadPayload) (string, error) {
	uploader := s3manager.NewUploader(c.session)

	objectKey := path.Join(c.baseDir, (uuid.New().String() + filepath.Ext(p.Filename)))

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(c.bucket),
		Key:                aws.String(objectKey),
		Body:               p.File,
		ContentType:        aws.String(p.ContentType),
		ContentDisposition: aws.String(fmt.Sprintf(`inline; filename="%s"`, p.Filename)),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}
