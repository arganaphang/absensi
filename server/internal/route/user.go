package route

import (
	"database/sql"
	"errors"
	"net/http"

	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/internal/middleware"
	"absensi/internal/service"
	"absensi/pkg"

	"github.com/gin-gonic/gin"
)

type UserRoute interface {
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	UpdatePassword(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type userRoute struct {
	Services *service.Services
}

func NewUserRoute(engine *gin.Engine, services *service.Services) UserRoute {
	route := &userRoute{
		Services: services,
	}
	// All Routes
	{
		engine.POST("/login", route.Login)
		engine.POST("/refresh-token", route.RefreshToken)
		if gin.Mode() == gin.DebugMode {
			engine.POST("/create-user", route.Create)
		}
	}
	user := engine.Group("/profile")
	user.Use(middleware.UseRoleAllMiddleware)
	{
		user.GET("", route.GetProfile)
		user.PUT("/update", route.UpdateProfile)
		user.POST("/password-change", route.UpdatePassword)
		// TODO: Add change password endpoint
	}
	// Admin Routes
	admin := engine.Group("/users")
	admin.Use(middleware.UseRoleAdminMiddleware)
	{
		admin.GET("", route.GetAll)
		admin.GET(":id", route.GetByID)
		admin.POST("", route.Create)
		admin.PUT(":id", route.Update)
		admin.DELETE(":id", route.Delete)
	}

	return route
}

func (r userRoute) GetProfile(c *gin.Context) {
	userAuth, ok := c.Get(middleware.JWTContextUserKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "failed to get authenticated user"})
		return
	}
	user, ok := userAuth.(*entity.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "authenticated user is broken"})
		return
	}
	user, err := r.Services.UserService.GetByID(c, user.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetProfileResponse{
		Success: true,
		Message: "get profile",
		Data:    dto.UserGetProfileResponseData{User: user},
	})
}

func (r userRoute) UpdateProfile(c *gin.Context) {
	userAuth, ok := c.Get(middleware.JWTContextUserKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "failed to get authenticated user"})
		return
	}
	user, ok := userAuth.(*entity.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "authenticated user is broken"})
		return
	}
	var body dto.UserUpdateProfileRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	userResponse, err := r.Services.UserService.Update(c, user.ID, dto.UserUpdateRequest{
		Email:     body.Email,
		Fullname:  body.Fullname,
		Birthdate: body.Birthdate,
		Position:  user.Position,
		Phone:     body.Phone,
		Address:   body.Address,
		Role:      user.Role,
	})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserUpdateProfileResponse{
		Success: true,
		Message: "profile updated",
		Data:    dto.UserUpdateProfileResponseData{User: userResponse},
	})
}

func (r userRoute) UpdatePassword(c *gin.Context) {
	userAuth, ok := c.Get(middleware.JWTContextUserKey)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "failed to get authenticated user"})
		return
	}
	user, ok := userAuth.(*entity.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "authenticated user is broken"})
		return
	}
	var body dto.PasswordChangeRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	if body.Password != body.RePassword {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: "password & re-password doesn't match"})
		return
	}
	userResponse, err := r.Services.UserService.UpdatePassword(c, user.ID, body)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserUpdateProfileResponse{
		Success: true,
		Message: "profile updated",
		Data:    dto.UserUpdateProfileResponseData{User: userResponse},
	})
}

func (r userRoute) GetAll(c *gin.Context) {
	queries := dto.UserGetAllRequest{}
	if err := c.BindQuery(&queries); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: "failed to parse queries"})
		return
	}
	users, meta, err := r.Services.UserService.GetAll(c, queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetAllResponse{
		Success: true,
		Message: "get users",
		Data:    dto.UserGetAllResponseData{Users: users, Meta: meta},
	})
}

func (r userRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := r.Services.UserService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetByIDResponse{
		Success: true,
		Message: "get user",
		Data:    dto.UserGetByIDResponseData{User: user},
	})
}

func (r userRoute) Create(c *gin.Context) {
	var body dto.UserCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	userExist, err := r.Services.UserService.GetByEmail(c, body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if userExist != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: "user already exist"})
		return
	}
	user, err := r.Services.UserService.Create(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.UserCreateResponse{
		Success: true,
		Message: "user created",
		Data:    dto.UserCreateResponseData{User: user},
	})
}

func (r userRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	user, err := r.Services.UserService.Update(c, id, body)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserUpdateResponse{
		Success: true,
		Message: "user updated",
		Data:    dto.UserUpdateResponseData{User: user},
	})
}

func (r userRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := r.Services.UserService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	err = r.Services.UserService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserDeleteResponse{Success: true, Message: "delete user"})
}

func (r userRoute) Login(c *gin.Context) {
	var body dto.LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	_, err := r.Services.UserService.GetByEmail(c, body.Email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: "user not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.Login(c, body)
	if err != nil && errors.Is(err, pkg.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.LoginResponse{
		Success: true,
		Message: "login success",
		Data:    dto.LoginResponseData{Token: token},
	})
}

func (r userRoute) RefreshToken(c *gin.Context) {
	var body dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.RefreshToken(c, body)
	if err != nil && errors.Is(err, pkg.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		Success: true,
		Message: "refresh token success",
		Data:    dto.RefreshTokenResponseData{Token: token},
	})
}
