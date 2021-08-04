package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	var logger *zap.Logger
	logger, _ = zap.NewProduction()
	logger.Debug("i am debug")       // 这行不会打印，因为默认日志级别是 INFO
	logger.Info("i am info")         // INFO  级别日志，这个会正常打印
	logger.Warn("i am warn")         // WARN  级别日志，这个会正常打印
	logger.Error("i am error")       // ERROR 级别日志，这个会打印，并附带堆栈信息
	logger.Fatal("i am fatal")       // FATAL 级别日志，这个会打印，附带堆栈信息，并调用 os.Exit 退出
	fmt.Println("can i be printed?") // 这行不会打印，呃...上面已经退出了
}


