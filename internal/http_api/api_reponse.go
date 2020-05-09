package http_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// w http.ResponseWriter, r *http.Request, ps httprouter.Params
// http.ResponseWriter, *http.Request, httprouter.Params

type APIHandler func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, error)

type Decorator func(APIHandler) APIHandler

type Err struct {
	Code int
	Text string
}

func (e Err) Error() string {
	return e.Text
}

func Decorate(f APIHandler, ds ...Decorator) httprouter.Handle {
	decorated := f
	for _, decorate := range ds {
		decorated = decorate(f)
	}
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		decorated(w, req, ps)
	}
}

func V1(f APIHandler) APIHandler {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
		data, err := f(w, req, ps)
		if err != nil {
			RespondV1(w, err.(Err).Code, err)
			return nil, nil
		}
		RespondV1(w, 200, data)
		return nil, nil
	}
}

func RespondV1(w http.ResponseWriter, code int, data interface{}) {
	var response []byte
	var err error
	var isJSON bool

	if code == 200 {
		switch data.(type) {
		case string:
			response = []byte(data.(string))
		case []byte:
			response = data.([]byte)
		case nil:
			response = []byte{}
		default:
			isJSON = true
			response, err = json.Marshal(data)
			if err != nil {
				code = 500
				data = err
			}
		}
	}

	if code != 200 {
		isJSON = true
		response, _ = json.Marshal(struct {
			Message string `json:"message"`
		}{fmt.Sprintf("%s", data)})
	}

	if isJSON {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	w.WriteHeader(code)
	w.Write(response)
}
