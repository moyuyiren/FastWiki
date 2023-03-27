package Logic

// 自定义一个error
type PwdError struct {
	msg string
	Pwd string
}

func (a *PwdError) Error() string {
	return a.msg + ", :" + a.Pwd
}
