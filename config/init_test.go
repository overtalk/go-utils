package config_test

import (
	"os"
	"testing"

	"github.com/qinhan-shu/go-utils/config"
)

func TestInit(t *testing.T) {
	os.Setenv("confPath", "/Users/qinhan/go/src/github.com/qinhan-shu/go-utils/config/config.toml")

	if err := config.Init(); err != nil {
		t.Error(err)
		return
	}

 	t.Log(config.Conf)
}
