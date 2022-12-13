package main

import (
	"fmt"
	"github.com/jieggii/lookupcfg"
	"os"
)

// try running it using
// PG_HOST=127.0.0.1 PG_PORT=5432 PG_USER=postgres PG_PASSWORD=password go run ./examples/env_lookup.go

// and just using
// go run ./examples/env_lookup.go
func main() {
	type Config struct { // defining our config struct
		PostgresHost string `env:"PG_HOST"`
		PostgresPort int    `env:"PG_PORT"`

		PostgresUser     string `env:"PG_USER"`
		PostgresPassword string `env:"PG_PASSWORD"`
	}

	config := Config{} // create Config instance
	result := lookupcfg.PopulateConfig(
		"env",
		os.LookupEnv,
		&config,
	) // populate it using source "env" and our os.LookupEnv function

	fmt.Printf("Population result: %+v\n", result) // print result of population.
	// there will be some useful information if
	// some mistakes were made in the environmental variables

	fmt.Printf("My config: %+v\n", config) // print our populated config instance
}
