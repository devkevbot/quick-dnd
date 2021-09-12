package main

import (
	"flag"

	_ "github.com/lib/pq"
)

func main() {
	configFile := flag.String("cfg", "config.yml", "Configuration YAML File")
	flag.Parse()
	startServer(configFile)
}
