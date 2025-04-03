package main

import (
	"absensi/internal/repository"
	"absensi/internal/route"
	"absensi/internal/service"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
	))
	if err != nil {
		zap.L().Fatal("failed to connect to database", zap.Error(err))
	}

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://google.com"}
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	app.Use(cors.New(config))

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	route.New(app, services)

	app.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "OK",
		})
	})

	app.Run("0.0.0.0:8000")
}
