package model

type Provider struct {
    AccessKeyID     string `yaml:"access_key_id"`
    SecretAccessKey string `yaml:"secret_access_key"`
    Region          string `yaml:"region"`
}