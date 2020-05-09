package admin

import (
	"goweb/internal/http_api"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	Router http.Handler
}

func NewHTTPServer() *httpServer {
	router := httprouter.New()
	s := &httpServer{
		Router: router,
	}

	router.GET("/", http_api.Decorate(s.index, http_api.V1))
	router.POST("/login", http_api.Decorate(s.login, http_api.V1))

	return s
}

func (s *httpServer) index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	return "welcome!", nil
}

func (s *httpServer) login(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	return "welcome, xxxx", nil
}
