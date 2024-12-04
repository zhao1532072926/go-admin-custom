package apis

import (
	"go-admin/app/custom/service"
	"go-admin/app/custom/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type Custom struct {
	api.Api
}

// Insert 创建自定义数据
// @Summary 创建自定义数据
// @Description 创建自定义数据
// @Tags 自定义
// @Accept  application/json
// @Product application/json
// @Param data body dto.CustomInsertReq true "自定义数据"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/custom [post]
func (e Custom) Insert(c *gin.Context) {
	req := &dto.CustomInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req).
		Errors
	if err != nil {
		e.Error(400, err, "数据解析失败")
		return
	}

	s := service.Custom{}
	err = s.Insert(*req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}

	e.OK(nil, "创建成功")
}
