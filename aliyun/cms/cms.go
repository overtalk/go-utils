package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/robfig/cron"
)

// CmsCron defines cms cron
type cmsCron struct {
	cron      *cron.Cron
	cmsClient *cms.Client
	//metricClient model.MetricClient
	//envMap map[string]string
}

// CmsCorn is the constructor of cms corn
func CmsCorn() *cmsCron {
	return &cmsCron{}
}

func (c *cmsCron) Start() {

}
