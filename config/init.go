package config

// 推荐使用toml文件作为配置文件
import (
	"os"
)

var (
	confPath = "confPath"
	Conf     = &config{}
)

// Interpreter gets config form source
type Interpreter interface {
	Local(path string) error
	Remote() error
}

// Source is the config source
type Source interface {
	Fetch(file string) ([]byte, error)
}

// Init get the config
func Init(i Interpreter) error {
	path, isExist := os.LookupEnv(confPath)
	if isExist {
		return i.Local(path)
	}
	return i.Remote()
}
