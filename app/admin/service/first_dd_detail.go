package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type FirstDdDetail struct {
	service.Service
}

// GetPage 获取FirstDdDetail列表
func (e *FirstDdDetail) GetPage(c *dto.FirstDdDetailGetPageReq, p *actions.DataPermission, list *[]models.FirstDdDetail, count *int64) error {
	var err error
	var data models.FirstDdDetail

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FirstDdDetailService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FirstDdDetail对象
func (e *FirstDdDetail) Get(d *dto.FirstDdDetailGetReq, p *actions.DataPermission, model *models.FirstDdDetail) error {
	var data models.FirstDdDetail

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFirstDdDetail error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FirstDdDetail对象
func (e *FirstDdDetail) Insert(c *dto.FirstDdDetailInsertReq) error {
    var err error
    var data models.FirstDdDetail
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FirstDdDetailService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FirstDdDetail对象
func (e *FirstDdDetail) Update(c *dto.FirstDdDetailUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.FirstDdDetail{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("FirstDdDetailService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除FirstDdDetail
func (e *FirstDdDetail) Remove(d *dto.FirstDdDetailDeleteReq, p *actions.DataPermission) error {
	var data models.FirstDdDetail

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveFirstDdDetail error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
