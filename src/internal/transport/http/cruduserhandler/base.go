package cruduserhandler

import (
	crudusersevice "crud/src/internal/service/crudusersevice"
	"net/http"
)

type HandlerUserCrud interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
type handlerusercrud struct {
	service crudusersevice.ServiceUserCrud
}

type handlerOption func(*handlerusercrud)

func WriteService(service crudusersevice.ServiceUserCrud) handlerOption {
	return func(h *handlerusercrud) {
		h.service = service
	}
}

func New(options ...handlerOption) HandlerUserCrud {
	k := &handlerusercrud{}
	for _, o := range options {
		o(k)
	}
	return k
}
