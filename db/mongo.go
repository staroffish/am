package db

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"global"

	mgo "gopkg.in/mgo.v2"
)

var dbSess *mgo.Session
var DB *mgo.Database

func Connect(config *global.Config) error {
	defer global.TraceLog("db.Connect")()
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

// GetAnime - 取得一个动漫
func GetAnime(_id bson.ObjectId) *Anime {
	defer global.TraceLog("db.GetAnime")()
	var anime Anime
	it := DB.C("anime").Find(bson.M{"_id": _id}).Iter()
	succ := it.Next(&anime)
	if succ == false {
		// 已经做了应答出了ERROR LOG 所以 直接返回NIL
		global.Log.Errorf("am:db.GetAnime:No anime Item:AnimeId:%s", _id)
		return nil
	}
	anime.AnimeID =anime.ID.Hex()
	defer it.Close()

	return &anime
}

// SaveAnime - 更新或者插入一个动漫
func SaveAnime(a *Anime) error{
	defer global.TraceLog("db.SaveAnime")()
	oldAni := GetAnime(a.ID)
	if oldAni != nil {
		if len(a.AnimeNameCn) == 0 {
			a.AnimeNameCn = oldAni.AnimeNameCn
		}
		if len(a.AnimeNameJp) == 0 {
			a.AnimeNameJp = oldAni.AnimeNameJp
		}
		if len(a.Cast) == 0 {
			a.Cast = oldAni.Cast
		}
		if len(a.SerialsDuri) == 0 {
			a.SerialsDuri = oldAni.SerialsDuri
		}
		if len(a.Type) == 0 {
			a.Type = oldAni.Type
		}
		if a.ImageBin == nil {
			a.ImageBin = oldAni.ImageBin
		}
		if len(a.StorDir) == 0 {
			a.StorDir = oldAni.StorDir
		}
		if len(a.PlayDir) == 0 {
			a.PlayDir = oldAni.PlayDir
		}
		if len(a.Status) == 0 {
			a.Status = oldAni.Status
		}
		if a.UpdateTime.IsZero() {
			a.UpdateTime = oldAni.UpdateTime
		}
	} 

	c := DB.C("anime")
	_, err := c.UpsertId(a.ID,&a)
	if err != nil {
		a.ImageBin = nil
		return fmt.Errorf("db.SaveAnime:UpsertId error:%v:%q", err, a)
	}

	return nil
}