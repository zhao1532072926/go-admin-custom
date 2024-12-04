package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/tools"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

// registerFirstDdDetailRouter
func registerFirstDdDetailRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FirstDdDetail{}
	g := v1.Group("/first-dd-detail")

	auth, err := tools.FirstUserAuth()
	if err != nil {
		log.Fatal("FirstUserAuth err:", err)
	}
	g.Use(auth.MiddlewareFunc()).POST("/insert", api.Insert)

	r := g.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
