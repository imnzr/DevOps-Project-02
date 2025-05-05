package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)

	// Login by email
	Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUsername(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
