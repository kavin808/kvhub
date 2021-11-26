package tool

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	// Listen port
	Listen string `yaml:"listen"`
	// LogDir log file folder
	LogDir string `yaml:"logDir"`
}

// NewConfig load config from path
func NewConfig(path string) (*Config, error) {
	c := &Config{}
	if err := c.load(path); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}
