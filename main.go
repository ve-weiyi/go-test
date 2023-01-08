package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Printf("init")
}

func main() {
	log.Printf("main")

	//initialize.MysqlConnect()
	//r := initialize.RouterInit()
	//
	//// 3.监听端口，默认在8080
	//// Run("里面不指定端口号默认为8080")
	//r.Run(":8000")
	deferTest()
}

// defer panic recovery

func deferTest() {
	defer log.Printf("1")

	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover")
		}
	}()

	log.Printf("2")
	log.Printf("3")

	panic("error")
}
