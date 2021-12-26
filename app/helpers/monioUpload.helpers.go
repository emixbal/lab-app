package helpers

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// func minioUpload(file_path string, file_name string) {
func MinioUpload(file_path string, file_name string, original_ext string) (bool, error) {
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACEESS_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACEESS_ID")
	useSSL := false

	// Initialize minio client object.
	minioClient, err_minio_connect := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err_minio_connect != nil {
		log.Println(err_minio_connect)
		return false, err_minio_connect
	}

	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	objectName := file_name + original_ext
	filePath := file_path
	contentType, err_content_type := GetFileContentType(file_path)
	if err_content_type != nil {
		log.Println(err_content_type)
		return false, err_content_type
	}

	info, err_minio := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err_minio != nil {
		log.Println(err_minio)
		return false, err_minio
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return true, nil
}
