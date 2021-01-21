package main

import (
	"fmt"
	"github.com/aelmel/ports-infra/client_api/internal/client/port"
	"github.com/aelmel/ports-infra/client_api/internal/parser/json"
	"github.com/aelmel/ports-infra/client_api/internal/transport"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "", flag.ExitOnError)
	var (
		grpcHost = fs.String("PORT_GRPC_HOST", "localhost", "GRPC host")
		grpcAddr = fs.String("PORT_GRPC_ADDR", "50051", "GRPC port")
		jsonPath = fs.String("FILE_PATH", "/tmp", "path to monitor for json")
		httpAddr = fs.String("HTTP_PORT", ":8080", "http port")
	)
	_ = fs.Parse(os.Args[1:])
	client, err := port.NewGrpcClient(*grpcHost, *grpcAddr)
	if err != nil {

	}
	parser := json.NewParser(client, *jsonPath)
	go parser.Monitor()

	var h http.Handler
	{
		h = transport.NewHTTPHandler(client)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Info("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	log.Info("err", <-errs)
}
