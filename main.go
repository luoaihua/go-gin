package main

import (
	"context"
	"flag"
	"fmt"
	"go-gin/cron"
	"go-gin/models"
	"go-gin/pkg/gredis"
	"go-gin/pkg/logging"
	"go-gin/pkg/setting"
	"go-gin/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	setting.Setup()
	gredis.Setup()
	models.Setup()
	logging.Setup()

	// 启动一个协程跑定时任务
	go cron.Start()

	runMode := flag.String("runMode", "debug", "runMode")
	flag.Parse()

	router := routers.InitRouter(*runMode)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Fatal("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
