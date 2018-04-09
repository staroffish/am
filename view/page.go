package view

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/staroffish/am/global"
)

// Page - The interface of web page
type Page interface {
	Init() error
	ShowPage(interface{}, http.ResponseWriter) error
}

// CommonPage - 公用页面类
type CommonPage struct {
	tmpl *template.Template
}

// ShowPage - 显示页面
func (c *CommonPage) ShowPage(data interface{}, w http.ResponseWriter) error {
	defer global.TraceLog("RdPage.ShowPage")()

	if err := c.tmpl.Execute(w, data); err != nil {
		return fmt.Errorf("RdPage.ShowPage:template.Execute:%v", err)
	}

	return nil
}
