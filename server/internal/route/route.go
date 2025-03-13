package route

import (
	"absensi/internal/service"

	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine, services *service.Services) {
	// User Route
	NewUserRoute(engine, services)
}
