package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
)

func lup(key string) (string, bool) { // our own simple lookup function
	switch key {
	case "app_name":
		return "My awesome application", true
	case "host":
		return "127.0.0.1", true
	case "port":
		return "8080", true
	case "some":
		return "hello", true
	case "N":
		return "1", true
	case "hewwo":
		return "0.5", true
	default:
		return "", false
	}
}

func main() {
	type Config struct { // defining our config struct
		AppName string  `my-source:"app_name"` // define value names in the source "my-source" using tags
		Host    string  `my-source:"host"`
		Port    int     `my-source:"port"`
		N       uint32  `my-source:"N"`
		Some    []byte  `my-source:"some"`
		MyFloat float32 `my-source:"hewwo"`
	}

	config := Config{} // create Config instance
	result := lookupcfg.PopulateConfig(
		"my-source",
		lup,
		&config,
	) // populate it using source "my-source" and our myLookUp function

	fmt.Printf(
		"Population result: %+v\n",
		result,
	) // print result of population (there were not errors so there is nothing interesting)
	fmt.Printf("My config: %+v\n", config) // print our populated config instance
	//fmt.Println(string(config.Some))
}
