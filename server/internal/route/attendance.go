package route

import (
	"absensi/internal/dto"
	"absensi/internal/middleware"
	"absensi/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceRoute interface {
	GetAttendances(c *gin.Context)
	GetTodayAttendance(c *gin.Context)
	CreateAttendance(c *gin.Context)
	ApproveAttendance(c *gin.Context)
}

type attendanceRoute struct {
	Services *service.Services
}

func NewAttendanceRoute(engine *gin.Engine, services *service.Services) AttendanceRoute {
	route := &attendanceRoute{
		Services: services,
	}
	// All Routes
	attendance := engine.Group("/attendances")
	attendance.Use(middleware.UseRoleAllMiddleware)
	{
		attendance.GET("", route.GetAttendances)
		attendance.GET("/today", route.GetTodayAttendance)
		attendance.POST("", route.CreateAttendance)
	}

	// Admin Routes
	{
		attendance.POST("/:id/approve", route.ApproveAttendance)
	}
	return route
}

func (r attendanceRoute) GetAttendances(c *gin.Context) {
	queries := dto.GetAttendancesRequest{}
	if err := c.BindQuery(&queries); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: "failed to parse queries"})
		return
	}
	attendances, meta, err := r.Services.AttendanceService.GetAttendances(c, queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.GetAttendancesResponse{
		Success: true,
		Message: "get attendances",
		Data:    dto.GetAttendancesResponseData{Attendances: attendances, Meta: meta},
	})
}

func (r attendanceRoute) GetTodayAttendance(c *gin.Context) {
	attendances, err := r.Services.AttendanceService.GetTodayAttendance(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.GetTodayAttendanceResponse{
		Success: true,
		Message: "get today attendance",
		Data:    dto.GetTodayAttendanceResponseData{Attendances: attendances},
	})
}

func (r attendanceRoute) CreateAttendance(c *gin.Context) {
	body := dto.CreateAttendanceRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: "failed to parse reuest body"})
		return
	}
	attendance, err := r.Services.AttendanceService.CreateAttendance(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.CreateAttendanceResponse{
		Success: true,
		Message: "create attendance",
		Data:    dto.CreateAttendanceResponseData{Attendance: attendance},
	})
}

func (r attendanceRoute) ApproveAttendance(c *gin.Context) {}
