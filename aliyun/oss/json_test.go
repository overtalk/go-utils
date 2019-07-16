package oss_test

import (
	"os"
	"testing"

	"github.com/qinhan-shu/go-utils/aliyun/oss"
)

func TestFetch(t *testing.T) {
	// get oss config form env var
	keys := []string{
		"accessKey",
		"secretKey",
		"endpoint",
		"bucket",
		"dir",
	}
	values := []string{}
	for _, key := range keys {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			t.Errorf("key `%s` is absent", key)
			return
		}
		values = append(values, value)
	}
	// get test.json
	aliOSS := oss.AliyunOSS(values[0], values[1], values[2], values[3], values[4])
	fileBytes, err := aliOSS.Fetch("test.json")
	if err != nil {
		t.Errorf("get file bytes error : %v", err)
		return
	}
	t.Log(string(fileBytes))
}
