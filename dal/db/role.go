package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/dy-dayan/user-srv-rbac/model"
	"time"
)

func GetRole(id uint) (*model.Role, error) {
	var r model.Role

	r.ID = id
	db := defaultDB.rClient.Table(r.TableName()).First(&r)
	if err := db.Error; err != nil {
		if db.RecordNotFound() {
			return nil, nil
		} else {
			logrus.Errorf("GetRole error, id:%d, err:%v", id, err)
			return nil, err
		}
	}
	return &r, nil
}

func CreateRole(role *model.Role) error {
	now := time.Now().Unix()
	role.CreatedAt = now
	role.UpdatedAt = now
	db := defaultDB.rClient.Table(role.TableName()).Create(role)
	if err := db.Error; err != nil {
		logrus.Errorf("CreateRole error. name:%s", role.Name)
		return err
	}

	return nil
}

func UpdateRole(role *model.Role) error {
	now := time.Now().Unix()

	//先查询
	roleDB, err := GetRole(role.ID)
	if err != nil {
		return err
	}

	if roleDB == nil {
		// 查询不到就创建
		return CreateRole(role)
	}

	role.UpdatedAt = now
	db := defaultDB.wClient.Table(role.TableName()).Save(role)
	if err = db.Error; err != nil {
		logrus.Errorf("UpdateRole error. err:%v", err)
		return err
	}

	return nil
}



