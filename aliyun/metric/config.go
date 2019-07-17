package metric

type Config struct {
	AccessKey string // aliyun access key
	SecretKey string // aliyun secret key
	Region    string // Region id
	GroupID   int64  // application group id
	Interval  string // push interval
}
