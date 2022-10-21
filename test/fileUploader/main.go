package main

import (
    "github.com/minio/minio-go/v6"
    "log"
)

func main() {
    endpoint := "ling11.top:9000"
    accessKeyID := "zyd5vDW8kgsRZJpl"
    secretAccessKey := "FRPjHVtacC8hmz2EmhIyWhimWdcNkbtu"
    useSSL := false

    // 初使化minio client对象。
    minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
    if err != nil {
        log.Fatalln(err)
    }
	log.Println("minioClient",minioClient)
    // 创建一个叫mymusic的存储桶。
    bucketName := "mymusic2"
    location := "us-east-1"

    err = minioClient.MakeBucket(bucketName, location)
    if err != nil {
        // 检查存储桶是否已经存在。
        exists, err := minioClient.BucketExists(bucketName)
        if err == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }
    }
    log.Printf("Successfully created %s\n", bucketName)

    // 上传一个zip文件。
    objectName := "1003.zip"
    filePath := "1003.zip"
    contentType := "application/zip"

    // 使用FPutObject上传一个zip文件。
    n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}