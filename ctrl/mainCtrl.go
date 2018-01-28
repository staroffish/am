package ctrl

import (
	"db"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"view"
)

// MainCtrl 主页逻辑控制
type MainCtrl struct {
	view.MainPage
}

// Init - 初始化
func (m *MainCtrl) Init() error {
	return m.MainPage.Init()
}

// Process 处理函数
func (m *MainCtrl) Process(_ *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("MainCtrl.Process")()
	aniLst := db.GetAnimeByKey(bson.M{}, 30)

	err := m.ShowPage(aniLst, w)
	if err != nil {
		return fmt.Errorf("MainCtrl.Process:%v", err)
	}

	return nil
}
