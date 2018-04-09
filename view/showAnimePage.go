package view

import (
	"fmt"
	"html/template"

	"github.com/staroffish/am/global"
)

var pageAnime = `
<script language='javascript' src='js/event.js' ></script>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>{{.Anime.AnimeNameJp}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
    <table  border="0">
        <tr>
			<td width="10%" height="10%"><img style="display:block;width:100%;" src="data:image/jpg;base64,{{.Anime.ImageBin|base64}}"></img></td>
            <td  valign="top" align="left">
                <table align="left" border="0">
                    <tr><td>中文名</td><td>{{.Anime.AnimeNameCn}}</td></tr>
                    <tr><td>日文名</td><td>{{.Anime.AnimeNameJp}}</td></tr>
                    <tr><td>主要声优</td><td>{{.Anime.Cast}}</td></tr>
                    <tr><td>类型</td><td>{{.Anime.Type}}</td></tr>
					<tr><td>连载时间</td><td>{{.Anime.SerialsDuri}} </td></tr>
					<tr><td>连载状态</td><td>{{.Anime.Status}} </td></tr>
					<tr><td>存储路径</td><td>{{.Anime.StorDir}} </td></tr>
					<tr><td>播放路径</td><td>{{.Anime.PlayDir}} </td></tr>
                    <tr><td><a href="javascript:void(0)" onclick='{{.PrePage}}'>返回</a>&nbsp;<a href="javascript:void(0)" onclick="javascript:edit_anime('{{.Anime.AnimeID}}','{{.PrePage}}')">编辑</a></td><td></td></tr>
                </table>
            </td>
        </tr>
    </table>
	<table border="0">
	{{range .Chapter}}
		<tr>
		<td><a href="{{.FullPath}}">{{.FileName}}</a></td>
		</tr>
	{{end}}
    </table>
</body>`

// 动漫显示 结构体
type AnimePage struct {
	CommonPage
}

func (a *AnimePage) Init() error {
	defer global.TraceLog("AnimePage.Init")()
	if a.tmpl == nil {
		tmp, err := template.New("anime").Funcs(template.FuncMap{"base64": global.EncodeBase64}).Parse(pageAnime)
		if err != nil {
			return fmt.Errorf("main:template.New:%v", err)
		}
		a.tmpl = tmp
	}
	return nil
}
