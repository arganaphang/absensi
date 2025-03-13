package middleware

import (
	"absensi/internal/dto"
	"absensi/pkg"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const JWTContextUserKey = "JWTContextUserKey"

func UseRoleAllMiddleware(ctx *gin.Context) {
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

		fmt.Println("OALAH MASUK KE ALL JANCOK")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.CommonResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.Set(JWTContextUserKey, user)
	ctx.Next()
}
