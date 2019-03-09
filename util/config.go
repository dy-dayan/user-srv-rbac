package util

import (
	"github.com/sirupsen/logrus"
	"github.com/micro/go-config"
)

type (
	PgSQLConfig struct {
		ReadAddr string `json:"readAddr"`
		WriteAddr string `json:"writeAddr"`
		Username string `json:"username"`
		Password string `json:"password"`
		ReadTimeout int `json:"readTimeout"`
		WriteTimeout int `json:"writeTimeout"`
	}
)

var (
	DefaultPgSQLConf PgSQLConfig
)

func Init() {
	// 加载db配置
	err := config.Get("pgsql").Scan(&DefaultPgSQLConf)
	if err != nil {
		logrus.Fatalf("get mgo config error: %s", err)
	}

	if len(DefaultPgSQLConf.ReadAddr) == 0 {
		logrus.Fatalf("invalid pgsql addr")
	}

}