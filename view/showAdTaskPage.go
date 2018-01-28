package view

import (
	"fmt"
	"global"
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
    <div style="text-align:right;width:98%"><input type="button" style="width:8%;" onclick="javascript:edit_adtask('')" value="添加新的自动下载"><input type="button" onclick="javascript:main()" style="width:8%;" value="返回"></div>
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

// AdTaskPage - AdTaskPage struct
type AdTaskPage struct {
	CommonPage
}

// Init - init mainpage
func (a *AdTaskPage) Init() error {
	defer global.TraceLog("AdTaskPage.Init")()
	if a.tmpl == nil {
		tmp, err := template.New("autodownload").
						Funcs(template.FuncMap{"showDate":global.FormatTime}).
						Parse(pageAd)
		if err != nil {
			return fmt.Errorf("AdTaskPage:template.New:%v", err)
		}
		a.tmpl = tmp
	}
	return nil
}