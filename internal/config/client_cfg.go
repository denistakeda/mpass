package config

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type ClientCfg struct {
	Address string `json:"address"`
}

func ParseClientCfg(configPath string) (ClientCfg, error) {
	var res ClientCfg

	content, err := os.ReadFile(configPath)
	if err != nil {
		return res, errors.Wrapf(err, "failed to read configuration file %q", configPath)
	}

	err = json.Unmarshal(content, &res)
	if err != nil {
		return res, errors.Wrapf(err, "failed to unmarshal configuration file %s", content)
	}

	return res, nil
}
