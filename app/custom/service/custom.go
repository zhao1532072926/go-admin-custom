package service

import (
	"encoding/json"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/custom/models"
	"go-admin/app/custom/service/dto"
	"gorm.io/datatypes"
)

type Custom struct {
	service.Service
}

func (e *Custom) Insert(req dto.CustomInsertReq) error {
	// 将interface{}转为JSON字节
	jsonBytes, err := json.Marshal(req.Content)
	if err != nil {
		return err
	}
	data := models.Custom{
		Content: datatypes.JSON(jsonBytes),
	}
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertCustomContent error:%s", err)
		return err
	}
	return nil
}
