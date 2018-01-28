// main包

package main

import (
	"db"
	"ad"
	"fmt"
	"global"
	"io"
	"net/http"
	"os"
	// "view"
	"ctrl"
)

var usage = `Usage:am [Config Path]`
var ctrlMap map[string]ctrl.Control

func init() {
	ctrlMap = make(map[string]ctrl.Control)
	ctrlMap["main"] = &ctrl.MainCtrl{}
	anime := &ctrl.AnimeCtrl{}
	ctrlMap["show_collection"] = anime
	ctrlMap["show_anime"] = anime
	ctrlMap["edit_anime"] = anime
	ctrlMap["update_anime"] = anime
	ctrlMap["del_anime"] = anime
	rdCtl := &ctrl.RdCtrl{}
	ctrlMap["get_task"] = rdCtl
	ctrlMap["pause_task"] = rdCtl
	ctrlMap["start_task"] = rdCtl
	ctrlMap["del_task"] = rdCtl
	ctrlMap["add_task"] = rdCtl
	adCtl := &ctrl.AdCtrl{}
	ctrlMap["show_adtask"] = adCtl
	ctrlMap["edit_adtask"] = adCtl
	ctrlMap["update_adTask"] = adCtl
	ctrlMap["delete_adTask"] = adCtl
}
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

	go autoDownload.Run()
	// for name, page := range viewMap {
	// 	err := page.Init()
	// 	if err != nil {
	// 		global.Log.Errorf("init %s page map err:%v", name, err)
	// 	}
	// }
	for name, ctrl := range ctrlMap {
		err := ctrl.Init()
		if err != nil {
			global.Log.Errorf("init %s ctrl map err:%v", name, err)
		}
	}
	http.HandleFunc("/", handler)
	http.Handle("/js/", http.FileServer(http.Dir("./")))
	http.Handle("/usb/", http.FileServer(http.Dir("/usb/")))
	global.Log.Infof("%v", http.ListenAndServe(global.Cfg.BindAddr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer global.TraceLog("main.handler")()
	body := make([]byte, 8192)
	n, err := r.Body.Read(body)
	var jReq *ctrl.JSONRequest

	if n == 0 {
		// 请求的是主页
		// jReq = &view.JSONRequest{Method: "main"}
		jReq = &ctrl.JSONRequest{Method: "main"}
	} else if err != io.EOF {
		global.Log.Errorf("am:Body.Read:%v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else {
		// 请求的其他页面
		// jReq, err = view.ParsePostData(body[:n])
		jReq, err = ctrl.ParsePostData(body[:n])
		if err != nil {
			global.Log.Errorf("am:view.ParsePostData:%v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	// page, ok := viewMap[jReq.Method]
	ctrl, ok := ctrlMap[jReq.Method]
	if !ok {
		global.Log.Debugf("am:unsupported method")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// err = page.ShowPageCtx(jReq, w)
	err = ctrl.Process(jReq, w)
	if err != nil {
		global.Log.Errorf("am:page.ShowPageCtx:%v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
