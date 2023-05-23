// package config parses the configuration for the server and provides it as a struct
// The config is read either from a file provided by "config" flag or by the "CONFIG_JSON"
// environment variable. In both cases config is expected to be in JSON format.
package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Config includes all the configuration options for the server.
type Config struct {
	Host   string `json:"host"`
	Secret string `json:"secret"`
}

// ParseServerCfg reads the configuration either from "config" flag or from the "CONFIG_JSON" env variable
func ParseServerCfg() (Config, error) {
	var conf Config

	content, err := getConfigJson()
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal([]byte(content), &conf)
	if err != nil {
		return conf, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	return conf, nil
}

func getConfigJson() (string, error) {
	envConf := os.Getenv("CONFIG_JSON")
	if envConf != "" {
		return envConf, nil
	}

	path := flag.String("config", "config/default.json", "Path to the configuration file")
	flag.Parse()
	if path == nil || *path == "" {
		return "", errors.New("confilg path was not provided")
	}

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read configuration file %q", *path)
	}

	return string(content), nil
}
