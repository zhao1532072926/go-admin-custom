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

type FirstUsers struct {
	service.Service
}

// GetPage 获取FirstUsers列表
func (e *FirstUsers) GetPage(c *dto.FirstUsersGetPageReq, p *actions.DataPermission, list *[]models.FirstUsers, count *int64) error {
	var err error
	var data models.FirstUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FirstUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FirstUsers对象
func (e *FirstUsers) Get(d *dto.FirstUsersGetReq, p *actions.DataPermission, model *models.FirstUsers) error {
	var data models.FirstUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFirstUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FirstUsers对象
func (e *FirstUsers) Insert(c *dto.FirstUsersInsertReq) error {
	var err error
	var data models.FirstUsers
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FirstUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FirstUsers对象
func (e *FirstUsers) Update(c *dto.FirstUsersUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FirstUsers{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FirstUsersService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FirstUsers
func (e *FirstUsers) Remove(d *dto.FirstUsersDeleteReq, p *actions.DataPermission) error {
	var data models.FirstUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFirstUsers error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// GetUserByPhone 根据手机号获取用户信息, 返回用户信息
func (e *FirstUsers) GetUserByPhone(phone string) (models.FirstUsers, error) {
	var data models.FirstUsers
	err := e.Orm.Model(&data).
		Where("phone = ?", phone).
		First(&data).Error
	if err != nil {
		e.Log.Errorf("FirstUsersService GetUserByPhone error:%s \r\n", err)
		return data, err
	}
	return data, nil
}
