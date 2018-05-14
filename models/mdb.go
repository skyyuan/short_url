package models

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var (
	db *mgo.Database
)

func GetDB() *mgo.Database {
	if db != nil {
		return db
	}
	session, err := mgo.Dial(beego.AppConfig.String("mongo_url"))
	if err != nil {
		panic(err)
	}
	db = session.DB(beego.AppConfig.String("mongo_database"))
	return db
}
