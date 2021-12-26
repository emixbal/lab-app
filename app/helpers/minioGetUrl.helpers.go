package helpers

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioGetUrl(object_name string) (string, error) {
	// ======
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACEESS_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACEESS_ID")
	useSSL := false
	minioClient, err_minio_conn := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err_minio_conn != nil {
		log.Println(err_minio_conn)
	}
	// =======

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	// Generates a presigned url which expires in a day.
	image_url, err := minioClient.PresignedGetObject(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), object_name, time.Second*24*60*60, reqParams)
	if err != nil {
		log.Println(err)
		return "", err
	}

	image_url_str := fmt.Sprintf("%v", image_url)
	return image_url_str, nil
}
