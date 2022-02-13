package service

import (
	"IM-gossip-space/model"
	"errors"
	"fmt"
	"log"
	"xorm.io/xorm"
)

const (
	driverName = "mysql"
	dsName     = "root:19990429@(127.0.0.1:3306)/chat?charset=utf8"
	showSQL    = true
	maxCon     = 10
	NONERROR   = "noerror" //没有错误
)

var DbEngine *xorm.Engine

func init() {
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, dsName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)      // show sql
	DbEngine.SetMaxOpenConns(2) // set max conn count
	// automatic user constructor
	DbEngine.Sync2(new(model.User))
	fmt.Println("Init database OK")
}
