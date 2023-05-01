package server

import (
	"log"
	"net/http"
	"time"

	"github.com/kuzminprog/service-provider-system/include/handler"
)

// StartServer() - starts the server with the specified address and port
func StartServer(addr string) {
	log.Println("Start server")
	server := &http.Server{
		Addr:         addr,
		Handler:      handler.InitRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
