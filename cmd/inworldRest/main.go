package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	loadConfigs()
}

var (
	Key               string
	Secret            string
	Scene             string
	DisconnectTimeout = 30 * time.Minute
	Emotions          = true

	Port = 3000
)

func main() {
	switch viper.GetString("GIN_MODE") {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	registerEndpoints(r)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Failed to start failed, err: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed, err: %s", err)
	}

	log.Println("Shut down server")
}

func loadConfigs() {
	if v := viper.GetString("INWORLD_KEY"); v == "" {
		log.Fatal("Inworld API Key not set")
	} else {
		Key = v
	}
	if v := viper.GetString("INWORLD_SECRET"); v == "" {
		log.Fatal("Inworld API Secret not set")
	} else {
		Secret = v
	}
	if v := viper.GetString("INWORLD_SCENE"); v == "" {
		log.Fatal("Inworld API Default Scene not set")
	} else {
		Scene = v
	}

	Emotions = viper.GetBool("EMOTIONS")
	if viper.Get("PORT") != nil {
		Port = viper.GetInt("PORT")
	}
	if viper.Get("DISCONNECT_TIMEOUT") != nil {
		DisconnectTimeout = viper.GetDuration("DISCONNECT_TIMEOUT")
	}
}
