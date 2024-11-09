package library

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)


func UploadFile(r *http.Request, fileField string, folder string, fileExt string) (string, error) {

	file, _, err := r.FormFile(fileField)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	fileName := fmt.Sprintf("%s_%s.%s", fileField, time.Now().Format("20060102150405"), fileExt)
	filePath := folder + "/" + fileName

	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory for file: %v", err)
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	return filePath, nil
}
