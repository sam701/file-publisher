package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sam701/file-publisher/provider"
	"github.com/sam701/file-publisher/receiver"
)

func startServer(addr string) {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.Path("/meta").Methods("POST").HandlerFunc(receiver.ReceiveMeta)
	api.Path("/files/{id}").Methods("POST").HandlerFunc(receiver.ReceiveFile)
	api.Path("/files/{id}/exists").Methods("GET").HandlerFunc(receiver.CheckFile)
	r.Path("/{id}").Methods("GET").HandlerFunc(provider.ServeFile)

	http.ListenAndServe(addr, r)
}
