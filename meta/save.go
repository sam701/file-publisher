package meta

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/sam701/file-publisher/config"
)

func metaPath(id string) string {
	return path.Join(config.Current.MetaDir(), id+".json")
}

func Save(id string, meta *PublishingRequest) error {
	f, err := os.OpenFile(metaPath(id), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln("ERROR: cannot open file:", err)
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(meta)
}

func Get(id string) *PublishingRequest {
	f, err := os.Open(metaPath(id))
	if err != nil {
		log.Println("WARN: no sharing with id:", id, err)
		return nil
	}
	defer f.Close()

	var pr PublishingRequest
	err = json.NewDecoder(f).Decode(&pr)
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	return &pr
}
