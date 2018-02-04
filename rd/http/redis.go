package http

import (
	"fmt"
	"global"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type redisPool struct {
	*redis.Pool
}

// Init - 初始化redis连接池
func (r *redisPool) initRedis(ipAddr string, pwd string, db int, maxIdle int) {
	defer global.TraceLog("redis.Init")()

	r.Pool = redis.NewPool(
		func() (redis.Conn, error) {
			optPwd := redis.DialPassword(pwd)
			optDb := redis.DialDatabase(db)
			c, err := redis.Dial("tcp", ipAddr, optPwd, optDb)
			if err != nil {
				return nil, fmt.Errorf("redis.Dial err:%v", err)
			}
			return c, nil
		}, maxIdle)
	r.Pool.MaxActive = maxIdle * 2
}

func (r redisPool) getAllTask(tasks map[string]*httpTask) error {
	defer global.TraceLog("redis.getAllTask")()
	conn := r.Get()
	defer conn.Close()
	var index int
	var taskIdList []string

	// 取得所有的KEY
	for {
		values, err := redis.Values(conn.Do("SCAN", index))
		if err != nil {
			return fmt.Errorf("redis.values err:%v", err)
		}

		values, err = redis.Scan(values, &index)
		if err != nil {
			return fmt.Errorf("redis.Scan err:%v", err)
		}

		var tempTaskIdList []string
		values, err = redis.Scan(values, &tempTaskIdList)
		if err != nil {
			return fmt.Errorf("redis.Scan err:%v", err)
		}

		// 一个KEY都不存在所以退出循环
		if len(tempTaskIdList) == 0 {
			break
		}

		taskIdList = append(taskIdList, tempTaskIdList...)

		// 遍历完了所有的key 推出循环
		if index == 0 {
			break
		}
	}

	// 由于 内存中的 任务 和 数据库中的任务出现了不同步的问题所以清空内存任务
	if len(taskIdList) != len(tasks) {
		for key := range tasks {
			delete(tasks, key)
		}
	}

	// 取得每个任务
	for _, id := range taskIdList {
		taskMap, err := redis.StringMap(conn.Do("HGETALL", id))
		if err != nil {
			return fmt.Errorf("conn.Do err:%v", err)
		}
		task, ok := tasks[id]
		if !ok {
			task = &httpTask{Controler: make(chan int)}
			tasks[id] = task
		}

		task.Id = id
		task.Path = taskMap["Path"]
		task.FileName = taskMap["FileName"]
		task.Url = taskMap["Url"]
		task.CanTruncate = taskMap["CanTruncate"]
		task.TaskType = taskMap["TaskType"]
		task.State = taskMap["State"]
		task.CreateTime = taskMap["CreateTime"]
		task.TotalSize, err = strconv.ParseFloat(taskMap["TotalSize"], 64)
		if err != nil {
			return fmt.Errorf("TotalSize convert err:%v", err)
		}

		task.FinishedSize, err = strconv.ParseFloat(taskMap["FinishedSize"], 64)
		if err != nil {
			return fmt.Errorf("FinishedSize convert err:%v", err)
		}
	}

	return nil
}
