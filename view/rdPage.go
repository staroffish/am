package view

import (
	"fmt"
	"global"
	"html/template"
	"rd"
)


var pageRd = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>离线下载</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body onload="reload()">
    <table width="96%" border="0" align="center"><tr><td><h1>离线下载</h1></td></tr></table>
    <table width="96%" border="1" align="center">
        <tr>
            <td width="18%">下载地址(磁链)</td>
            <td><input style="width:82%;" type="text" name="" id="magnet"/></td>
        </tr>
        <tr>
             <td>存储路径</td>
            <td><input style="width:82%;" type="text" name="" id="savePath"/>
                <input type="button" style="width:6%;" onclick="javascript:add_task()" value="提交">
				<input type="button" style="width:6%;" onclick="javascript:main()" value="返回">
			</td>
        </tr>
    </table>
    <table width="96%" border="1" align="center">
    <tr align="left" >
        <th>名称</th>
        <th>大小</th>
        <th>存储路径</th>
		<th width="5%">进度</th>
		<th>状态</th>
        <th>操作</th>
	</tr>
	{{range .}}
    <tr align="left" >
        <td>{{.Name}}</td>
        <td>{{.Size|SizeToString}}</td>
        <td>{{.SavePath}}</td>
		<td>{{.Progress}}%</td>
		<td>{{.State}}</td>
        <td><input type="button" style="width:50%;" onclick="javascript:{{.State|GetPauseBtnEvent}}'{{.Ids}}')" value="{{.State|GetPauseBtnValue}}"><input type="button" style="width:50%;" onclick="del_task('{{.Ids}}')" value="删除"></td>
	</tr>
	{{end}}
    </table>
</body>
</html>`

// MainPage - Main page struct
type RdPage struct {
	CommonPage
}

func GetPauseBtnValue(status string) string {
	if status == rd.Paused {
		return "开始"
	}
	return "暂停"
}

func GetPauseBtnEvent(status string) template.JS {
	if status == rd.Paused {
		return "start_task("
	}
	return "pause_task("
}

// Init - init mainpage
func (r *RdPage) Init() error {
	defer global.TraceLog("RdPage.Init")()
	if r.tmpl == nil {
		tmp, err := template.New("remotedownload").
						Funcs(template.FuncMap{"GetPauseBtnValue":GetPauseBtnValue}).
						Funcs(template.FuncMap{"SizeToString":global.SizeToString}).
						Funcs(template.FuncMap{"GetPauseBtnEvent":GetPauseBtnEvent}).
						Parse(pageRd)
		if err != nil {
			return fmt.Errorf("RdPage:template.New:%v", err)
		}
		r.tmpl = tmp
	}
	return nil
}