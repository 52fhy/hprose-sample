package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//获取项目根目录，依据：可执行文件在根目录
func GetRootPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
