package helpers

import (
	"encoding/csv"
	"io"
	"mime/multipart"
)

func CheckMatchFileType(fileHeader *multipart.FileHeader, expectFileTypes []string) bool {
	fileType := fileHeader.Header.Get("content-Type")

	for _, expectFileType := range expectFileTypes {
		if fileType == expectFileType {
			return true
		}
	}

	return false
}

func ReadFileContent(csvFile multipart.File) (content []string, err error) {
	reader := csv.NewReader(csvFile)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}

		content = append(content, row[0])
	}

	return content, nil
}
