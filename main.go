package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/200lab-training-1/models"
	"github.com/200lab-training-1/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/200labnotes?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&models.Note{})
	fileWrite, err := os.Create("access.log")
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = fileWrite

	r := gin.Default()
	routes.InitRoutes(r, db)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8081"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
}
