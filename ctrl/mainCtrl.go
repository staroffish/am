package ctrl

import (
	"fmt"
	"net/http"

	"github.com/staroffish/am/global"

	"github.com/staroffish/am/db"
	"github.com/staroffish/am/view"

	"gopkg.in/mgo.v2/bson"
)

// MainCtrl 主页逻辑控制
type MainCtrl struct {
	view.MainPage
}

// Init - 初始化
func (m *MainCtrl) Init() error {
	return m.MainPage.Init()
}

func (m *MainCtrl) Close() {}

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
