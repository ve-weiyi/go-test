package main

import (
	"go-test/initialize"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Printf("init")
}

func main() {
	initialize.MysqlConnect()    //初始化链接数据库
	r := initialize.RouterInit() //调用路由   包名.文件中的函数

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8081")

	//deferTest()
}

// defer panic recovery  延迟 恐慌 恢复

//func deferTest() {
//	defer log.Printf("1")
//
//	defer func() {
//		if r := recover(); r != nil {
//			log.Printf("recover")
//		}
//	}()
//
//	log.Printf("2")
//	log.Printf("3")
//
//	panic("error")
//}
