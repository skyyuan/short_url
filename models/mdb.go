package models

import (
	"gopkg.in/mgo.v2"
	"short_url/conf"
)

var (
	db *mgo.Database
)

func GetDB() *mgo.Database {
	if db != nil {
		return db
	}
	url, err1 := conf.Read("mongo_url")
	database, err2 := conf.Read("mongo_database")
	if err1 != nil || err2 != nil {
		panic(err1)
		panic(err2)
	}
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	db = session.DB(database)
	return db
}
