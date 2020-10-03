package module

import (
	"github.com/Netflix/go-env"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

type module struct {
	Name    string   `toml:"name"`
	Desc    string   `toml:"description"`
	Version string   `toml:"version"`
	License string   `toml:"license"`
	Authors []string `toml:"authors"`
}

type options struct {
	NoColor bool `toml:"no_color" env:"NO_COLOR"`
}

type Config struct {
	Module  module  `toml:"module"`
	Options options `toml:"options"`
}

var DefaultConfigPath = "mod.toml"

func ConfigLoad() (*Config, error) {
	cfg := Config{}

	if _, err := os.Stat(DefaultConfigPath); err == nil || os.IsExist(err) {
		buf, err := ioutil.ReadFile(DefaultConfigPath)
		if err != nil {
			return nil, err
		}
		err = toml.Unmarshal(buf, &cfg)
		if err != nil {
			return nil, err
		}
	}

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ConfigParse(source []byte) (*Config, error) {
	cfg := Config{}

	if source != nil {
		err := toml.Unmarshal(source, &cfg)
		if err != nil {
			return nil, err
		}
	}

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ConfigExists() bool {
	if _, err := os.Stat(DefaultConfigPath); err == nil || os.IsExist(err) {
		return true
	}
	return false
}
