package view

import (
	"fmt"
	"html/template"

	"github.com/staroffish/am/global"
)

var pageCollection = `
<script language='javascript' src='js/event.js' ></script>
<html>
<head>
    <title>已收藏动漫</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
    <table width="96%" border="0" align="center"><tr><td><h1>已收藏动漫</h1></td></tr></table>
	<table width="96%" border="0" align="center">
		<tr>
			<td>
				<input type="text" value="{{.Key}}" id="searchInput"/>
				<input type="button" onclick="javascript:search_collection()" value="查询"/>
				<input type="button" onclick="javascript:edit_anime('','show_collection()')" value="新建"/>
				<input type="button" value="返回" onclick="javascript:main()"/>
			</td>
		</tr>
	</table>
    <table width="96%" border="1" align="center">
    <tr align="left">
		<th>日文名</th>
		<th>中文名</th>
		<th>动漫ID</th>
		<th>更新日期</th>
        <th>操作</th>
	</tr>
	{{range .Anime}}
    <tr align="left">
		<td>
			<a href="javascript:void(0)" onclick="javascript:show_anime('{{.AnimeID}}','show_collection()')">{{.AnimeNameJp}}</a>
		</td>
		<td>
			{{.AnimeNameCn}}
		</td>
		<td>{{.AnimeID}}</td>
		<td>{{.UpdateTime|showDate}}</td>
		<td>
			<input type="button" onclick="javascript:del_anime({{.AnimeID}})" value="删除"/>
		</td>
	</tr>
	{{end}}
    </table>
</body>

</html>
`

// 动漫显示 结构体
type AnimeCollectionPage struct {
	CommonPage
}

func (s *AnimeCollectionPage) Init() error {
	defer global.TraceLog("AnimeCollectionPage.Init")()
	tmp, err := template.New("collection").
		Funcs(template.FuncMap{"showDate": global.FormatTime}).
		Parse(pageCollection)
	if err != nil {
		return fmt.Errorf("Init:template.New:%v", err)
	}
	s.tmpl = tmp
	return nil
}
