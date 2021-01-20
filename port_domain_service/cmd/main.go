package main

import (
	"fmt"
	pb "github.com/aelmel/ports-infra/port_domain_service/internal/proto"
	"github.com/aelmel/ports-infra/port_domain_service/internal/repository/mongo"
	"github.com/aelmel/ports-infra/port_domain_service/internal/service/port"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "", flag.ExitOnError)
	var (
		grpcAddr        = fs.String("GRPC_ADDR", ":50051", "GRPC port")
		mongoPort       = fs.String("MONGO_PORT", "27017", "Mongo port")
		mongoHost       = fs.String("MONGO_HOST", "localhost", "Mongo host")
		mongoUser       = fs.String("MONGO_USERNAME", "", "Mongo username")
		mongoPwd        = fs.String("MONGO_PWD", "", "Mongo password")
		mongoDb         = fs.String("MONGO_DATABASE", "", "Mongo database")
		mongoCollection = fs.String("MONGO_COLLECTION", "", "Mongo collection")
	)
	_ = fs.Parse(os.Args[1:])
	repo, err := mongo.NewRepo(mongo.Configuration{
		Port:       *mongoPort,
		Host:       *mongoHost,
		Username:   *mongoUser,
		Password:   *mongoPwd,
		Database:   *mongoDb,
		Collection: *mongoCollection,
	})

	if err != nil {
		log.Fatalf("Issue connecting to repo %v", err)
		os.Exit(1)
	}

	errChan := make(chan error)
	go func() {
		listener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := port.NewGrpcPortSvc(repo)
		grpcServer := grpc.NewServer()
		pb.RegisterPortServiceServer(grpcServer, handler)
		errChan <- grpcServer.Serve(listener)
	}()
	log.Info(fmt.Sprintf("Starting grpc server on port %s", *grpcAddr))
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		_ = <-gracefulStop
		repo.Close()
		// handle it
		os.Exit(0)
	}()
	log.Fatalf("exit %v", <-errChan)
}
