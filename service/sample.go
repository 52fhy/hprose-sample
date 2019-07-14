package service

import (
	"sample/model"
	"sample/util"
)

//定义服务
type SampleService struct {
}

//服务里的方法
func (this *SampleService) GetUserInfo(uid int64) util.State {
	var state util.State

	if uid <= 0 {
		util.Logger.Debug("uid error. uid:%d", uid)
		return state.SetErrCode(1001).SetErrMsg("uid不正确").End()
	}

	var user model.User
	user.Id = uid
	user.Name = "test"
	return state.SetData(user).End()
}
