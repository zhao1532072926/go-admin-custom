package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func InitBusinessRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("/api/v1")
	// 注册 FirstUsersRouter
	registerFirstUsersRouter(g, authMiddleware)
	// 注册 FirstDdDetailRouter
	registerFirstDdDetailRouter(g, authMiddleware)

	return g
}
