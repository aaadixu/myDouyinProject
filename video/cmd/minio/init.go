package minio

import (
	"context"
	"douyinProject/video/cmd/consts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client

func InitMinioClient() {
	var err error

	// 初使化 minio client对象。
	// Initialize minio client object.
	MinioClient, err = minio.New(consts.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(consts.AccessKeyID, consts.SecretAccessKey, ""),
		Secure: consts.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 判断桶是否存在
	exists, errBucketExists := MinioClient.BucketExists(context.Background(), consts.VideoBucketName)
	if errBucketExists != nil {
		log.Fatalln(err)
	}
	if !exists { // 不存在则创建
		err = MinioClient.MakeBucket(context.Background(), consts.VideoBucketName, minio.MakeBucketOptions{Region: consts.Location})
		if err != nil {
			log.Fatalln(err)
		}
	}

	// 判断桶是否存在
	exists, errBucketExists = MinioClient.BucketExists(context.Background(), consts.CoverBucketName)
	if errBucketExists != nil {
		log.Fatalln(err)
	}
	if !exists { // 不存在则创建
		err = MinioClient.MakeBucket(context.Background(), consts.CoverBucketName, minio.MakeBucketOptions{Region: consts.Location})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
