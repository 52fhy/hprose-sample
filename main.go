package main

import (
	"flag"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"sample/model"
	"sample/util"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

//定义服务
type SampleService struct {

}

//服务里的方法
func (this *SampleService) GetUserInfo(uid int64) util.State {
	var state util.State

	if uid <= 0 {
		return  state.SetErrCode(1001).SetErrMsg("uid不正确").End()
	}

	var user model.User
	user.Id = uid
	user.Name = "test"
	return state.SetData(user).End()
}

func main() {

	configFile := flag.String("c", "config/rd.ini", "config file")
	flag.Parse()

	err := util.InitConfig(*configFile)
	if err != nil {
		fmt.Printf("load config file fail, err:%v\n", err)
		return
	}

	fmt.Printf("tcp server is running at %s\n", util.Cfg.ListenAddr)
	server := rpc.NewTCPServer("tcp4://" +util.Cfg.ListenAddr + "/")

	//注册func
	server.AddFunction("hello", hello)

	//注册struct，命名空间是SampleService
	server.AddInstanceMethods(&SampleService{}, rpc.Options{NameSpace: "Sample"})
	err = server.Start()
	if err != nil {
		fmt.Printf("start server fail, err:%v\n", err)
		return
	}
}