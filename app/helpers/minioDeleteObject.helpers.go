package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioDeleteObject(object_name string) (string, error) {
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

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		// VersionID:        "myversionid",
	}
	err := minioClient.RemoveObject(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), object_name, opts)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return "", nil
}
