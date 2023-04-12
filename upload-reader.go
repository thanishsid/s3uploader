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

// Uploads data from provided reader and then returns the object key.
func (c *Client) UploadFromReader(data io.Reader, fileName string, contentType *string) (string, error) {
	uploader := s3manager.NewUploader(c.session)

	objectKey := path.Join(c.baseDir, (uuid.New().String() + filepath.Ext(fileName)))

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(c.bucket),
		Key:                aws.String(objectKey),
		Body:               data,
		ContentType:        contentType,
		ContentDisposition: aws.String(fmt.Sprintf(`inline; filename="%s"`, fileName)),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}
