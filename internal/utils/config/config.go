package config

import (

)

type Config struct {
	Env      string        `yaml:"env" env-required:"true"`
}
