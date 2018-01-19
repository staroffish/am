package view

import (
	"gopkg.in/mgo.v2/bson"
	"db"
	"net/http"
	"fmt"
	"global"
	"time"
	"html/template"
)


var pageAd = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>自动下载</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
    <table width="96%" border="0" align="center"><tr><td><h1>自动下载</h1></td></tr></table>
    <div style="text-align:right;width:98%"><input type="button" style="width:8%;" onclick="javascript:add_adtask()" value="添加新的自动下载"><input type="button" onclick="javascript:main()" style="width:8%;" value="返回"></div>
    <table width="96%" border="1" align="center">
    <tr align="left" >
        <th>任务名</th>
        <th>最新集数</th>
		<th>存储路径</th>
		<th>更新时间</th>
        <th>操作</th>
	</tr>
	{{range .}}
    <tr align="left" >
        <td>{{.AnimeNameJp}}</td>
        <td>{{.SchChapt}}</td>
		<td>{{.StorDir}}</td>
		<td>{{.UpdateTime|showDate}}</td>
        <td width="9%" align="right"><input type="button" style="width:50%" onclick="javascript:edit_adtask({{.Id}})" value="编辑"><input type="button" onclick="delete_adTask({{.Id}})" style="width:50%" value="删除"></td>
	</tr>
	{{end}}
    </table>
           
</body>
</html>`

// AdPage - Main page struct
type ShowAdTaskPage struct {
	tmpl *template.Template
}

// Init - init mainpage
func (a *ShowAdTaskPage) Init() error {
	defer global.TraceLog("ShowAdTaskPage.Init")()
	if a.tmpl == nil {
		tmp, err := template.New("autodownload").
						Funcs(template.FuncMap{"showDate":global.FormatTime}).
						Parse(pageAd)
		if err != nil {
			return fmt.Errorf("ShowAdTaskPage:template.New:%v", err)
		}
		a.tmpl = tmp
	}
	return nil
}

// ShowPageCtx - 显示页面
func (a *ShowAdTaskPage) ShowPageCtx(jr *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("ShowAdTaskPage.ShowPageCtx")()

	type taskData struct {
		Id string
		AnimeNameJp string
		SchChapt int
		StorDir string
		UpdateTime time.Time
	}
	tds := []taskData{}

	tasks := db.GetAdTaskByKey(bson.M{});
	if tasks != nil {
		animeList := db.GetAnimeByKey(bson.M{}, 0);
		if animeList != nil {
			var animeMap = make(map[string]*db.Anime)
			for i := 0; i < len(animeList); i++ {
				animeList[i].ImageBin = nil
				animeMap[animeList[i].AnimeID] = &animeList[i]	
			}
			for _, task := range tasks {
				var td taskData
				td.Id = task.Id.Hex()
				td.SchChapt = task.SchChapt
				anime, ok := animeMap[task.AnimeID.Hex()]
				if !ok {
					global.Log.Debugf("am:db.ShowAdTaskPage:animeID in adtask but not in anime:%s", task.AnimeID.Hex())
					continue;
				}
				td.StorDir = anime.StorDir
				td.AnimeNameJp = anime.AnimeNameJp
				td.UpdateTime = task.UpdateTime
				tds = append(tds, td)
			}
		}
	}
	if err := a.tmpl.Execute(w, &tds); err != nil {
		return fmt.Errorf("ShowAdTaskPage:template.Execute:%v", err)
	}

	return nil
}
