package view

import (
	"db"
	"fmt"
	"global"
	"html/template"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
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
				<input type="button" onclick="javascript:update_anime('{{.Anime.AnimeID}}','show_anime(\'{{.Anime.AnimeID}}\',\'{{.PrePage}}\')')" value="提交">
			</td>
		</tr>
    </table>
</body>
</html>
`

// EditAnimePage - 动漫编辑 结构体
type EditAnimePage struct {
	tmpl *template.Template
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

// ShowPageCtx - 编辑添加动漫的页面
func (e *EditAnimePage) ShowPageCtx(req *JSONRequest, w http.ResponseWriter) error {
	defer global.TraceLog("EditAnimePage.ShowPageCtx")()
	var sa  struct {
		Anime   db.Anime
		PrePage template.JS 
	}
	if len(req.Params) == 1 {
		prePage, ok := req.Params[0].(string)
		if !ok {
			return fmt.Errorf("ShowAnime:ShowPageCtx:Parameter type error:%T", req.Params[0])
		}
		sa.PrePage = template.JS(prePage)
		// 新增页面所以没有参数直接返回页面
		if err := e.tmpl.Execute(w, &sa); err != nil {
			return fmt.Errorf("EditAnimePage:ShowPageCtx.Execute:%v", err)
		}
		return nil
	}

	// 取得是修改还是显示编辑页面的FLAG
	// 显示编辑页面
	if len(req.Params) < 3 {
		return fmt.Errorf("EditAnimePage:ShowPageCtx:Parameter Num is less then 2")
	}

	updateFlag, ok := req.Params[2].(bool)
	if !ok {
		return fmt.Errorf("EditAnimePage:ShowPageCtx. parameter convert error %v", req.Params[2])
	}
	prePage, ok := req.Params[1].(string)
	if !ok {
		return fmt.Errorf("ShowAnime:ShowPageCtx:Parameter type error:%T", req.Params[1])
	}

	_id, ok := req.Params[0].(string)
	if !ok {
		return fmt.Errorf("EditAnimePage:ShowPageCtx:Parameter type error:%v", req.Params[0])
	}

	if updateFlag == false {
		ani := db.GetAnime(bson.ObjectIdHex(_id))
		if ani == nil {
			// 已经做了错误应答并出了ERROR LOG 所以 直接返回NIL
			http.Error(w, "Page Not Found", http.StatusNotFound)
			return nil
		}
		sa.Anime = *ani
		sa.PrePage = template.JS(prePage)
		if err := e.tmpl.Execute(w, &sa); err != nil {
			return fmt.Errorf("EditAnimePage:ShowPageCtx.Execute:%v", err)
		}

		return nil
	}

	var ani db.Anime
	for _, param := range req.Params[3:] {
		_, ok := param.(string)
		if !ok {
			return fmt.Errorf("EditAnimePage:ShowPageCtx:Parameter type error:value:%v:type:%T", param, param)
		}
	}

	var i = 3

	// 添加新项目时ani.AnimeID不存在,生成一个新的object id
	if _id == "" {
		_id = bson.NewObjectId().Hex()
		prePage = "show_collection('main()')"
	}
	ani.ID = bson.ObjectIdHex(_id)
	ani.AnimeNameCn, ok = req.Params[i].(string)
	i++
	ani.AnimeNameJp, ok = req.Params[i].(string)
	i++
	ani.Cast, ok = req.Params[i].(string)
	i++
	ani.Type, ok = req.Params[i].(string)
	i++
	ani.Status, ok = req.Params[i].(string)
	i++
	ani.SerialsDuri, ok = req.Params[i].(string)
	i++
	ani.StorDir, ok = req.Params[i].(string)
	i++
	ani.PlayDir, ok = req.Params[i].(string)
	i++
	imageUrl, ok := req.Params[i].(string)
	i++

	// 如果图片路径存在则取得图片
	if imageUrl != "" {
		resp, err := http.Get(imageUrl)
		if err != nil {
			global.Log.Errorf("am:EditAnimePage:ShowPageCtx:Get image error:%v", err)
		} else {
			defer resp.Body.Close()
			ani.ImageBin, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				global.Log.Errorf("am:EditAnimePage:ShowPageCtx:Read image error:%v", err)
			}
		}
	}

	// 更新数据库
	err := db.SaveAnime(&ani)
	if err != nil {
		return fmt.Errorf("EditAnimePage:ShowPageCtx:SaveAnime:%v", err)
	}

	// 返回页面请求拼装
	_, err = w.Write([]byte(prePage))
	if err != nil {
		return fmt.Errorf("EditAnimePage:ShowPageCtx:write:%v", err)
	}

	return nil
}
