package bootstrap

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	appServer *gin.Engine
)

func initServer() {
	appServer = gin.Default()
	appServer.Use(gin.Logger())
}

func StartServer() {
	initConfig()
	initServer()
	initService()
	initController()
	initRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", appConfig.App.Port),
		Handler:        appServer,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("down")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
