package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type GenderResponse struct {
	Gender string `json:"gender"`
}

func DetectGenderML(fileHeader *multipart.FileHeader) (string, error) {

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	formFile, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return "", err
	}

	_, err = bytes.NewBuffer(nil).ReadFrom(file)
	if err != nil {
		return "", err
	}

	file2, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file2.Close()

	_, err = io.Copy(formFile, file2)
	if err != nil {
		return "", err
	}

	writer.Close()

	fastapiURL := os.Getenv("FASTAPI_GENDER_URL")
	if fastapiURL == "" {
		fastapiURL = "http://127.0.0.1:8000/detect-gender"
	}

	req, err := http.NewRequest("POST", fastapiURL, &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var result GenderResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Gender, nil
}
