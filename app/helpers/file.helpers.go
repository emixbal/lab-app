package helpers

import (
	"log"
	"net/http"
	"os"
)

func GetFileContentType(file_path string) (string, error) {
	f, err_open := os.Open(file_path)
	if err_open != nil {
		log.Println(err_open)
		return "", err_open
	}

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err_read := f.Read(buffer)
	if err_read != nil {
		return "", err_read
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
