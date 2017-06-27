package config

import (
	"os"
	"path"
)

type Config struct {
	dataDir string
	BaseURL string
}

func (s *Config) FileDir() string {
	return path.Join(s.dataDir, "files")
}

func (s *Config) MetaDir() string {
	return path.Join(s.dataDir, "meta")
}

var Current *Config

func Read(dataDir, baseURL string) *Config {
	Current = &Config{dataDir, baseURL}
	os.MkdirAll(Current.FileDir(), 0700)
	os.MkdirAll(Current.MetaDir(), 0700)
	return Current
}
