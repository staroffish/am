// main包

package main

import (
	"ad"
	"db"
	"fmt"
	"global"
	"io"
	"net/http"
	"os"
	"view"
)

var usage = `Usage:am [Config Path]`
var viewMap = map[string]view.Page{"main": &view.MainPage{}}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(-1)
	}
	err := global.NewConfig(os.Args[1])
	if err != nil {
		global.Log.Errorf("am:Read config error:%v", err)
		os.Exit(-1)
	}

	global.NewLogger(global.Cfg.DebugOn)
	global.Log.Infof("Start Program")
	err = db.Connect(global.Cfg)
	if err != nil {
		global.Log.Errorf("am:Connect error:%v", err)
		os.Exit(-1)
	}
	defer db.Close()
	autoDownload := ad.New(global.Cfg)

	autoDownload.Run()
	if err != nil {
		global.Log.Errorf("am:autoDownload.Run:%v", err)
		os.Exit(-1)
	}
	http.HandleFunc("/", handler)
	http.Handle("/js/", http.FileServer(http.Dir("./")))
	global.Log.Infof("%v", http.ListenAndServe(global.Cfg.BindAddr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer global.TraceLog("main.handler")()
	body := make([]byte, 1024)
	n, err := r.Body.Read(body)
	var jReq *view.JSONRequest

	if n == 0 {
		// 请求的是主页
		jReq = &view.JSONRequest{Method: "main"}
	} else if err != io.EOF {
		global.Log.Errorf("am:Body.Read:%v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else {
		// 请求的其他页面
		jReq, err = view.ParsePostData(body[:n])
		if err != nil {
			global.Log.Errorf("am:view.ParsePostData:%v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	page, ok := viewMap[jReq.Method]
	if !ok {
		global.Log.Debugf("am:unsupported method")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = page.ShowPageCtx(jReq, w)
	if err != nil {
		global.Log.Errorf("am:page.ShowPageCtx:%v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
