package service

import (
	"github.com/unrolled/render"
	"net/http"
)

type handles struct {
	Handle []Handle
}

type Service func(w http.ResponseWriter, r *http.Request)

type Handle struct {
	Url     string
	Service Service
}

var Render *render.Render

func Router(r *render.Render, mux *http.ServeMux) {
	Render = r

	router := handles{
		[]Handle{
			{"/api/user", UserPage},
		},
	}

	for _, handle := range router.Handle {
		mux.HandleFunc(handle.Url, handle.Service)
	}

}
