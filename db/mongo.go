package db

import (
	"time"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	"math"

	mgo "gopkg.in/mgo.v2"
)

var dbSess *mgo.Session
var DB *mgo.Database

func Connect(config *global.Config) error {
	defer global.TraceLog("db.Connect")()
	sess, err := mgo.Dial(config.DBIP + ":" + config.DBPort)
	if err != nil {
		
		return fmt.Errorf("mgo.Dial:%v", err)
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
	anime.AnimeID = anime.ID.Hex()
	defer it.Close()

	return &anime
}

// GetAllAnime - 取得所有动漫动漫
//		cnt - 取的件数(小于等于0时取得所有)
func GetAnimeByKey(key bson.M, cnt int) []Anime {
	defer global.TraceLog("db.GetAllAnime")()
	var aniLst = make([]Anime, 0)
	it := DB.C("anime").Find(key).Sort("-updatetime").Iter()

	if cnt <= 0 {
		cnt = math.MaxInt32
	}

	for i := 0; i < cnt; i++ {
		var anime Anime
		if !it.Next(&anime) {
			if err := it.Err(); err != nil {
				global.Log.Errorf("am:db.GetAllAnime:it.Next:%v", err)
				return nil
			}
			break
		}
		anime.AnimeID = anime.ID.Hex()
		aniLst = append(aniLst, anime)
	}
	defer it.Close()

	return aniLst
}

// SaveAnime - 更新或者插入一个动漫
func SaveAnime(a *Anime) error {
	defer global.TraceLog("db.SaveAnime")()
	oldAni := GetAnime(a.ID)
	if oldAni != nil {
		if a.ImageBin == nil {
			a.ImageBin = oldAni.ImageBin
		}
	}

	a.UpdateTime = time.Now()
	c := DB.C("anime")
	_, err := c.UpsertId(a.ID, &a)
	if err != nil {
		a.ImageBin = nil
		return fmt.Errorf("db.SaveAnime:UpsertId error:%v:%q", err, a)
	}

	return nil
}

// DeletedAnime - 删除一个动画
func DeletedAnime(id bson.ObjectId) error {
	defer global.TraceLog("db.DeletedAnime")()

	c := DB.C("anime")
	err := c.Remove(bson.M{"_id":id})
	if err != nil {
		return fmt.Errorf("db.DeletedAnime:UpsertRemoveId error:%v:%q", err, id.Hex())
	}

	return nil
}

// GetAdTaskByKey - 根据ID取得自动下载任务
func GetAdTaskByKey(key bson.M) []AdTask{
	defer global.TraceLog("db.GetAdTaskByKey")()
	var adTasks = make([]AdTask, 0)
	it := DB.C("adtask").Find(key).Sort("-updatetime").Iter()

	for {
		var adtask AdTask
		if !it.Next(&adtask) {
			if err := it.Err(); err != nil {
				global.Log.Errorf("am:db.GetAdTaskByKey:it.Next:%v", err)
				return nil
			}
			break
		}
		// adtask.Id = bson.ObjectId(adtask.Id).Hex()
		adTasks = append(adTasks, adtask)
	}
	defer it.Close()

	return adTasks
}

// SaveAdTask - 更新或者插入一个自动下载任务
func SaveAdTask(a *AdTask) error {
	defer global.TraceLog("db.SaveAdTask")()

	a.UpdateTime = time.Now()
	c := DB.C("adtask")
	_, err := c.UpsertId(a.Id, a)
	if err != nil {
		return fmt.Errorf("db.SaveAdTask:UpsertId error:%v:%q", err, a)
	}

	return nil
}

// DeletedTask - 删除一个自动下载任务
func DeleteAdTask(id bson.ObjectId) error {
	defer global.TraceLog("db.DeletedTask")()

	c := DB.C("adtask")
	err := c.Remove(bson.M{"_id":id})
	if err != nil {
		return fmt.Errorf("db.DeletedTask:Remove error:%v:%q", err, id.Hex())
	}

	return nil
}