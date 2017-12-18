package http

import (
	"runtime"
	"path"
	"github.com/BurntSushi/toml"
	"errors"
)

type Config struct {
	Listen string
}

func GetConfig() Config {
	var c Config
	_, filename, _, _ := runtime.Caller(1)
	cfg := path.Join(path.Dir(filename), "/../config/http.toml")
	if _, err := toml.DecodeFile(cfg, &c); err != nil {
		panic(errors.New("parse http.toml fail: " + err.Error()))
	}
	if c.Listen == "" {
		panic(errors.New("invalid file http.toml"))
	}
	return c
}
