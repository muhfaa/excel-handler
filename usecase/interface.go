package usecase

import (
	"io"
)

type FileUsecase interface {
	ReadExcelFile(file io.Reader) ([][]string, error)
}
