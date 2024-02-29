package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VicOsewe/secure-pass/pkg/config"
	"github.com/VicOsewe/secure-pass/pkg/infrastructure/databases/postgres"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() {
	apiBasePath := "/api"
	router := gin.Default()
	apiGroup := router.Group(apiBasePath)
	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGroup.POST("/signup")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)

	go func(quit chan os.Signal) {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	}(quit)

	log.Println("ShutDown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}

	log.Println("Server exiting")
}

// InitHandlers initializes all the handlers dependencies
func InitializeHandlers() {
	// ctx := context.Background()
	env := config.NewEnv()
	db := postgres.NewSecurePassDB(env)
	fmt.Println(db)
}
