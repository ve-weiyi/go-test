package initialize

import (
	"fmt"
	"go-test/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

const dns = "root:mysql7914@(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

func MysqlConnect() {
	var err error
	// 连接数据库
	global.CONFIG.DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "tb_",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	log.Println("mysql connection done")
}
