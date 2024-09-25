package model

import (
	providerModel "github.com/hamza-007/ret3iac/pkg/model/provider"
	resourceModel "github.com/hamza-007/ret3iac/pkg/model/resource"
)

type Config struct {
	Provider providerModel.Provider    `yaml:"providers"`
	Resources []*resourceModel.Resource `yaml:"resources"`
}
