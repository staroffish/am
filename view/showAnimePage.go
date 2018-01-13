package view

import (
	"gopkg.in/mgo.v2/bson"
	"db"
	"fmt"
	"global"
	"html/template"
	"io/ioutil"
	"net/http"
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
type ShowAnimePage struct {
	tmpl *template.Template
}

type showAnime struct {
	Anime   db.Anime
	Chapter []struct {
		FullPath string
		FileName string
	}
	PrePage template.JS
}

func (a *ShowAnimePage) Init() error {
	defer global.TraceLog("ShowAnimePage.Init")()
	if a.tmpl == nil {
		tmp, err := template.New("anime").Funcs(template.FuncMap{"base64": global.EncodeBase64}).Parse(pageAnime)
		if err != nil {
			return fmt.Errorf("main:template.New:%v", err)
		}
		a.tmpl = tmp
	}
	return nil
}

// 动漫显示方法
func (a *ShowAnimePage) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("ShowAnimePage.ShowPageCtx")()
	if len(req.Params) <= 1 {
		return fmt.Errorf("ShowAnime:ShowPageCtx:Parameter Num is zero")
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("ShowAnime:ShowPageCtx:Parameter type error:%T", req.Params[0])
	}

	var shower showAnime
	ani := db.GetAnime(bson.ObjectIdHex(_id))
	if ani == nil {
		// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return nil
	}

	shower.Anime = *ani
	prePage, ok := req.Params[1].(string) 
	if !ok {
		return fmt.Errorf("ShowAnime:ShowPageCtx:Parameter type error:%T", req.Params[1])
	}
	shower.PrePage = template.JS(prePage)
	// 取得文件夹下面的所有文件
	fileList, err := ioutil.ReadDir(shower.Anime.StorDir)
	if err != nil {
		return fmt.Errorf("ShowAnime:ShowPageCtx:read dir error:%s:%v", shower.Anime.StorDir, err)
	}
	shower.Chapter = []struct {
		FullPath string
		FileName string
	}{}
	for _, file := range fileList {
		playPath := fmt.Sprintf("%s/%s", shower.Anime.PlayDir, file.Name())
		shower.Chapter = append(shower.Chapter,
			struct {
				FullPath string
				FileName string
			}{playPath, file.Name()})
	}

	if err := a.tmpl.Execute(w, &shower); err != nil {
		return fmt.Errorf("ShowAnime:ShowPageCtx.Execute:%v", err)
	}

	return nil
}
