// main包

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/staroffish/am/ad"
	"github.com/staroffish/am/ctrl"
	"github.com/staroffish/am/db"
	"github.com/staroffish/am/global"
	_ "github.com/staroffish/am/rd/deluge"
	_ "github.com/staroffish/am/rd/http"
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
		fmt.Fprintf(os.Stderr, "am:Read config error:%v", err)
		os.Exit(-1)
	}

	global.Log.Infof("Start Program")

	workDir, err := os.Getwd()
	if err != nil {
		global.Log.Errorf("am:os.Getwd error:%v", err)
		os.Exit(-1)
	}
	fmt.Printf("dir=%s\n", dir)

	err = db.Connect(global.Cfg)
	if err != nil {
		global.Log.Errorf("am:Connect error:%v", err)
		os.Exit(-1)
	}
	defer db.Close()
	autoDownload, err := ad.New(global.Cfg)
	if err != nil {
		global.Log.Errorf("am:ad.New:%v", err)
		os.Exit(-1)
	}

	go autoDownload.Run()
	for name, ctrl := range ctrlMap {
		err := ctrl.Init()
		if err != nil {
			global.Log.Errorf("am:init %s ctrl map err:%v", name, err)
			os.Exit(-1) // http.HandleFunc("/rdtask", PushTasks)
		}
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/rdtask", ctrlMap["get_task"].(*ctrl.RdCtrl).PushTasks)
	http.Handle("/js/", http.FileServer(http.Dir("./")))
	http.Handle(global.Cfg.SaveDir, http.FileServer(http.Dir(workDir)))
	ln, err := net.Listen("tcp", global.Cfg.BindAddr)
	if err != nil {
		global.Log.Errorf("am:net.Listen error:%v", err)
		os.Exit(-1)
	}

	s := http.Server{}
	s.SetKeepAlivesEnabled(true)
	sigCh := make(chan os.Signal)
	signal.Ignore(syscall.SIGHUP)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		for _, ctrl := range ctrlMap {
			ctrl.Close()
		}

		err := s.Shutdown(nil)
		if err != nil {
			global.Log.Errorf("am:s.Shutdown error:%v", err)
		}
	}()
	err = s.Serve(ln)
	if err != nil && err != http.ErrServerClosed {
		global.Log.Errorf("am:s.Serve error:%v", err)
		os.Exit(-1)
	}

	global.Log.Infof("End program")
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer global.TraceLog("main.handler")()
	var body []byte

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		global.Log.Errorf("am:Body.Read:%v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var jReq *ctrl.JSONRequest

	if len(body) == 0 {
		// 请求的是主页
		// jReq = &view.JSONRequest{Method: "main"}
		jReq = &ctrl.JSONRequest{Method: "main"}
	} else {
		// 请求的其他页面
		jReq, err = ctrl.ParsePostData(body)
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
