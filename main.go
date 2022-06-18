package main

import (
	"go_office/internal/routers"
)

func main(){
	router := routers.Load()
	// 默认端口是8080被tomcat使用了,用指定端口 r.Run(":8888")
	_ = router.Run(":8888")
}