package metric_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"

	"github.com/qinhan-shu/go-utils/aliyun/metric"
	"github.com/qinhan-shu/go-utils/random"
)

type testClient struct {
	key string
}

func (t testClient) GetCustomMetricMetricList() cms.PutCustomMetricMetricList {
	num, _ := random.RandInt(0, 100)
	list := cms.PutCustomMetricMetricList{
		MetricName: t.key,
		Dimensions: fmt.Sprintf(`{"dimensionality":"%s"}`, t.key),
		Values:     fmt.Sprintf(`{"value":%d}`, num),
	}
	return list
}

func TestCmsGroup(t *testing.T) {
	accessKey, isExist := os.LookupEnv("AccessKey")
	if !isExist {
		t.Errorf("key `AccessKey` is not Exist")
	}

	secretKey, isExist := os.LookupEnv("SecretKey")
	if !isExist {
		t.Errorf("key `SecretKey` is not Exist")
	}

	c := metric.Config{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    "cn-shanghai",
		GroupID:   0,
		Interval:  "*/20 * * * * *",
	}

	client1 := testClient{key: "test1"}
	client2 := testClient{key: "test2"}
	client3 := testClient{key: "test3"}
	client4 := testClient{key: "test4"}

	g, err := metric.CmsGroup(c, client1, client2, client3, client4)
	if err != nil {
		t.Error(err)
		return
	}

	g.Start()

	time.Sleep(5 * time.Minute)
}
