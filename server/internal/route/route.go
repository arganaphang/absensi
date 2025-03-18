package route

import (
	"absensi/internal/service"

	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine, services *service.Services) {
	// User Routes
	NewUserRoute(engine, services)

	// Attendance Routes
	NewAttendanceRoute(engine, services)
}
