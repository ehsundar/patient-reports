package main

import (
	"github.com/ehsundar/reports/internal/user"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	userController := user.NewController()
	userController.Register(r.PathPrefix("/user").Subrouter())

	if err = http.Serve(listener, r); err != nil {
		panic(err)
	}
}
