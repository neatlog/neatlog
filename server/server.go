package server

import (
	"fmt"
	"net/http"
	"time"
)

var (
	HTTPServer = &http.Server{
		Handler:     http.DefaultServeMux,
		ReadTimeout: 15 * time.Second,
	}
)

func Start() {
	HTTPServer.Addr = ":8080" // TODO make configurable

	http.HandleFunc("/health", healthEndpoint)
}

func ListenAndServer() error {
	return HTTPServer.ListenAndServe()
}

func healthEndpoint(w http.ResponseWriter, _ *http.Request) {
	var oerr error
	defer func() {
		if err := recover(); err != nil {
			oerr = fmt.Errorf("recovering from: %v\n", err)
		}
		if oerr != nil {
			fmt.Println(oerr) // TODO handle better logging
		}
	}()

	_, oerr = w.Write([]byte("online"))
	if oerr != nil {
		return
	}
}
