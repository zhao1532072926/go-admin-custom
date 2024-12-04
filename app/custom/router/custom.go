package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/custom/apis"
)

func registerCustomRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Custom{}
	r := v1.Group("/")
	{
		r.POST("", api.Insert)
	}
}
