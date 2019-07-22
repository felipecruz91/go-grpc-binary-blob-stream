package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"../protos/chunker"
)

func main() {
	conn, err := grpc.Dial("localhost:8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cc := chunker.NewChunkerClient(conn)
	client, err := cc.Chunker(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var blob []byte
	start := time.Now()
	for {
		c, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(blob))
				elapsed := time.Since(start)
				log.Printf("Download took %s", elapsed)
				return
			}

			panic(err)
		}

		blob = append(blob, c.Chunk...)
	}
}
