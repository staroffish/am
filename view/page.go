package view

import (
	"net/http"
)

// Page - The interface of web page
type Page interface {
	ShowPageCtx(*JSONRequest, http.ResponseWriter) error
}
