package main

import (
	"flag"
	"fmt"
	"os"

	"bitbucket.org/gpascual2/gin-seed/api/config"
)

func main() {
	var env string
	flag.StringVar(&env, "e", "dev", "Working environment. Default: dev")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Environment: ", env)
	config.Init(env)

	c := config.GetConfig()
	fmt.Printf("HTTP Port: %v\n", c.Get("server.port"))
}
