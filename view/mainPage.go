package view

import (
	"net/http"
	"db"
	"fmt"
	"global"
	"html/template"

	"gopkg.in/mgo.v2/bson"
)


var pageMain = `
<script language='javascript' src='js/event.js' ></script>
<html>
<head>
    <title>动漫管理</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
    <table width="96%" border="0" align="center"><tr><td><h1>动漫管理</h1></td></tr></table>
    <table width="96%" border="0" align="center">
        <tr><td>
        <div style="text-align:left;width:98%">
            <input type="button" style="width:9%;" value="离线下载">
            <input type="button" style="width:9%;" value="自动下载管理">
            <input type="button" style="width:9%;" value="管理已下载动漫">
            <input type="button" style="width:9%;" value="设置">
        </div>
    </td></tr>
    </table>
    <table width="96%" border="1" align="center">
    <tr align="left" >
        <th>最近下载</th>
        <th>更新时间</th>
    </tr>
    {{range .}}
    <tr align="left" >
        <td><a href="javascript:void(0)" onclick="javascript:show_anime('{{.AnimeID}}','main()')">{{.AnimeNameJp}}</a></td>
        <td>{{.UpdateTime|showDate}}</td>
    </tr>
    {{end}}
    </table>
            
</body>
</html>`

// MainPage - Main page struct
type MainPage struct {
	tmpl *template.Template
}

// Init - init mainpage
func (m *MainPage) Init() error {
	defer global.TraceLog("MainPage.Init")()
	if m.tmpl == nil {
		tmp, err := template.New("main").
						Funcs(template.FuncMap{"showDate":global.FormatTime}).
						Parse(pageMain)
		if err != nil {
			return fmt.Errorf("main:template.New:%v", err)
		}
		m.tmpl = tmp
	}
	return nil
}

// ShowPageCtx - 显示页面
func (m *MainPage) ShowPageCtx(_ *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("MainPage.ShowPageCtx")()
	aniLst := make([]db.Anime, 0)
	it := db.DB.C("anime").Find(bson.M{}).Sort("-updatetime").Iter()
	// it := db.DB.C("anime").Find(bson.M{}).Iter()

	for i := 0; i < 30; i++ {
		var anime db.Anime
		if !it.Next(&anime) {
			if err := it.Err(); err != nil {
				return err
			}
			break
		}
		anime.AnimeID = anime.ID.Hex()
		aniLst = append(aniLst, anime)
	}
	defer it.Close()

	if err := m.tmpl.Execute(w, aniLst); err != nil {
		return fmt.Errorf("main:template.Execute:%v", err)
	}

	return nil
}
