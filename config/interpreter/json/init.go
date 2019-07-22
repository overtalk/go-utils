package toml

import (
	"encoding/json"

	"github.com/qinhan-shu/go-utils/config"
)

type Json struct{}

func (j *Json) Local(path string) error {
	return json.Unmarshal(nil, config.Conf)
}

func (j *Json) Remote() error {
	return json.Unmarshal(nil, config.Conf)
}
