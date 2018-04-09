package view

import (
	"fmt"
	"html/template"

	"github.com/staroffish/am/global"
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
type EditAdTaskPage struct {
	CommonPage
}

// Init - init mainpage
func (s *EditAdTaskPage) Init() error {
	defer global.TraceLog("EditAdTaskPage.Init")()
	if s.tmpl == nil {
		tmp, err := template.New("autodownload").
			Parse(pageEditAd)
		if err != nil {
			return fmt.Errorf("EditAdTaskPage:template.New:%v", err)
		}
		s.tmpl = tmp
	}
	return nil
}
