package httpserver

import (
	"net/http"
	"time"
)

type HttpServer struct{
	mux *http.ServeMux
}

func NewHttpServer() HttpServer{
	return HttpServer{
		mux:http.NewServeMux(),
	}
}


func (this *HttpServer) Run(cfg Config) error{
	server := &http.Server{

		Addr:        cfg.Address,

		ReadTimeout:  cfg.ReadTimeOut * time.Millisecond,

		WriteTimeout: cfg.WriteTimeOut * time.Millisecond,
		Handler: this.mux,
	}
	if cfg.Https{
		return server.ListenAndServeTLS(cfg.CertFile,cfg.KeyFile)
	}else{
		return server.ListenAndServe()
	}
}

func (this *HttpServer) Bind(pattern string,hander func(w http.ResponseWriter,r *http.Request)){
	this.mux.HandleFunc(pattern,hander)
}


