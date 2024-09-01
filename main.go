package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(ApiKeyAuth())
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"Number": 10}) })
	router.POST("/", func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, nil)
		}
		log.Println(string(data))
		c.JSON(http.StatusOK, nil)
	})
	srv := &http.Server{
		Addr:    ":80",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server.....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds")
	}
	log.Println("Server exiting")
}
