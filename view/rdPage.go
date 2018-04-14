package view

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/staroffish/am/global"
	"github.com/staroffish/am/rd"
	"golang.org/x/net/websocket"
)

var pageRd = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>离线下载</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body onload="get_task_by_ws()">
    <table width="96%" border="0" align="center"><tr><td><h1>离线下载</h1></td></tr></table>
    <table width="96%" border="1" align="center">
        <tr>
            <td width="18%">下载地址(磁链/HTTP下载)</td>
            <td><input style="width:82%;" type="text" name="" id="magnet"/></td>
        </tr>
        <tr>
             <td>存储路径</td>
            <td><input style="width:82%;" type="text" name="" id="savePath"/>
                <input type="button" style="width:6%;" onclick="javascript:add_task()" value="提交">
				<input type="button" style="width:6%;" onclick="location='/'" value="返回">
			</td>
        </tr>
    </table>
    <div id="tasks"></div>
</body>
</html>`

var pageTask = `<table width="96%" border="1" align="center">
<tr align="left" >
	<th>名称</th>
	<th>大小</th>
	<th>存储路径</th>
	<th width="5%">进度</th>
	<th>状态</th>
	<th>创建时间</th>
	<th>操作</th>
</tr>
{{range .}}
<tr align="left" >
	<td>{{.Name}}</td>
	<td>{{.Size|SizeToString}}</td>
	<td>{{.SavePath}}</td>
	<td>{{.Progress}}%</td>
	<td>{{.State}}</td>
	<td>{{.CreateTime}}</td>
	<td><input type="button" style="width:50%;" {{.State|GetPauseBtnDisnable}} onclick="javascript:{{.State|GetPauseBtnEvent}}'{{.Ids}}:{{.TaskType}}')" value="{{.State|GetPauseBtnValue}}"><input type="button" style="width:50%;" onclick="del_task('{{.Ids}}:{{.TaskType}}')" value="删除"></td>
</tr>
{{end}}
</table>`

// MainPage - Main page struct
type RdPage struct {
	CommonPage
	Head []byte
}

func getPauseBtnValue(status string) string {
	if status == rd.Paused || status == rd.Error || status == rd.Finished {
		return "开始"
	}
	return "暂停"
}

func getPauseBtnEvent(status string) template.JS {
	if status == rd.Paused || status == rd.Error || status == rd.Finished {
		return "start_task("
	}
	return "pause_task("
}

func getPauseBtnDisnable(status string) template.JS {
	if status == rd.Finished {
		return "disabled"
	}
	return ""
}

// Init - init mainpage
func (r *RdPage) Init() error {
	defer global.TraceLog("RdPage.Init")()
	r.Head = []byte(pageRd)
	if r.tmpl == nil {
		tmp, err := template.New("remotedownload").
			Funcs(template.FuncMap{"GetPauseBtnValue": getPauseBtnValue}).
			Funcs(template.FuncMap{"SizeToString": global.SizeToString}).
			Funcs(template.FuncMap{"GetPauseBtnEvent": getPauseBtnEvent}).
			Funcs(template.FuncMap{"GetPauseBtnDisnable": getPauseBtnDisnable}).
			Parse(pageTask)
		if err != nil {
			return fmt.Errorf("RdPage:template.New:%v", err)
		}
		r.tmpl = tmp
	}
	return nil
}

func (r *RdPage) PushPageCtx(ws *websocket.Conn, task []rd.RdTask) error {
	defer global.TraceLog("RdPage.PushPageCtx")()
	buff := bytes.NewBuffer([]byte{})
	if err := r.tmpl.Execute(buff, task); err != nil {
		return fmt.Errorf("PushPageCtx error:%v", err)
	}

	return websocket.Message.Send(ws, buff.String())
}
