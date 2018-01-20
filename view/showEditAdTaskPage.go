package view

import (
	"db"
	"fmt"
	"global"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

var pageEditAd = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>编辑</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
	<table align="left" border="0">
		<tr><td>名称</td><td  width="90%"><input style="width:100%;" type="text" value="{{.AnimeNameJp}}" /></td></tr>
		<tr><td>抓取网页地址</td><td  width="90%"><input style="width:100%;" type="text" value="{{.Url}}" /></td></tr>
		<tr><td>抓取网页条件</td><td  width="90%"><input style="width:100%;" type="text" value="{{.UrlParam}}" /></td></tr>
        <tr><td>抓取网页正则</td><td><input style="width:100%;" type="text" value="{{.SchExp}}" /></td></tr>
		<tr><td>抓取磁链正则</td><td><input style="width:100%;" type="text" value="{{.MagExp}}" /></td></tr>
		<tr><td>最新集数</td><td><input style="width:100%;" type="text" value="{{.SchChapt}}" /></td></tr>
        <tr><td>储存地址</td><td><input type="text" {{.Disabled}} style="width:100%;" value="{{.StorDir}}" /></td></tr>
		<tr><td>播放地址</td><td><input type="text" {{.Disabled}} style="width:100%;" value="{{.PlayDir}}" /> </td></tr>
		<tr><td>收藏动漫ID</td><td><input type="text" style="width:100%;" value="{{.AnimeID}}" /> </td></tr>
		{{.CheckBox}}
        <tr><td></td><td align="right"><input type="button" onclick="javascript:update_adTask({{.Id}})" value="提交"><input type="button" onclick="javascript:show_adtask()" value="返回"></td></tr>
    </table>
</body>`

// AdPage - Main page struct
type ShowEditAdTaskPage struct {
	tmpl *template.Template
}

// Init - init mainpage
func (s *ShowEditAdTaskPage) Init() error {
	defer global.TraceLog("ShowEditAdTaskPage.Init")()
	if s.tmpl == nil {
		tmp, err := template.New("autodownload").
			Parse(pageEditAd)
		if err != nil {
			return fmt.Errorf("ShowEditAdTaskPage:template.New:%v", err)
		}
		s.tmpl = tmp
	}
	return nil
}

// ShowPageCtx - 显示页面
func (s *ShowEditAdTaskPage) ShowPageCtx(jr *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("ShowEditAdTaskPage.ShowPageCtx")()

	type taskData struct {
		Id          string
		AnimeID     string
		AnimeNameJp string
		Url         string
		UrlParam    string
		SchExp      string
		SchChapt    int
		MagExp      string
		StorDir     string
		PlayDir     string
		CheckBox    template.HTML
		Disabled	template.HTMLAttr
	}
	var td taskData

	td.SchChapt = 1
	// 添加自动下载
	if jr.Method == "edit_adtask" {
		//显示编辑页面
		id, ok := jr.Params[0].(string)
		if !ok {
			return fmt.Errorf("ShowEditAdTaskPage:show edit page:Id type error:%T", jr.Params[0])
		}
		tasks := db.GetAdTaskByKey(bson.M{"_id": bson.ObjectIdHex(id)})
		if tasks == nil {
			return fmt.Errorf("ShowEditAdTaskPage:tasks not found:id:%s", id)
		}
		td.Id = tasks[0].Id.Hex()
		td.Url = tasks[0].Url
		td.UrlParam = tasks[0].UrlParam
		td.SchExp = tasks[0].SchExp
		td.SchChapt = tasks[0].SchChapt
		td.MagExp = tasks[0].MagExp
		td.AnimeID = tasks[0].AnimeID.Hex()

		anime := db.GetAnimeByKey(bson.M{"_id": tasks[0].AnimeID}, 0)
		if anime == nil {
			return fmt.Errorf("ShowEditAdTaskPage:anime of tasks not found:id:%s", tasks[0].AnimeID.Hex())
		}
		td.AnimeNameJp = anime[0].AnimeNameJp
		td.StorDir = anime[0].StorDir
		td.PlayDir = anime[0].PlayDir
		td.Disabled = "disabled"
	} else {
		td.CheckBox = template.HTML(`<tr><td>是否创建动漫选项</td><td><input type="checkbox" checked="true" /> </td></tr>`)
	}

	if err := s.tmpl.Execute(w, &td); err != nil {
		return fmt.Errorf("ShowEditAdTaskPage:template.Execute:%v", err)
	}
	return nil
}
