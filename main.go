package main

import (
	"flag"
	"fmt"
	"os"

	"safebsc/config"
	dbconnect "safebsc/db"
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
	db := dbconnect.GetConnection()
	defer db.Close()
	q := "select now();"
	fmt.Println(q)
	fmt.Println(db.Conn().Exec(q))

	server.Init()
}
