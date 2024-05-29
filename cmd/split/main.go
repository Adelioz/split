package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Adelioz/split/internal/handlers"
	"github.com/Adelioz/split/internal/repository/mongo"
	"github.com/Adelioz/split/internal/service"
	"github.com/Adelioz/split/pkg/logging"
)

var args struct {
	devel bool
}

func main() {
	flag.BoolVar(&args.devel, "devel", false, "enable dev mode")
	flag.Parse()

	logger := logging.NewLogger(args.devel)

	r, err := mongo.NewRepository("mongodb://test_user:p4ssw0rd@mongodb:27017")
	if err != nil {
		logger.Sugar().Fatalf("Failed to connect to database %s", err)
	}

	// mock repo
	// r, _ := mock.NewRepository()

	s := service.NewService(r)

	router := handlers.NewRouter(logger, s)
	server := http.Server{
		Addr:    ":65000",
		Handler: router,
	}
	logger.Sugar().Infof("Starting server %s", server.Addr)
	fmt.Println(server.ListenAndServe())
}
