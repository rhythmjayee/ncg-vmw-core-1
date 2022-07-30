package api

import (
	. "DIaaS/diaas/auth"
	. "DIaaS/diaas/constants"
	"net/http"
)

func NewServer() *Server {
	srv := &Server{
		ServeMux: *http.NewServeMux(),
		Registry: NewRegistry(),
	}
	srv.HandleFunc("/login", Login)
	srv.HandleFunc("/register", Register)
	srv.HandleFunc(TriggerPrefix, srv.trigger)
	srv.HandleFunc(BindPrefix, srv.bind)
	srv.HandleFunc("/Functions/", srv.myfunctions)
	srv.HandleFunc("/describe/", srv.describe)
	srv.HandleFunc("/logs/", srv.logs)
	srv.HandleFunc("/deploy", srv.deploy)
	srv.HandleFunc("/listAlias", srv.listAlias)
	srv.HandleFunc("/list", srv.list)
	return srv
}
