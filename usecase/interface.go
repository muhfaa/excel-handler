package usecase

import (
	"mime/multipart"
)

type FileUsecase interface {
	ReadExcelFile(file *multipart.File) ([][]string, error)
}
