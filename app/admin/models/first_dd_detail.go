package models

import (
    "time"

	"go-admin/common/models"

)

type FirstDdDetail struct {
    models.Model
    
    Sender string `json:"sender" gorm:"type:varchar(15);comment:sender"` 
    SpuId string `json:"spuId" gorm:"type:varchar(20);comment:spu_id"` 
    SpuName string `json:"spuName" gorm:"type:text;comment:spu_name"` 
    ShopId string `json:"shopId" gorm:"type:varchar(15);comment:shop_id"` 
    ShopName string `json:"shopName" gorm:"type:text;comment:shop_name"` 
    SendTime time.Time `json:"sendTime" gorm:"type:timestamp;comment:send_time"` 
    DataBody string `json:"dataBody" gorm:"type:text;comment:data_body"` 
    Remove int64 `json:"remove" gorm:"type:int;comment:remove"` 
    Youhui string `json:"youhui" gorm:"type:text;comment:youhui"` 
    models.ModelTime
    models.ControlBy
}

func (FirstDdDetail) TableName() string {
    return "first_dd_detail"
}

func (e *FirstDdDetail) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FirstDdDetail) GetId() interface{} {
	return e.Id
}