// try running this program using
//     PG_HOST=127.0.0.1 PG_PORT=5432 PG_USER=postgres PG_PASSWORD=password go run ./examples/env_lookup.go
// and just using
//     go run ./examples/env_lookup.go

package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
	"os"
)

func main() {
	// defining our config struct
	type Config struct {
		PostgresHost string `env:"PG_HOST"`
		PostgresPort int    `env:"PG_PORT"`

		PostgresUser     string `env:"PG_USER"`
		PostgresPassword string `env:"PG_PASSWORD"`
	}

	// create Config instance
	config := Config{}

	// populate it using source "env" and our os.LookupEnv function
	result := lookupcfg.PopulateConfig("env", os.LookupEnv, &config)

	// print result of population. There will be some useful information if
	// any mistakes were made in the environmental variables
	fmt.Printf("Population result: %+v\n", result)

	// print our populated config instance
	fmt.Printf("My config: %+v\n", config)
}
