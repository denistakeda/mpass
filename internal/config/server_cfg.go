package server_cfg

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Config struct {
	Host string `json:"host"`
}

func ParseCfg() (Config, error) {
	var conf Config

	path := flag.String("config", "config/default.json", "Path to the configuration file")
	if path == nil || *path == "" {
		return conf, errors.New("confilg path was not provided")
	}

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		return conf, errors.Wrapf(err, "failed to read configuration file %q", *path)
	}

	err = json.Unmarshal(content, &conf)
	if err != nil {
		return conf, errors.Wrapf(err, "failed to unmarshal the configuration file %q", *path)
	}

	return conf, nil
}
