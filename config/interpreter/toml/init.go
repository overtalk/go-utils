package toml

import (
	"github.com/BurntSushi/toml"

	"github.com/qinhan-shu/go-utils/config"
)

type Toml struct{}

func (t *Toml) Local(path string) error {
	_, err := toml.DecodeFile(path, config.Conf)
	return err
}

func (t *Toml) Remote() error {
	_, err := toml.Decode("", config.Conf)
	return err
}
