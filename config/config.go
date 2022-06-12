package config

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "github.com/lib/pq"
)

type YamlStruce struct {
	Common struct{
		Host string `yaml:"host"`
		Schema string `yaml:"schema"`
		Imghost string `yaml:"imghost"`
		DocumentRoot string `yaml:"document_root"`
		Vendor string `yaml:"vendor"`
		Cache string `yaml:"cache"`
		Data string `yaml:"data"`
		Cdndomain string `yaml:"cdndomain"`
		Envname string `yaml:"envname"`
		Wkhtmltopdf string `yaml:"wkhtmltopdf"`
		RsyncModule string `yaml:"rsync_module"`
		DeployName string `yaml:"deploy_name"`
	}
	Redis struct{
		Servers []string `yaml:"servers"`
		Prefix string `yaml:"prefix"`
	}
	Db map[string]struct{
		Adapter string `yaml:"Adapter"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Dbname string `yaml:"dbname"`
		Port int `yaml:"port"`
		Host string `yaml:"host"`
		Slaves []struct{
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}
	}
}

var YamlFile *YamlStruce

var Redisdb *redis.Client

var Postgres *sql.DB

func init() {
	// 初始化main.yaml配置文件
	configFile, err := ioutil.ReadFile("./src/main.yaml")
	//configFile, err := ioutil.ReadFile("/Users/hezhongli/data/www/htdocs/officeflow/server-api/config/main.yaml")
	if err != nil {
		panic(err.Error())
	}
	YamlFile = &YamlStruce{}
	err = yaml.Unmarshal(configFile,YamlFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("配置文件main.yaml解析成功")
	fmt.Println(YamlFile)

	// 初始化 redis
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     YamlFile.Redis.Servers[0], // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	_, err = Redisdb.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Redis连接成功")

	// 初始化postgreSQL
	pgdsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", YamlFile.Db["db"].Host, YamlFile.Db["db"].Port, YamlFile.Db["db"].User, YamlFile.Db["db"].Pass, YamlFile.Db["db"].Dbname)
	Postgres, err := sql.Open("postgres", pgdsn)
	fmt.Println(pgdsn)
	if err != nil {
		panic(err.Error())
	}
	if err = Postgres.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("主数据库Postgres连接成功")
}