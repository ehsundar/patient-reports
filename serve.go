package main

import (
	"net"
	"net/http"

	"github.com/ehsundar/reports/internal/user"
	"github.com/gorilla/mux"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
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
