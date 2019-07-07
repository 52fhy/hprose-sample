package main

import (
	"flag"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
	"sample/service"
	"sample/util"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {

	configFile := flag.String("c", "config/rd.ini", "config file")
	protocol := flag.String("P", "tcp", "the protocol, such as tcp,http")
	flag.Parse()

	err := util.InitConfig(*configFile)
	if err != nil {
		fmt.Printf("load config file fail, err:%v\n", err)
		return
	}

	fmt.Printf("server is running at %s\n", util.Cfg.ListenAddr)

	if *protocol == "tcp" {
		server := rpc.NewTCPServer("tcp4://" + util.Cfg.ListenAddr + "/")

		//注册func
		server.AddFunction("hello", hello)

		//注册struct，命名空间是SampleService
		server.AddInstanceMethods(&service.SampleService{}, rpc.Options{NameSpace: "Sample"})
		err = server.Start()
		if err != nil {
			fmt.Printf("start server fail, err:%v\n", err)
			return
		}
	} else if *protocol == "http" {
		server := rpc.NewHTTPService()

		//注册func
		server.AddFunction("hello", hello)

		//注册struct，命名空间是SampleService
		server.AddInstanceMethods(&service.SampleService{}, rpc.Options{NameSpace: "Sample"})

		err = http.ListenAndServe(util.Cfg.ListenAddr, server)
		if err != nil {
			fmt.Printf("start server fail, err:%v\n", err)
			return
		}
	} else {
		fmt.Printf("err protocol config : %v\n", *protocol)
		return
	}

}
