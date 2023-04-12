package s3uploader

import (
	"bytes"
	"encoding/base64"
	"path"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

var b64PrefixReg = regexp.MustCompile(`(?:data:)(.*)(?:base64,)`)

// Decodes the provided base64 string and uploads it to s3 and then returns the object key.
func (c *Client) UploadBase64(b64File string) (string, error) {
	prefix := b64PrefixReg.FindString(b64File)

	dataString := strings.TrimPrefix(b64File, prefix)

	if err := validation.Validate(dataString, is.Base64); err != nil {
		return "", err
	}

	uploader := s3manager.NewUploader(c.session)

	objectKey := path.Join(c.baseDir, uuid.New().String())

	data, err := base64.StdEncoding.DecodeString(dataString)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(data)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(objectKey),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}
