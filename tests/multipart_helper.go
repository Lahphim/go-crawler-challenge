package tests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	"go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/core/logs"
)

func GetMultipartAttributesFromFile(filePath string, contentType string) (multipart.File, *multipart.FileHeader, error) {
	realPath := fmt.Sprintf("%s/%s", helpers.RootDir(), filePath)
	headers, payload := CreateMultipartRequestInfo(realPath, contentType)
	req, err := http.NewRequest("POST", "", payload)
	if err != nil {
		return nil, nil, err
	}

	req.Header = headers
	file, fileHeader, err := req.FormFile("file")

	return file, fileHeader, err
}

func CreateMultipartRequestInfo(filePath string, contentType string) (http.Header, *bytes.Buffer) {
	file, err := os.Open(filePath)
	if err != nil {
		logs.Error("Failed to open file: ", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := createFormFile(writer, "file", filepath.Base(filePath), contentType)
	if err != nil {
		logs.Error("Failed to create part from file: ", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		logs.Error("Failed to copy file: ", err)
	}
	writer.Close()

	headers := http.Header{}
	headers.Set("Content-Type", writer.FormDataContentType())

	return headers, body
}

func createFormFile(w *multipart.Writer, fieldname string, filePath string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filePath))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}
