package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var config Config

func main() {
	// parse args and config file
	var configPath string
	args := os.Args
	processArgs(args, &configPath, &config)

	// start http server
	http.Handle("/sync", RequestMiddleware(Sync))
	http.Handle("/update_whitelist", RequestMiddleware(UpdateWhitelist))

	addr := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("GAU server listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

