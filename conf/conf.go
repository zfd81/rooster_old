package conf

import (
	"github.com/BurntSushi/toml"
	"os/user"
)

type Config struct {
	Path string `toml:"path"`
}

func (c *Config) Load(confFile string) error {
	_, err := toml.DecodeFile(confFile, c)
	return err
}

var defaultConf = Config{
	Path: "/tmp/rooster/meta",
}

var globalConf = defaultConf

func init() {
	user, _ := user.Current()
	defaultConf.Path = user.HomeDir + "/rooster/meta"
}

func NewConfig() *Config {
	conf := defaultConf
	return &conf
}

func GetGlobalConfig() *Config {
	return &globalConf
}
