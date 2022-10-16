package controllers

type ResponseFormat struct {
	E int         `json:"e"` // error
	M string      `json:"m"` // message
	D interface{} `json:"d"` // data
}

var codeMessage = map[int]string{
	0:     "操作成功",
	1:     "操作失败",
	10010: "获取access_token失败",
}

const (
	SUCC             = 0
	FAIL             = 1
	ERR_ACCESS_TOKEN = 10010
)

func Succ(data interface{}) (r ResponseFormat) {
	r = ResponseFormat{
		E: SUCC,
		M: codeMessage[SUCC],
		D: data,
	}
	return r
	//r := new ResponseFormat
	//r.E = SUCC
	//r.M = codeMessage[SUCC]
	//r.D = data
	//return r
}

func Fail(message string) (r ResponseFormat) {
	r = ResponseFormat{
		E: FAIL,
		M: message,
		D: "",
	}
	if message == "" {
		r.M = codeMessage[FAIL]
	}
	return r
}

func FailByCode(code int) (r ResponseFormat) {
	r = ResponseFormat{
		E: code,
		M: codeMessage[code],
		D: "",
	}
	return r
}
