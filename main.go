package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"math/rand"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"github.com/spf13/cobra"
	"github.com/dustinkirkland/golang-petname"

	"github.com/ZackButcher/echo-server/api"
)

func main() {
	var (
		servingPort uint16
		id string
	)


	root := &cobra.Command{
		Use:   "server",
		Short: "Starts Mixer as a server",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				rand.Seed(time.Now().UTC().UnixNano())
				id = petname.Generate(2, "-")
				log.Printf("no ID provided at startup, picking a random one")
			}
			log.Printf("starting with ID: %v", id)

			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 80))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			grpcServer := grpc.NewServer()
			api.RegisterServiceServer(grpcServer, &server{id})
			grpcServer.Serve(lis)
		},
	}

	root.PersistentFlags().Uint16VarP(&servingPort, "server-port", "s", 80, "Main port to serve on; always on /echo")
	root.PersistentFlags().StringVar(&id, "id", "", "Name that identifies this instance (the ID is returned as part of the response)")

	if err := root.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}

type server struct{
	id string
}

// Echo replies to a request by sending back exactly the requests data, with the server's time
func (s server) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	now, err := ptypes.TimestampProto(time.Now())
	log.Printf("handling request at %s", now.String())
	return &api.EchoResponse{
		Id:         s.id,
		Body:       req.Body,
		ServerTime: now,
	}, err
}