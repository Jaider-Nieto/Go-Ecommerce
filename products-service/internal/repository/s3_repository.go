package repository

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jaider-nieto/ecommerce-go/products-service/internal/interfaces"
)

type S3Repository struct {
	s3Client *s3.S3
	bucket   string
}

func NewS3Repository(s3Client *s3.S3, bucket string) interfaces.S3RepositoryInterface {
	return &S3Repository{
		s3Client: s3Client,
		bucket:   bucket,
	}
}

func (r *S3Repository) UploadFile(file multipart.File) (string, error) {
	fileName := fmt.Sprintf("%d", time.Now().Unix())
	buffer, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	result, errAws := r.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer),
		ACL:    aws.String("public-read"),
		ContentType: aws.String("image/jpeg"),
	})
	if errAws != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}
	if result == nil {
		return "", fmt.Errorf("upload response is nil, file may not have been uploaded")
	}

	log.Println(result)

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.bucket, "us-east-2", fileName)

	return url, nil
}
