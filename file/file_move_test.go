package file

import "testing"

// 测试复制单文件
func TestCopySingle(t *testing.T) {
	Copy("../io", "../1212/sadas")
}

// 测试复制文件夹
func TestCopyDir(t *testing.T) {

}
