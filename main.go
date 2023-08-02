package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	SGB "github.com/fbaube/sqlite-gobroem/gobroem"
)

const version = "0.1.0"

var options struct {
	db   string
	host string
	port uint
}

// printHeader print the welcome header.
func printHeader() {
	println("sqlite-gobroem v." + version)
}

// initConfig parse CLI config
func initConfig() {
	flag.StringVar(&options.host, "host", "localhost", "HTTP server hostname")
	flag.StringVar(&options.db, "db", "test/test.db", "SQLite DB file")
	flag.UintVar(&options.port, "port", 8000, "HTTP server listen port")
	flag.Parse()
}

// startServer initialize and start the web server.
func startServer() {
	api, err := SGB.NewAPI(options.db)
	if err != nil {
		log.Fatal("can not open DB:", err)
	}
	http.ListenAndServe(
		fmt.Sprintf("%s:%d", options.host, options.port),
		api.Handler("/", "/static/"),
	)
}

func main() {
	printHeader()
	initConfig()
	startServer()
}
