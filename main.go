package main

import (
	handler "excel-handler/handlers"
	"excel-handler/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the file usecase
	fileUsecase := usecase.NewFileUsecase()

	// Initialize the handler with the repository
	handler := &handler.Handler{FileUsecase: fileUsecase}

	// Endpoint for file upload
	router.POST("/upload", handler.UploadFile)

	// Run the server
	router.Run(":8080")
}
