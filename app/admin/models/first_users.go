package models

import (

	"go-admin/common/models"

)

type FirstUsers struct {
    models.Model
    
    Phone string `json:"phone" gorm:"type:varchar(15);comment:phone"` 
    Password string `json:"password" gorm:"type:varchar(20);comment:password"` 
    models.ModelTime
    models.ControlBy
}

func (FirstUsers) TableName() string {
    return "first_users"
}

func (e *FirstUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FirstUsers) GetId() interface{} {
	return e.Id
}