package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"sample/util"
	"testing"
)

//stub
type clientStub struct {
	Hello       func(string) string
	GetUserInfo func(uid int64) util.State
}

func GetClient() *rpc.TCPClient {
	return rpc.NewTCPClient("tcp4://127.0.0.1:8050")
}

//测试服务里的方法
func TestSampleService_GetUserInfo(t *testing.T) {
	client := GetClient()

	defer client.Close()
	var stub clientStub
	client.UseService(&stub, "Sample") //使用命名空间

	rep := stub.GetUserInfo(10001)
	if rep.ErrCode > 0 {
		t.Error(rep.ErrMsg)
	} else {
		t.Log(rep.Data)
	}
}

//测试普通方法
func TestHello(t *testing.T) {
	client := GetClient()

	defer client.Close()
	var stub clientStub
	client.UseService(&stub)

	rep := stub.Hello("func")
	if rep == "" {
		t.Error(rep)
	} else {
		t.Log(rep)
	}
}
