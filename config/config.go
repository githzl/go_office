package config

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlStruce struct {
	Common struct {
		Host         string `yaml:"host"`
		Schema       string `yaml:"schema"`
		Imghost      string `yaml:"imghost"`
		DocumentRoot string `yaml:"document_root"`
		Vendor       string `yaml:"vendor"`
		Cache        string `yaml:"cache"`
		Data         string `yaml:"data"`
		Cdndomain    string `yaml:"cdndomain"`
		Envname      string `yaml:"envname"`
		Wkhtmltopdf  string `yaml:"wkhtmltopdf"`
		RsyncModule  string `yaml:"rsync_module"`
		DeployName   string `yaml:"deploy_name"`
	}
	Redis struct {
		Servers []string `yaml:"servers"`
		Auth    string   `yaml:"auth"`
		Db      int      `yaml:"db"`
		Prefix  string   `yaml:"prefix"`
	}
	Db map[string]struct {
		Adapter string `yaml:"Adapter"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Dbname  string `yaml:"dbname"`
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		Slaves  []struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}
	}
}

var YamlFile *YamlStruce

var Redisdb *redis.Client

func init() {
	// 初始化main.yaml配置文件
	configFile, err := ioutil.ReadFile("./src/main.yaml")
	//configFile, err := ioutil.ReadFile("/Users/hezhongli/data/www/htdocs/officeflow/server-api/config/main.yaml")
	if err != nil {
		panic(err.Error())
	}
	YamlFile = &YamlStruce{}
	err = yaml.Unmarshal(configFile, YamlFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("配置文件main.yaml解析成功")

	// 初始化 redis
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     YamlFile.Redis.Servers[0], // use default Addr
		Password: "",                        // no password set
		DB:       0,                         // use default DB
	})
	_, err = Redisdb.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Redis连接成功")
}
