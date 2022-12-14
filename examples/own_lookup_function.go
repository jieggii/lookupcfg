package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
)

// our own simple lookup function
func myLookUp(key string) (string, bool) {
	switch key {
	case "app_name":
		return "My awesome application", true
	case "host":
		return "127.0.0.1", true
	case "port":
		return "8080", true
	default:
		return "", false
	}
}

func main() {
	// defining our config struct
	type Config struct {
		AppName string `my-source:"app_name"` // define value names in the source "my-source" using tags
		Host    string `my-source:"host"`
		Port    int    `my-source:"port"`
	}

	// create Config instance
	config := Config{}

	// populate it using source "my-source" and our myLookUp function
	result := lookupcfg.PopulateConfig("my-source", myLookUp, &config)

	// print result of population (there were not errors so there is nothing interesting)
	fmt.Printf("Population result: %+v\n", result)
	// >>> Population result: &{MissingFields:[] IncorrectTypeFields:[]}

	// print our populated config instance
	fmt.Printf("My config: %+v\n", config)
	// >>> My config: {AppName:My awesome application Host:127.0.0.1 Port:8080}
}
