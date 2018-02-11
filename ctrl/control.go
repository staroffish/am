package ctrl

import (
	"net/http"
)

// Control - 逻辑层的基类
type Control interface {
	Init() error
	Process(*JSONRequest, http.ResponseWriter) error
	Close()
}
