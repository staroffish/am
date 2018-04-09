package view

import (
	"fmt"
	"html/template"

	"github.com/staroffish/am/global"
)

var pageEditAnime = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>编辑</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
    <table align="left" border="0">
        <tr><td>中文名</td><td  width="90%"><input style="width:100%;" type="text" value="{{.Anime.AnimeNameCn}}" /></td></tr>
        <tr><td>日文名</td><td><input style="width:100%;" type="text" value="{{.Anime.AnimeNameJp}}" /></td></tr>
        <tr><td>声优</td><td><input style="width:100%;" type="text" value="{{.Anime.Cast}}" /></td></tr>
		<tr><td>类型</td><td><input type="text" style="width:100%;" value="{{.Anime.Type}}" /></td></tr>
		<tr><td>连载状态</td><td><input type="text" style="width:100%;" value="{{.Anime.Status}}" /></td></tr>
		<tr><td>连载期间</td><td><input type="text" style="width:100%;" value="{{.Anime.SerialsDuri}}" /> </td></tr>
		<tr><td>存储路径</td><td><input type="text" style="width:100%;" value="{{.Anime.StorDir}}" /> </td></tr>
		<tr><td>播放路径</td><td><input type="text" style="width:100%;" value="{{.Anime.PlayDir}}" /> </td></tr>
		<tr><td>上传图片</td><td><input type="text" style="width:100%;" value="" /> </td></tr>
		<tr>
			<td></td>
			<td align="right">
				<input type="button" onclick="javascript:show_anime('{{.Anime.AnimeID}}','{{.PrePage}}')" value="返回">
			</td>
			<td align="right">
				<input type="button" onclick="javascript:update_anime('{{.Anime.AnimeID}}','{{.PrePage}}')" value="提交">
			</td>
		</tr>
    </table>
</body>
</html>
`

// EditAnimePage - 动漫编辑 结构体
type EditAnimePage struct {
	CommonPage
}

func (e *EditAnimePage) Init() error {
	defer global.TraceLog("EditAnimePage.Init")()
	if e.tmpl == nil {
		tmp, err := template.New("anime").Parse(pageEditAnime)
		if err != nil {
			return fmt.Errorf("main:template.New:%v", err)
		}
		e.tmpl = tmp
	}
	return nil
}
