package controllers

type ResponseFormat struct {
	E int  `json:"e"`// error
	M string `json:"m"`// message
	D interface{}  `json:"d"` // data
}

var codeMessage = map[int]string{
	0:"操作成功",
	1:"操作失败",
	10010:"获取access_token失败",
}

const (
	SUCC = 0
	FAIL = 1
	ERR_ACCESS_TOKEN = 10010
)
func (r *ResponseFormat)Succ(data interface{}) *ResponseFormat {
	r.E = SUCC
	r.M = codeMessage[SUCC]
	r.D = data
	return r
}

func (r *ResponseFormat)Fail(message string) *ResponseFormat {
	if message == "" {
		r.M = codeMessage[FAIL]
	}
	return r
}

func (r *ResponseFormat)FailByCode(code int) *ResponseFormat {
	r.E = code
	r.M = codeMessage[code]
	return r
}



