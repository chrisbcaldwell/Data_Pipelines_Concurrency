package main

import (
	"context"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(ctx context.Context, paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		defer close(out)
		for _, p := range paths {
			select {
			case <-ctx.Done():
				return
			default:
				job := Job{InputPath: p,
					OutPath: strings.Replace(p, "images/", "images/output/", 1)}
				job.Image = imageprocessing.ReadImage(p)
				out <- job
			}
		}
	}()
	return out
}

func resize(ctx context.Context, input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		defer close(out)
		for job := range input { // Read from the channel
			select {
			case <-ctx.Done():
				return
			default:
				job.Image = imageprocessing.Resize(job.Image)
				out <- job
			}
		}

	}()
	return out
}

func convertToGrayscale(ctx context.Context, input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for job := range input { // Read from the channel
			select {
			case <-ctx.Done():
				return
			default:
				job.Image = imageprocessing.Grayscale(job.Image)
				out <- job
			}
		}

	}()
	return out
}

func saveImage(ctx context.Context, input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		for job := range input { // Read from the channel
			select {
			case <-ctx.Done():
				return
			default:
				imageprocessing.WriteImage(job.OutPath, job.Image)
				out <- true
			}
		}

	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	channel1 := loadImage(ctx, imagePaths)
	channel2 := resize(ctx, channel1)
	channel3 := convertToGrayscale(ctx, channel2)
	writeResults := saveImage(ctx, channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
