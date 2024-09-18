package usecase

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

type FileUsecaseImpl struct{}

func NewFileUsecase() *FileUsecaseImpl {
	return &FileUsecaseImpl{}
}

func (u FileUsecaseImpl) ReadExcelFile(file *multipart.File) ([][]string, error) {
	logCtx := "ReadExcelFile"
	// Read file content
	fileBytes, err := ioutil.ReadAll(*file)
	if err != nil {
		log.Printf("%v, ERROR: failed to read file: %v", logCtx, err)
		return nil, err
	}

	f, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		log.Printf("%v, ERROR: failed to opening file: %v", logCtx, err)
		return nil, err
	}

	// Assume we are working with the first sheet
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		log.Printf("%v, ERROR: sheet not found", logCtx)
		return nil, fmt.Errorf("sheet not found")
	}

	// Read rows from the sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Printf("%v, ERROR: failed getting rows: %v", logCtx, err)
		return nil, err
	}

	var data [][]string
	for _, row := range rows {
		var rowData []string
		for _, cell := range row {
			rowData = append(rowData, cell)
		}
		data = append(data, rowData)
	}

	return data, nil
}
