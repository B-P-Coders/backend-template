package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

func Load() error {
	tempConfig := Config{}
	if err := env.Parse(&tempConfig); err != nil {
		return err
	}
	Conf = &tempConfig
	fmt.Println(Conf)
	return nil
}
