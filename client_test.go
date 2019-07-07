package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"sample/service"
	"testing"
)

func GetClient() *rpc.TCPClient {
	return rpc.NewTCPClient("tcp4://127.0.0.1:8050")
}

func TestSampleService_GetUserInfo(t *testing.T) {
	client := GetClient()

	defer client.Close()
	var comment service.SampleService
	client.UseService(&comment, "Sample")

	rep := comment.GetUserInfo(10001)
	if rep.ErrCode > 0 {
		t.Error(rep.ErrMsg)
	} else {
		t.Log(rep.Data)
	}
}
