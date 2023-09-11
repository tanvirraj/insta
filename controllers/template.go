package controllers

import "net/http"

type Template interface {
	Excute(w http.ResponseWriter, data interface{})
}
