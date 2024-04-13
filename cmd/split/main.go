package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Adelioz/split/internal/handlers"
	"github.com/Adelioz/split/pkg/logging"
)

var args struct {
	devel bool
}

func main() {
	flag.BoolVar(&args.devel, "devel", false, "enable dev mode")
	flag.Parse()

	logger := logging.NewLogger(args.devel)
	router := handlers.NewRouter(logger)
	server := http.Server{
		Addr:    ":65000",
		Handler: router,
	}
	logger.Sugar().Infof("Starting server %s", server.Addr)
	fmt.Println(server.ListenAndServe())
}
