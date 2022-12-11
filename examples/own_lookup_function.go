package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
)


type Config struct { // defining our config struct
	AppName string `my-source:"app_name"` // define value names in the source "my-source" using tags
	Host string `my-source:"host"`
	Port int `my-source:"port"`
}


func myLookUp(key string) (string, bool) { // our own simple lookup function
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
	config := Config{} // create Config instance
	result := lookupcfg.PopulateConfig(
		"my-source",
		myLookUp,
		&config,
	) // populate it using source "my-source" and our myLookUp function
	fmt.Println("Population result:", result) // print result of population (there were not errors so there is nothing interesting)
	fmt.Println("My config:", config) // print our populated config instance
}
