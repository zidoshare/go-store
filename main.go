package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zidoshare/go-store/logs"

	"github.com/gorilla/mux"
	"github.com/zidoshare/go-store/common"
	"github.com/zidoshare/go-store/controller"
	"github.com/zidoshare/go-store/service"
)

var logger = logs.NewLogger(os.Stdout)

// main load common and start up the server on port 8080
func main() {

	//load configuration
	common.LoadConf()
	logger.Info("preparing some jobs...")
	//connect db
	service.Connect()
	//disconnect on exits
	defer service.DisConnect()
	//load routers
	r := mux.NewRouter()
	controller.Load(r)
	//config server
	srv := &http.Server{
		Addr:         common.Conf.Server,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	//async start
	go func() {
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
	logger.Infof("finish,lisenning on [%s]", common.Conf.Server)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err)
	}

}
