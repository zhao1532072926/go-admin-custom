package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FirstUsersGetPageReq struct {
	dto.Pagination     `search:"-"`
    Phone string `form:"phone"  search:"type:exact;column:phone;table:first_users" comment:"phone"`
    FirstUsersOrder
}

type FirstUsersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:first_users"`
    Phone string `form:"phoneOrder"  search:"type:order;column:phone;table:first_users"`
    Password string `form:"passwordOrder"  search:"type:order;column:password;table:first_users"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:first_users"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:first_users"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:first_users"`
    
}

func (m *FirstUsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FirstUsersInsertReq struct {
    Id int `json:"-" comment:""` // 
    Phone string `json:"phone" comment:"phone"`
    Password string `json:"password" comment:"password"`
    common.ControlBy
}

func (s *FirstUsersInsertReq) Generate(model *models.FirstUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Phone = s.Phone
    model.Password = s.Password
}

func (s *FirstUsersInsertReq) GetId() interface{} {
	return s.Id
}

type FirstUsersUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Phone string `json:"phone" comment:"phone"`
    Password string `json:"password" comment:"password"`
    common.ControlBy
}

func (s *FirstUsersUpdateReq) Generate(model *models.FirstUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Phone = s.Phone
    model.Password = s.Password
}

func (s *FirstUsersUpdateReq) GetId() interface{} {
	return s.Id
}

// FirstUsersGetReq 功能获取请求参数
type FirstUsersGetReq struct {
     Id int `uri:"id"`
}
func (s *FirstUsersGetReq) GetId() interface{} {
	return s.Id
}

// FirstUsersDeleteReq 功能删除请求参数
type FirstUsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FirstUsersDeleteReq) GetId() interface{} {
	return s.Ids
}
