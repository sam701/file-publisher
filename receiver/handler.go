package receiver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sam701/file-publisher/config"
	"github.com/sam701/file-publisher/files"
	"github.com/sam701/file-publisher/meta"
	"github.com/sam701/file-publisher/util"
)

func ReceiveMeta(w http.ResponseWriter, r *http.Request) {
	log.Println("receiving meta")
	var pr meta.PublishingRequest
	json.NewDecoder(r.Body).Decode(&pr)

	shareId := util.GenerateNewId(20)
	resp := &meta.PublishingResponse{
		SharingURL: fmt.Sprintf("%s/%s", config.Current.BaseURL, shareId),
	}

	err := meta.Save(shareId, &pr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := files.Save(id, r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.WriteHeader(204)
	}
}

func CheckFile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if files.Exists(id) {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}
