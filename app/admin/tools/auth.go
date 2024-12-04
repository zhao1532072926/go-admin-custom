package tools

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"time"
)

type loginVals struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// FirstUserAuth first_user jwt验证
func FirstUserAuth() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.JwtConfig.Timeout != 0 {
		timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
	}

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "first_user",
		Key:             []byte(config.JwtConfig.Secret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     FirstUserPayloadFunc,
		IdentityHandler: FirstUserIdentityHandler,
		Authenticator:   FirstUserAuthenticator,
		Unauthorized:    FirstUserUnauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}

// FirstUserPayloadFunc 载荷处理
func FirstUserPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(models.FirstUsers); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v.Id,
			"phone":         v.Phone,
		}
	}
	return jwt.MapClaims{}
}

// FirstUserIdentityHandler 身份处理
func FirstUserIdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"Phone":       claims["phone"],
	}
}

// FirstUserAuthenticator 认证处理
func FirstUserAuthenticator(c *gin.Context) (interface{}, error) {
	var loginVals struct {
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	s := service.FirstUsers{}
	s.Orm = db
	user, err := s.GetUserByPhone(loginVals.Phone)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	if user.Password != loginVals.Password {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

// FirstUserUnauthorized 未授权处理
func FirstUserUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  message,
	})
}
