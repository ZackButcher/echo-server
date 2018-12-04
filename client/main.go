package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/spf13/cobra"

	"github.com/ZackButcher/echo-server/api"
)

func main() {
	var (
		address string
	)

	root := &cobra.Command{
		Use:   "server",
		Short: "Starts Mixer as a server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatal(err)
			}
			client := api.NewServiceClient(conn)

			resp, err := client.Echo(context.Background(), &api.EchoRequest{
				Id:   "1",
				Body: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			})
			log.Printf("got response %v", resp)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	root.PersistentFlags().StringVar(&address, "address", "localhost:80", "Main port to serve on; always on /echo")

	if err := root.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}