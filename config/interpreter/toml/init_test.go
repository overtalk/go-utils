package toml_test

import (
	"os"
	"testing"

	"github.com/qinhan-shu/go-utils/config"
	"github.com/qinhan-shu/go-utils/config/interpreter/toml"
)

func TestLocal(t *testing.T) {
	if err := os.Setenv("confPath", "/Users/qinhan/go/src/github.com/qinhan-shu/go-utils/config/interpreter/toml/test.toml"); err != nil {
		t.Error(err)
		return
	}
	if err := config.Init(&toml.Toml{}, nil); err != nil {
		t.Error(err)
		return
	}
	t.Log(config.Conf)
}
