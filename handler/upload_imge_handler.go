package handler

import (
	"io"
	"mime/multipart"
	"os"
	"time"

	"com.test.users_api_test/api/constants"
)

func UploadFileHandler(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	// Create directory if it doesn't exist
	uploadPath := constants.UploadPath + "/"
	err := os.MkdirAll(uploadPath, 0755)
	if err != nil {
		return "", err
	}

	// Get the uploaded file
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	imageName := timestamp + "_" + fileHeader.Filename

	// Create a new file on the server to copy the uploaded file to
	uploadedFilePath := uploadPath + "/" + imageName
	newFile, err := os.Create(uploadedFilePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	// Copy the uploaded file to the newly created file
	_, err = io.Copy(newFile, file)
	if err != nil {
		return "", err
	}
	return imageName, nil
}
