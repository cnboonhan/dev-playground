package main

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/docker"
)

func main() {
	ref, err := docker.ParseReference("//fedora")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	img, err := ref.NewImage(ctx, nil)
	if err != nil {
		panic(err)
	}

	defer img.Close()
	b, _, err := img.Manifest(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", string(b))
}
