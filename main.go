// main包

package main

import (
	"db"
	"fmt"
	"global"
	"io"
	"net/http"
	"os"
	"rd"
	_ "rd/deluge"
	"view"
)

var usage = `Usage:am [Config Path]`
var viewMap = map[string]view.Page{"main": &view.MainPage{},
									"show_anime": &view.ShowAnimePage{},
									"edit_anime": &view.EditAnimePage{},
									"update_anime": &view.EditAnimePage{}}

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
	// autoDownload := ad.New(global.Cfg)

	// autoDownload.Run()
	// if err != nil {
	// 	global.Log.Errorf("am:autoDownload.Run:%v", err)
	// 	os.Exit(-1)
	// }
	rd.InitDownloader()
	// dnldr := deluge.NewDownloader()
	// if err = rd.PauseTask(&rd.RdTask{Ids: "c5d436eb9b05e37ca6025f27f87df8ab3f018cdf", SavePath: "/usb/"},"magnet"); err != nil {
	// 	fmt.Printf("am:%v\n", err)
	// 	os.Exit(-1)
	// }
	// if err = dnldr.PauseTask(&rd.RdTask{Ids:"c5d436eb9b05e37ca6025f27f87df8ab3f018cdf"}); err != nil {
	// 	fmt.Printf("am:%v\n", err);
	// 	os.Exit(-1)
	// }
	// if err = dnldr.DeleteTask(&rd.RdTask{Ids:"8a2e355195bf39add94525e8bdc929433c6a32e7"}); err != nil {
	// 	fmt.Printf("am:%v\n", err);
	// 	os.Exit(-1)
	// }
	// if err = dnldr.ResumeTask(&rd.RdTask{Ids:"c5d436eb9b05e37ca6025f27f87df8ab3f018cdf"}); err != nil {
	// 	fmt.Printf("am:%v\n", err);
	// 	os.Exit(-1)
	// }
	
	// tasks, err := dnldr.GetAllTask()
	// if err != nil {
	// 	fmt.Printf("am:%v\n", err)
	// 	os.Exit(-1)
	// }
	tasks, err := rd.GetAllTask()
	if err != nil {
		fmt.Printf("am:%v\n", err)
		os.Exit(-1)
	}
	for _, t := range tasks {
		fmt.Printf("%v\n", t)
	}
	for name, page := range viewMap {
		err := page.Init()
		if err != nil {
			global.Log.Errorf("init %s page map err:%v", name, err)
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
