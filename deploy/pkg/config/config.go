package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config is application configuration
type Config struct {
	AwsAccessKeyID    string `env:"AWS_ACCESS_KEY_ID"`
	AwsRegion         string `env:"AWS_REGION"`
	CdkDefaultAccount string `env:"CDK_DEFAULT_ACCOUNT"`
	CdkDefaultRegion  string `env:"CDK_DEFAULT_REGION"`

	Vpc struct {
		Name      string `env:"VPC_NAME"`
		CidrBlock string `env:"CIDR_BLOCK"`
	}
	Repository struct {
		Name string `env:"REPOSITORY_NAME"`
	}
	Cluster struct {
		Name string `env:"CLUSTER_NAME"`
	}
}

// Load returns configuration made from environment variables
func Load() *Config {
	_ = godotenv.Load()
	c := Config{}
	_ = env.Parse(&c)

	return &c
}