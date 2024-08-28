package utils_test

import (
	"testing"
	"github.com/0days-ru/UTMStack/installer/utils"
)
// Test CheckDisk
func TestCheckDisk(t *testing.T) {
	err := utils.CheckDisk(100)
	t.Log(err)
}