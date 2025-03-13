package middleware

import (
	"absensi/internal/dto"
	"absensi/internal/entity"
	"absensi/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UseRoleAdminMiddleware(ctx *gin.Context) {
	header := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(header) != 2 && header[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.CommonResponse{
			Success: false,
			Message: "failed to parse Authorization",
		})
		return
	}
	user, err := pkg.JWTDecode(header[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.CommonResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	if user.Role != entity.UserRoleAdmin {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.CommonResponse{
			Success: false,
			Message: "you are not allowed to access this feature",
		})
		return
	}
	ctx.Next()
}
