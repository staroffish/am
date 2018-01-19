package view

import (
	"gopkg.in/mgo.v2/bson"
	"db"
	"fmt"
	"global"
	"html/template"
	"net/http"
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
				<input type="text" id="searchInput"/>
				<input type="button" onclick="javascript:search_collection('{{.PrePage}}')" value="查询"/>
				<input type="button" onclick="javascript:add_anime('show_collection(\'{{.PrePage}}\')')" value="新建"/>
				<input type="button" value="返回" onclick="javascript:{{.PrePage}}"/>
			</td>
		</tr>
	</table>
    <table width="96%" border="1" align="center">
    <tr align="left">
        <th>动漫名</th>
        <th>操作</th>
	</tr>
	{{range .Anime}}
    <tr align="left">
		<td style="width:91%;">
			<a href="javascript:void(0)" onclick="javascript:show_anime('{{.AnimeID}}','show_collection(\'{{$.PrePage}}\')')">{{.AnimeNameJp}}</a>
		</td>
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
type ShowCollectionPage struct {
	tmpl *template.Template
}

type showCollection struct {
	Anime   []db.Anime
	PrePage template.JS
}

func (s *ShowCollectionPage) Init() error {
	defer global.TraceLog("ShowCollectionPage.Init")()
	tmp, err := template.New("collection").Parse(pageCollection)
	if err != nil {
		return fmt.Errorf("Init:template.New:%v", err)
	}
	s.tmpl = tmp
	return nil
}

// 动漫显示方法
func (s *ShowCollectionPage) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("ShowCollectionPage.ShowPageCtx")()
	if len(req.Params) < 1 ||
		(len(req.Params) < 2 && req.Method == "search_collection"){
		return fmt.Errorf("ShowCollectionPage:ShowPageCtx:Parameter Num is invalid")
	}

	prePage, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("ShowCollectionPage:ShowPageCtx:Parameter type error:prePage:%T", req.Params[0])
	}
	var key = bson.M{}
	if req.Method != "show_collection"{
		keyword, ok := req.Params[1].(string)
		if !ok {
			return fmt.Errorf("ShowCollectionPage:ShowPageCtx:Parameter type error:keyword:%T", req.Params[0])
		}
		if len(keyword) > 0 {
			key = bson.M{"$or":[]bson.M{
				bson.M{"animenamecn":bson.RegEx{Pattern:keyword}},
				bson.M{"animenamejp":bson.RegEx{Pattern:keyword}},
				bson.M{"cast":bson.RegEx{Pattern:keyword}},
				bson.M{"serialsduri":bson.RegEx{Pattern:keyword}},
				bson.M{"type":bson.RegEx{Pattern:keyword}},
				bson.M{"status":bson.RegEx{Pattern:keyword}},
				}}
		}
	}

	var col showCollection
	col.Anime = db.GetAnimeByKey(key, 0)
	if col.Anime == nil {
		// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
		retStr := fmt.Sprintf(`<html><body>动漫未找到<input type="button" value="返回" onclick="javascript:%s"/></body></html>`, prePage)
		http.Error(w, retStr, http.StatusNotFound)
		return nil
	}
	col.PrePage = template.JS(prePage)

	if err := s.tmpl.Execute(w, col); err != nil {
		return fmt.Errorf("ShowCollectionPage:template.Execute:%v", err)
	}

	return nil
}
