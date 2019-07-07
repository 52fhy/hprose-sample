package util

type State struct {
	ErrCode int64
	ErrMsg  string
	Data    interface{}
}

func (r *State) SetErrCode(errCode int64) *State {
	r.ErrCode = errCode
	return r
}

func (r *State) SetErrMsg(errMsg string) *State {
	r.ErrMsg = errMsg
	return r
}

func (r *State) SetData(data interface{}) *State {
	r.Data = data
	return r
}

func (r *State) End() State {
	return *r
}
