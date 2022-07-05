package pgmodel

type Thirdparty struct {
	Id        int    `json:"id"`
	Appid     string `json:"appid"`
	Appsecret string `json:"appsecret"`
	Whitelist string `json:"white_list"`
	Enable    string `json:"enable"`
	Detail    string `json:"detail"`
	Created   string `json:"created"`
	Creator   int    `json:"creator"`
}
