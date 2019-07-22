package oss

import (
	"fmt"
	"io/ioutil"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
aliyunOSS demo
{
	accessKey : "xxx"
	secretKey : "xxx"
	endpoint  : "endpoint=oss-cn-shanghai.aliyuncs.com"
	bucket    : "qinhan"
	dir       : "test/"
}
=> 对于的取文件的格式
=> qinhan@oss-cn-shanghai.aliyuncs.com/test/xxx
*/

type aliyunOSS struct {
	accessKey string // aliyun access key
	secretKey string // aliyun secret key
	endpoint  string // aliyun oss endpoint
	bucket    string // aliyun oss bucket
	dir       string // aliyun oss dir
}

func AliyunOSS(accessKey, secretKey, endpoint, bucket, dir string) *aliyunOSS {
	return &aliyunOSS{
		accessKey: accessKey,
		secretKey: secretKey,
		endpoint:  endpoint,
		bucket:    bucket,
		dir:       dir,
	}
}

func (a *aliyunOSS) Fetch(fileName string) ([]byte, error) {
	client, err := oss.New(a.endpoint, a.accessKey, a.secretKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(a.bucket)
	if err != nil {
		return nil, err
	}

	body, err := bucket.GetObject(a.dir + fileName)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		//logger.GetLogger().Errorf("file %s is empty", a.dir+fileName)
		return nil, fmt.Errorf("file %s is empty", a.dir+fileName)
	}

	return data, nil
}
