package cleanup

import (
	"log"
	"time"

	"github.com/sam701/file-publisher/files"
	"github.com/sam701/file-publisher/meta"
)

func Run() {
	doCleanup()
	for range time.Tick(time.Hour) {
		doCleanup()
	}
}

func doCleanup() {
	log.Println("cleaning up")

	fileMap := map[string]bool{}
	ff, err := files.ListAll()
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	for _, f := range ff {
		fileMap[f] = false
	}

	metas, err := meta.ListAll()
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	for sharingId, pr := range metas {
		if !pr.Expired() {
			fileMap[pr.FileHash] = true
			continue
		}

		log.Println("Removing sharing", sharingId)
		err = meta.Remove(sharingId)
		if err != nil {
			log.Println(err)
			return
		}
	}

	for hash, used := range fileMap {
		if used {
			continue
		}

		log.Println("Removing file", hash)
		err = files.Remove(hash)
		if err != nil {
			log.Println(err)
		}
	}
}
