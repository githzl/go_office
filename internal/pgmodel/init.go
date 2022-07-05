package pgmodel

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_office/config"
)

var DB *gorm.DB

const SSLMODE = "disable"

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.YamlFile.Db["db"].Host, config.YamlFile.Db["db"].Port, config.YamlFile.Db["db"].User, config.YamlFile.Db["db"].Dbname, config.YamlFile.Db["db"].Pass, SSLMODE)
	fmt.Println(dsn)
	DB, _ = gorm.Open("postgres", dsn)
	if DB.Error != nil {
		fmt.Println(DB.Error)
		panic(DB.Error)
	}
	DB.SingularTable(true)
}
