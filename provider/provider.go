package provider

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
	"github.com/sam701/file-publisher/files"
	"github.com/sam701/file-publisher/meta"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("serving file", id)
	me := meta.Get(id)

	ex, err := time.Parse(time.RFC3339, me.ExpirationTime)
	if err != nil {
		log.Println("ERROR while serving", id, err)
		w.WriteHeader(500)
		return
	}

	if ex.After(time.Now()) {
		if files.Exists(me.FileHash) {
			w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, me.FileName))
			http.ServeFile(w, r, files.FilePath(me.FileHash))
		} else {
			w.WriteHeader(500)
		}
	} else {
		log.Println("Expired")
		w.WriteHeader(404)
	}
}
