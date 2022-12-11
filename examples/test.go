package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
)

type AppConfig struct {
	Host    string `env:"HOST"`
	Port    int    `env:"PORT"`
	AppName string `env:"APP_NAME"`
}

func envLookup(key string) (string, bool) {
	if key == "APP_NAME" {
		return "my app", true
	} else if key == "PORT" {
		return "bruh", true
	}
	return "localhost", true
}

func main() {
	cfg := AppConfig{}
	result := lookupcfg.PopulateConfig("env", envLookup, &cfg)
	fmt.Println("result:", result)
	fmt.Println("config: ", cfg)
}
