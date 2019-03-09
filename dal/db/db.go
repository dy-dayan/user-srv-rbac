package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"sync"
	"github.com/dy-dayan/user-srv-rbac/util"
)

type DB struct {
	rClient *gorm.DB
	wClient *gorm.DB

	sync.RWMutex
}

const (
	//库定义
	DBRBAC = "dayan_rbac"

	// 表定义
	CRole = "role"
	CAccess = "access"
	CRoleAccess  = "role_access"
	CUserRole = "user_role"
)

var (
	defaultDB = &DB{}
	dbArgsFormat = "host=%s dbname=%s sslmode=disable user=%s  password=%s"
)

func Init() {
	var err error
	rArgs := fmt.Sprintf(dbArgsFormat,
		util.DefaultPgSQLConf.ReadAddr,
		DBRBAC,
		util.DefaultPgSQLConf.Username,
		util.DefaultPgSQLConf.Password)
	db, err := gorm.Open("postgres", rArgs)
	if err != nil {
		logrus.Fatalf("open postgres error:%v", err)
	}

	defaultDB.rClient = db
}
