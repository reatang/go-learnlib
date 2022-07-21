package utils

import (
	"strings"
	"testing"
)

func TestGetCurrentAbPath(t *testing.T) {
	path := GetCurrentAbPath()

	if !strings.HasSuffix(path, "go-learnlib") {
		t.Error("测试失败")
	}
}
