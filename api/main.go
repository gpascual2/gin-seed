package main

import (
	"flag"
	"fmt"
	"os"
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
}
