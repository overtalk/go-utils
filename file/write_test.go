package file_test

import (
	"testing"

	"github.com/qinhan-shu/go-utils/file"
)

func TestWrite(t *testing.T) {
	path := "/Users/qinhan/judge/1.txt"
	writeBytes := []byte("qinhan")
	if err := file.Write(path, writeBytes); err != nil {
		t.Error(err)
		return
	}
}
