package main

import (
	"fmt"
	"testing"

	"git.daynightplay.cn/kepler/backend/core/kepler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gogap/logrus"
)

const (
	user      = "kepler"
	password  = "amo6662018"
	host      = "123.56.27.55"
	port      = 3306
	charset   = "utf8"
	location  = "Asia/Shanghai"
	showSQL   = true
	isDefault = true
	db        = "kepler"
)

// NewXormEngine 测试
func TestNewXormEngine(t *testing.T) {
	var (
		engine *xorm.Engine
		// tables []*core.Table
		err error
	)
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s", user, password, host, port, db, charset))

	if err != nil {
		logrus.Errorln("new engine err: ", err)
		// return
	}
	engine.ShowSQL(showSQL)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	// tables, err = engine.DBMetas()
	// if err != nil {
	// 	logrus.Errorln("DBMetas err: ", err)
	// 	return
	// }
	// for i, v := range tables {
	// 	fmt.Println(i, v)
	// }
	var stars []*kepler.Star
	err = engine.Alias("s").Where("id>1").Limit(100).Find(&stars)
	if err != nil {
		logrus.Errorln("find star err: ", err)
	}
	logrus.Info(stars)
	// for i, v := range stars {
	// 	fmt.Println(i, v)
	// }
	return
}
