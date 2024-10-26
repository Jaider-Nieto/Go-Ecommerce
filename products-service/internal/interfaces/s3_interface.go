package interfaces

import "mime/multipart"

type S3RepositoryInterface interface {
	UploadFile(file multipart.File) (string, error)
}
