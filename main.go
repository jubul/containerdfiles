package main

import (
	"context"
	"log"
	
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd"
)

func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}

func redisExample () error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
	
	ctx := namespaces.WithNamespace(context.Background(), "example")
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	if err != nil {
		return err
}
log.Printf("Successfully pulles %s image\n", image.Name())

return nil

}





