package files

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/sam701/file-publisher/config"
)

func FilePath(id string) string {
	return path.Join(config.Current.FileDir(), id)
}

func Exists(id string) bool {
	_, err := os.Stat(FilePath(id))
	return err == nil
}

func Save(id string, r io.Reader) error {
	log.Println("Saving file", id)
	f, err := os.OpenFile(FilePath(id), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln("ERROR: cannot open file:", err)
	}
	defer f.Close()

	_, err = io.Copy(f, r)

	return err
}

func Remove(id string) error {
	return os.Remove(FilePath(id))
}

func ListAll() ([]string, error) {
	result := []string{}
	fis, err := ioutil.ReadDir(config.Current.FileDir())
	if err != nil {
		return nil, err
	}

	for _, fi := range fis {
		result = append(result, fi.Name())
	}

	return result, nil
}
