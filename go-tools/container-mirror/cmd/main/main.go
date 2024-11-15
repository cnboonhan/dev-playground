package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/types"
)

func downloadImage(ctx context.Context, ref types.ImageReference) error {
	img, err := ref.NewImage(ctx, nil)
	if err != nil {
		return fmt.Errorf("creating image: %w", err) // Wrap errors for better context
	}
	defer img.Close() // Crucial for resource cleanup
	// ... (rest of download logic)
	return nil
}

func main() {
	imageListPathPtr := flag.String("imagelist-path", "sample-imagelist", "Text file of all container images on each line.")
	log.Println("imagelist-path: ", *imageListPathPtr)

	concurrencyPtr := flag.Int("concurrency", 1, "Number of simultaneous downloads.")
	log.Println("concurrency:", *concurrencyPtr)

	imagelist_file, err := os.Open(*imageListPathPtr)
	if err != nil {
		panic(err)
	}
	defer imagelist_file.Close()

	log.Println("Validating ImageList contents...")
	imagelist_scanner := bufio.NewScanner(imagelist_file)
	for imagelist_scanner.Scan() {
		image_ref := imagelist_scanner.Text()
		log.Println(image_ref)
		_, err := docker.ParseReference("//" + image_ref)
		if err != nil {
			panic(err)
		}
	}

	imagelist_file.Seek(0, 0)
	imagelist_scanner = bufio.NewScanner(imagelist_file)
	for imagelist_scanner.Scan() {
		image_ref := imagelist_scanner.Text()
		log.Println("Downloading image: " + image_ref)
		log.Println(image_ref)
		ref, err := docker.ParseReference("//" + image_ref)
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

}
