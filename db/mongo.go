package db

import (
	"fmt"
	"global"

	mgo "gopkg.in/mgo.v2"
)

var dbSess *mgo.Session
var DB *mgo.Database

func Connect(config *global.Config) error {
	sess, err := mgo.Dial(config.DBIP + ":" + config.DBPort)
	if err != nil {
		return fmt.Errorf("mgo.Dial:%v", err);
	}
	dbSess = sess
	db := dbSess.DB(config.DBName)
	err = db.Login(config.DBUser, config.DBPasswd)
	if err != nil {
		return fmt.Errorf("Login:%v", err)
	}
	DB = db
	return nil
}

func Close() {
	dbSess.Close()
}
