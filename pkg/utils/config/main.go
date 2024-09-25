package config

import (
	"os"

	configModel "github.com/hamza-007/ret3iac/pkg/model/config"

	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(filename string) (*configModel.Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config configModel.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
