package handler

import (
	"excel-handler/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	FileUsecase usecase.FileUsecase
}

// UploadFile handles the Excel file upload and processing
// Method: POST
// URI: localhost:8080/upload
func (h *Handler) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if fileHeader.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Please upload an Excel file."})
		return
	}

	extractedData, err := h.FileUsecase.ReadExcelFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process file"})
		return
	}

	log.Printf("success extrac file %v", fileHeader.Filename)
	log.Println("Extracted Data:", extractedData)

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded and processed successfully",
		"data":    extractedData,
	})
}
