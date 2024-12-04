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

type FirstUsers struct {
	api.Api
}

// GetPage 获取FirstUsers列表
// @Summary 获取FirstUsers列表
// @Description 获取FirstUsers列表
// @Tags FirstUsers
// @Param phone query string false "phone"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FirstUsers}} "{"code": 200, "data": [...]}"
// @Router /api/v1/first-users [get]
// @Security Bearer
func (e FirstUsers) GetPage(c *gin.Context) {
    req := dto.FirstUsersGetPageReq{}
    s := service.FirstUsers{}
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
	list := make([]models.FirstUsers, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FirstUsers失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取FirstUsers
// @Summary 获取FirstUsers
// @Description 获取FirstUsers
// @Tags FirstUsers
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FirstUsers} "{"code": 200, "data": [...]}"
// @Router /api/v1/first-users/{id} [get]
// @Security Bearer
func (e FirstUsers) Get(c *gin.Context) {
	req := dto.FirstUsersGetReq{}
	s := service.FirstUsers{}
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
	var object models.FirstUsers

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FirstUsers失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建FirstUsers
// @Summary 创建FirstUsers
// @Description 创建FirstUsers
// @Tags FirstUsers
// @Accept application/json
// @Product application/json
// @Param data body dto.FirstUsersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/first-users [post]
// @Security Bearer
func (e FirstUsers) Insert(c *gin.Context) {
    req := dto.FirstUsersInsertReq{}
    s := service.FirstUsers{}
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
		e.Error(500, err, fmt.Sprintf("创建FirstUsers失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FirstUsers
// @Summary 修改FirstUsers
// @Description 修改FirstUsers
// @Tags FirstUsers
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FirstUsersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/first-users/{id} [put]
// @Security Bearer
func (e FirstUsers) Update(c *gin.Context) {
    req := dto.FirstUsersUpdateReq{}
    s := service.FirstUsers{}
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
		e.Error(500, err, fmt.Sprintf("修改FirstUsers失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除FirstUsers
// @Summary 删除FirstUsers
// @Description 删除FirstUsers
// @Tags FirstUsers
// @Param data body dto.FirstUsersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/first-users [delete]
// @Security Bearer
func (e FirstUsers) Delete(c *gin.Context) {
    s := service.FirstUsers{}
    req := dto.FirstUsersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除FirstUsers失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
