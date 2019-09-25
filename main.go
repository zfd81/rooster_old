package main

import (
	"github.com/zfd81/rooster/conf"
	"github.com/zfd81/rooster/http"
)

const (
	fileName = "config.toml"
)

// Rooster runtime environment
var (
	Rre    map[string]interface{} = make(map[string]interface{})
	config                        = conf.NewConfig()
)

func main() {
	http.Start()
}
