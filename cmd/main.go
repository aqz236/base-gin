package main

import (
	"base-gin/wire"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化应用
	app, cleanup, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("初始化应用失败: %v", err)
	}
	defer cleanup()

	// 启动 HTTP 服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: app.Router,
	}

	// 在 goroutine 中启动服务器
	go func() {
		app.Logger.Info("服务器启动在端口 :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务器失败: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Logger.Info("正在关闭服务器...")

	// 给服务器 5 秒的时间来完成现有的请求
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// 关闭服务器
	if err := server.Close(); err != nil {
		log.Fatalf("服务器关闭失败: %v", err)
	}

	// 等待一秒钟让所有进程结束
	time.Sleep(1 * time.Second)
	app.Logger.Info("服务器已关闭")
}
