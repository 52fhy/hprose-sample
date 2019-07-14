package util

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

var Logger *logs.BeeLogger

func InitLog() error {
	Logger = logs.NewLogger(10)

	err := Logger.SetLogger(logs.AdapterMultiFile, fmt.Sprintf(`{"filename":"%s/logs/main.log", "daily":true,"maxdays":7,"rotate":true}`, GetRootPath()))
	if err != nil {
		return errors.New("init beego log error:" + err.Error())
	}
	Logger.Async(1000)
	return nil
}
