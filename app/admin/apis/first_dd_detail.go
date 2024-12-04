package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type FirstDdDetail struct {
	api.Api
}

// GetPage 获取FirstDdDetail列表
// @Summary 获取FirstDdDetail列表
// @Description 获取FirstDdDetail列表
// @Tags FirstDdDetail
// @Param sender query string false "sender"
// @Param spuId query string false "spu_id"
// @Param spuName query string false "spu_name"
// @Param shopId query string false "shop_id"
// @Param shopName query string false "shop_name"
// @Param remove query int64 false "remove"
// @Param youhui query string false "youhui"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FirstDdDetail}} "{"code": 200, "data": [...]}"
// @Router /api/v1/first-dd-detail [get]
// @Security Bearer
func (e FirstDdDetail) GetPage(c *gin.Context) {
    req := dto.FirstDdDetailGetPageReq{}
    s := service.FirstDdDetail{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FirstDdDetail, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FirstDdDetail失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取FirstDdDetail
// @Summary 获取FirstDdDetail
// @Description 获取FirstDdDetail
// @Tags FirstDdDetail
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FirstDdDetail} "{"code": 200, "data": [...]}"
// @Router /api/v1/first-dd-detail/{id} [get]
// @Security Bearer
func (e FirstDdDetail) Get(c *gin.Context) {
	req := dto.FirstDdDetailGetReq{}
	s := service.FirstDdDetail{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.FirstDdDetail

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FirstDdDetail失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建FirstDdDetail
// @Summary 创建FirstDdDetail
// @Description 创建FirstDdDetail
// @Tags FirstDdDetail
// @Accept application/json
// @Product application/json
// @Param data body dto.FirstDdDetailInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/first-dd-detail [post]
// @Security Bearer
func (e FirstDdDetail) Insert(c *gin.Context) {
    req := dto.FirstDdDetailInsertReq{}
    s := service.FirstDdDetail{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建FirstDdDetail失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FirstDdDetail
// @Summary 修改FirstDdDetail
// @Description 修改FirstDdDetail
// @Tags FirstDdDetail
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FirstDdDetailUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/first-dd-detail/{id} [put]
// @Security Bearer
func (e FirstDdDetail) Update(c *gin.Context) {
    req := dto.FirstDdDetailUpdateReq{}
    s := service.FirstDdDetail{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改FirstDdDetail失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除FirstDdDetail
// @Summary 删除FirstDdDetail
// @Description 删除FirstDdDetail
// @Tags FirstDdDetail
// @Param data body dto.FirstDdDetailDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/first-dd-detail [delete]
// @Security Bearer
func (e FirstDdDetail) Delete(c *gin.Context) {
    s := service.FirstDdDetail{}
    req := dto.FirstDdDetailDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除FirstDdDetail失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
