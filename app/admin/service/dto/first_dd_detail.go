package dto

import (
    "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type FirstDdDetailGetPageReq struct {
	dto.Pagination     `search:"-"`
    Sender string `form:"sender"  search:"type:exact;column:sender;table:first_dd_detail" comment:"sender"`
    SpuId string `form:"spuId"  search:"type:exact;column:spu_id;table:first_dd_detail" comment:"spu_id"`
    SpuName string `form:"spuName"  search:"type:exact;column:spu_name;table:first_dd_detail" comment:"spu_name"`
    ShopId string `form:"shopId"  search:"type:exact;column:shop_id;table:first_dd_detail" comment:"shop_id"`
    ShopName string `form:"shopName"  search:"type:exact;column:shop_name;table:first_dd_detail" comment:"shop_name"`
    Remove int64 `form:"remove"  search:"type:exact;column:remove;table:first_dd_detail" comment:"remove"`
    Youhui string `form:"youhui"  search:"type:exact;column:youhui;table:first_dd_detail" comment:"youhui"`
    FirstDdDetailOrder
}

type FirstDdDetailOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:first_dd_detail"`
    Sender string `form:"senderOrder"  search:"type:order;column:sender;table:first_dd_detail"`
    SpuId string `form:"spuIdOrder"  search:"type:order;column:spu_id;table:first_dd_detail"`
    SpuName string `form:"spuNameOrder"  search:"type:order;column:spu_name;table:first_dd_detail"`
    ShopId string `form:"shopIdOrder"  search:"type:order;column:shop_id;table:first_dd_detail"`
    ShopName string `form:"shopNameOrder"  search:"type:order;column:shop_name;table:first_dd_detail"`
    SendTime string `form:"sendTimeOrder"  search:"type:order;column:send_time;table:first_dd_detail"`
    DataBody string `form:"dataBodyOrder"  search:"type:order;column:data_body;table:first_dd_detail"`
    Remove string `form:"removeOrder"  search:"type:order;column:remove;table:first_dd_detail"`
    Youhui string `form:"youhuiOrder"  search:"type:order;column:youhui;table:first_dd_detail"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:first_dd_detail"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:first_dd_detail"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:first_dd_detail"`
    
}

func (m *FirstDdDetailGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FirstDdDetailInsertReq struct {
    Id int `json:"-" comment:""` // 
    Sender string `json:"sender" comment:"sender"`
    SpuId string `json:"spuId" comment:"spu_id"`
    SpuName string `json:"spuName" comment:"spu_name"`
    ShopId string `json:"shopId" comment:"shop_id"`
    ShopName string `json:"shopName" comment:"shop_name"`
    SendTime time.Time `json:"sendTime" comment:"send_time"`
    DataBody string `json:"dataBody" comment:"data_body"`
    Remove int64 `json:"remove" comment:"remove"`
    Youhui string `json:"youhui" comment:"youhui"`
    common.ControlBy
}

func (s *FirstDdDetailInsertReq) Generate(model *models.FirstDdDetail)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Sender = s.Sender
    model.SpuId = s.SpuId
    model.SpuName = s.SpuName
    model.ShopId = s.ShopId
    model.ShopName = s.ShopName
    model.SendTime = s.SendTime
    model.DataBody = s.DataBody
    model.Remove = s.Remove
    model.Youhui = s.Youhui
}

func (s *FirstDdDetailInsertReq) GetId() interface{} {
	return s.Id
}

type FirstDdDetailUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Sender string `json:"sender" comment:"sender"`
    SpuId string `json:"spuId" comment:"spu_id"`
    SpuName string `json:"spuName" comment:"spu_name"`
    ShopId string `json:"shopId" comment:"shop_id"`
    ShopName string `json:"shopName" comment:"shop_name"`
    SendTime time.Time `json:"sendTime" comment:"send_time"`
    DataBody string `json:"dataBody" comment:"data_body"`
    Remove int64 `json:"remove" comment:"remove"`
    Youhui string `json:"youhui" comment:"youhui"`
    common.ControlBy
}

func (s *FirstDdDetailUpdateReq) Generate(model *models.FirstDdDetail)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Sender = s.Sender
    model.SpuId = s.SpuId
    model.SpuName = s.SpuName
    model.ShopId = s.ShopId
    model.ShopName = s.ShopName
    model.SendTime = s.SendTime
    model.DataBody = s.DataBody
    model.Remove = s.Remove
    model.Youhui = s.Youhui
}

func (s *FirstDdDetailUpdateReq) GetId() interface{} {
	return s.Id
}

// FirstDdDetailGetReq 功能获取请求参数
type FirstDdDetailGetReq struct {
     Id int `uri:"id"`
}
func (s *FirstDdDetailGetReq) GetId() interface{} {
	return s.Id
}

// FirstDdDetailDeleteReq 功能删除请求参数
type FirstDdDetailDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FirstDdDetailDeleteReq) GetId() interface{} {
	return s.Ids
}
