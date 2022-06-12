package main

import (
	//"github.com/gin-gonic/gin"
	//"go_office/config"
	//"go_office/internal/cacheKey"
	//"go_office/internal/middleware"
	"go_office/internal/routers"
	//"net/http"

)

func main(){
	//configFile, _ := ioutil.ReadFile("./src/main.yaml")
	//configFile, _ := ioutil.ReadFile("/Users/hezhongli/data/www/htdocs/officeflow/server-api/config/main.yaml")
	//conf := &config.Config{}
	//_ = yaml.Unmarshal(configFile,conf)

	//redisdb := redis.NewClient(&redis.Options{
	//	Addr:     config.YamlFile.Redis.Servers[0], // use default Addr
	//	Password: "",               // no password set
	//	DB:       0,                // use default DB
	//})
	//fmt.Println(redisdb)

	//router := gin.Default()
	//apiRouter := router.Group("/api").Use(middleware.ApiMiddleware())
	//apiRouter.GET("home/site-options",func(c *gin.Context) {
	//	val, _ := config.Redisdb.Get("officeflow"+cacheKey.SiteOption()).Result()
	//	c.String(http.StatusOK,"Go1 "+ val)
	////router.GET("api/home/site-options",
	//})
	router := routers.Load()
	// 默认端口是8080被tomcat使用了,用指定端口 r.Run(":8888")
	_ = router.Run(":8888")
}