package main

import (
	"fmt"
	"log"
	"syscall"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"github.com/fvbock/endless"
)

func main() {
	fmt.Println("exec main func")

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	//router := routers.InitRouter()
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	// 启动服务
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()
}
