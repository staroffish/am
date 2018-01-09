package view

import (
	"net/http"
)

// Page - The interface of web page
type Page interface {
	Init() error
	ShowPageCtx(*JSONRequest, http.ResponseWriter) error
}
