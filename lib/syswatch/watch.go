package syswatch

import (
	"go-app/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

// 用于监听系统退出程序的信号
//
// fn可传递一个退出后需要执行的清理逻辑
// 只能kill -15 不能kill -9
func After(fn func()) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func(fn func()) {
		logger.Info("等待停止信号")
		<-c
		logger.Info("收到停止信号")
		fn()
		os.Exit(1)
	}(fn)
}
