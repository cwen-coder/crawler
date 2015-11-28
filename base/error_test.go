package base

import (
	"testing"
)

func Test_NewCrawlerError(t *testing.T) {
	err := NewCrawlerError(DOWNLOADER_ERROR, "下载错误")
	if err.Type() != DOWNLOADER_ERROR {
		t.Error("Type 出错")
	} else {
		t.Log("Type 通过")
	}

}
