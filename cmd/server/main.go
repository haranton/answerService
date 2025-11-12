package main

import (
	"answerService/internals/app"
	"answerService/internals/config"
	"answerService/internals/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.MustLoad()
	logger := logger.GetLogger(cfg.Env)

	application := app.New(cfg, logger)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go application.MustStart()

	<-stop
	application.Close()
	logger.Info("app succesfully stop")

}
