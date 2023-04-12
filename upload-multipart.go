package s3uploader

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// Uploads data from provided reader and then returns the object key.
func (c *Client) UploadMultipartFile(h *multipart.FileHeader) (string, error) {
	if h == nil {
		return "", errors.New("multipart file header is nil")
	}

	data, err := h.Open()

	if err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(c.session)

	objectKey := path.Join(c.baseDir, (uuid.New().String() + filepath.Ext(h.Filename)))

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(c.bucket),
		Key:                aws.String(objectKey),
		Body:               data,
		ContentType:        aws.String(h.Header.Get("Content-Type")),
		ContentDisposition: aws.String(fmt.Sprintf(`inline; filename="%s"`, h.Filename)),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}
