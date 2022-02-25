package main

import (
	"flag"
	"fmt"
	"os"

	"safebsc/config"
	"safebsc/db"
	"safebsc/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.GetConnection()
	server.Init()
}
