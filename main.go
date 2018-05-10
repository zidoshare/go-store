package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zidoshare/go-store/controller"
	"github.com/zidoshare/go-store/logs"
	"github.com/zidoshare/go-store/service"

	"github.com/zidoshare/go-store/common"
)

var logger = logs.NewLogger(os.Stdout)

// main start up the server on port 8080
func main() {
	//connect db
	service.Connect()
	//config server
	srv := &http.Server{
		Addr:         common.Conf.Server,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      controller.StoreRouters,
	}
	//async start
	go func() {
		//disconnect on exits
		defer service.DisConnect()
		//waiting ^c signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), common.Conf.Wait)
		defer cancel()

		srv.Shutdown(ctx)

		fmt.Println()
		logger.Info("service shutting down ok")
		os.Exit(0)
	}()
	logger.Infof("Preparation is completed,listenning on [%s]", common.Conf.Server)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err)
	}
}
