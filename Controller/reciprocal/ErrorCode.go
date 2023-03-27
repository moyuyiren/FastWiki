package reciprocal

type ResCode int64

const (
	CodeSuccess ResCode = 10000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeUserPermissionDenied
	CodeExchangeUserPasswd
	CodePageError

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{ //错误码对应的提示信息

	CodeSuccess:              "success",
	CodeInvalidParams:        "请求参数错误",
	CodeUserExist:            "用户已存在",
	CodeUserNotExist:         "用户不存在",
	CodeInvalidPassword:      "用户名或密码错误",
	CodeServerBusy:           "服务器繁忙",
	CodeUserPermissionDenied: "用户权限不足",
	CodeExchangeUserPasswd:   "用户密码修改失败,请检查并重新输入",
	CodePageError:            "分页信息输入错误",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的Token",
}

func (c ResCode) GetMsg() string {
	return codeMsgMap[c]
}
