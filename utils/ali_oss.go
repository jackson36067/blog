package utils

import (
	"blog/global"
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadImage(file *multipart.FileHeader) (string, error) {
	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d_%d%s", time.Now().Unix(), rand.Intn(10000), fileExt)
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// 获取ali_oss配置信息
	aliOssConf := global.Conf.AliOss
	// 创建oss客户端
	client, err := oss.New(aliOssConf.EndPoint, aliOssConf.AccessKeyId, aliOssConf.AccessKeySecret)
	if err != nil {
		return "", err
	}
	// 获取bucket
	bucket, err := client.Bucket(aliOssConf.BucketName)
	// 上传文件流到oss
	err = bucket.PutObject(fileName, src)
	if err != nil {
		return "", err
	}
	fileURL := fmt.Sprintf("http://%s.%s/%s", aliOssConf.BucketName, strings.Split(aliOssConf.EndPoint, "//")[1], fileName)
	return fileURL, nil
}
