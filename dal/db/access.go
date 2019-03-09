package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/dy-dayan/user-srv-rbac/model"
	"time"
)

func GetAccess(id uint) (*model.Access, error) {
	var r model.Access

	r.ID = id
	db := defaultDB.rClient.Table(r.TableName()).First(&r)
	if err := db.Error; err != nil {
		if db.RecordNotFound() {
			return nil, nil
		} else {
			logrus.Errorf("GetAccess error, id:%d, err:%v", id, err)
			return nil, err
		}
	}
	return &r, nil
}

func CreateAccess(access *model.Access) error {
	now := time.Now().Unix()
	access.CreatedAt = now
	access.UpdatedAt = now
	db := defaultDB.rClient.Table(access.TableName()).Create(access)
	if err := db.Error; err != nil {
		logrus.Errorf("CreateAccess error. name:%s", access.Title)
		return err
	}

	return nil
}

func UpdateAccess(access *model.Access) error {
	now := time.Now().Unix()

	//先查询
	accessDB, err := GetAccess(access.ID)
	if err != nil {
		return err
	}

	if accessDB == nil {
		// 查询不到就创建
		return CreateAccess(access)
	}

	access.UpdatedAt = now
	db := defaultDB.wClient.Table(access.TableName()).Save(access)
	if err = db.Error; err != nil {
		logrus.Errorf("UpdateAccess error. err:%v", err)
		return err
	}

	return nil
}




