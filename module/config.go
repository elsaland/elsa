package module

import (
	"github.com/BurntSushi/toml"
	"github.com/Netflix/go-env"
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

var configPath = "mod.toml"

func GetConfig() (*Config, error) {
	cfg := Config{}

	if _, err := os.Stat(configPath); err == nil || os.IsExist(err) {
		buf, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		source := string(buf)
		_, err = toml.Decode(source, &cfg)
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
