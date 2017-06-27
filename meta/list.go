package meta

import (
	"io/ioutil"
	"os"

	"github.com/sam701/file-publisher/config"
)

func ListAll() (map[string]*PublishingRequest, error) {
	result := map[string]*PublishingRequest{}
	fis, err := ioutil.ReadDir(config.Current.MetaDir())
	if err != nil {
		return nil, err
	}

	for _, fi := range fis {
		name := fi.Name()
		id := name[:len(name)-5]

		pr := Get(id)
		result[id] = pr
	}

	return result, nil
}

func Remove(id string) error {
	return os.Remove(metaPath(id))
}
