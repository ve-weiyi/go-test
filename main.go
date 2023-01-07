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
	log.Printf("main")

	initialize.MysqlConnect()
	r := initialize.RouterInit()

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
