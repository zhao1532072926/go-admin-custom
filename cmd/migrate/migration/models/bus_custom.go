package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Custom struct {
	gorm.Model
	Content datatypes.JSON `gorm:"type:json;comment:定制数据"` //使用 datatypes.JSON 类型
}
