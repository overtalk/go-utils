package config

// 推荐使用toml文件作为配置文件
import (
	"os"

	"github.com/BurntSushi/toml"
)

var (
	confPath = "confPath"
	Conf     = &config{}
)

// Init get the config
func Init() error {
	path, isExist := os.LookupEnv(confPath)
	if isExist {
		return local(path)
	}
	return remote()
}

func local(path string) error {
	return nil
}

func remote() error {
	_, err := toml.Decode("", Conf)
	return err
}
