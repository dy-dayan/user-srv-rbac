package main

import (
	"github.com/dy-dayan/common-srv-atomicid/idl/dayan/user/srv-rbac"
	"github.com/dy-dayan/user-srv-rbac/idl/dayan/user/srv-rbac"
	"github.com/dy-dayan/user-srv-rbac/dal/db"
	"github.com/dy-dayan/user-srv-rbac/handler"
	"github.com/dy-dayan/user-srv-rbac/util/config"
	"github.com/dy-gopkg/kit/micro"
	"github.com/sirupsen/logrus"
)

func main() {
	micro.Init()

	// 初始化配置
	uconfig.Init()

	// 初始化数据库
	db.Init()

	//TODO 初始化缓存
	//cache.CacheInit()

	err := dayan_user_srv_rbac.RegisterRBACHandler(micro.Server(), &handler.Handler{})
	if err != nil {
		logrus.Fatalf("RegisterPassportHandler error:%v", err)
	}

	micro.Run()
}