package main

import (
	"ByLCY/ShortenUrl/router"
	"ByLCY/ShortenUrl/variable"
	"log"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", variable.HttpAdderss)

	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	r := router.NewRouter()
	defer listener.Close()
	server := &http.Server{
		Addr:    variable.HttpAdderss,
		Handler: r.Handler(),
	}

	if variable.Crt == "" || variable.Key == "" {
		server.ListenAndServe()
		return
	}
	server.ListenAndServeTLS(variable.Crt, variable.Key)
}
