# MSDS 431 Week 6 Assignment: Data Pipelines with Concurrency

Chris Caldwell<br>
Northwestern University<br>
MSDS 431<br>
May 11, 2025

## Assignment Requirements

Amrit Singh ([CODEHEIM](https://www.codeheim.io/)) offers an example of a Go image processing pipeline with concurrency. He provides a [GitHub](https://github.com/code-heim/go_21_goroutines_pipeline) code repository and [video tutorial](https://www.youtube.com/watch?v=8Rn8yOQH62k). Let's replicate his work using image files that we select.

* Clone the GitHub repository for image processing. 
* Build and run the program in its original form.
* Add error checking for image file input and output.
* Replace the four input image files with files of your choosing.
* Add unit tests to the code repository.
* Add benchmark methods for capturing pipeline throughput times. Design the program so it can be run with and without goroutines. 
* Make additional code modifications as you see fit.
* Build, test, and run the pipeline program with and without goroutines. Compare processing times with and without goroutines.
* Prepare a complete README.md file documenting your work.

(Optional) Note that resizing images can cause distortion. How can we preserve the aspect ratio of images? Suppose we detect the size of the input images in pixels and ensure that the output image has the same shape, rather than the 500x500 shape in the image_processing.Resize helper function.

## Clone the GitHub repository for image processing.  Build and run the program in its original form.

### Clone the repository

* Create a new empty Github repository, and a new local folder.
* Visit [the original project's Github repository](https://github.com/code-heim/go_21_goroutines_pipeline) and fork the repository into the new respository that was just created.
* In the terminal, navigate to the new local project folder, then run `git clone (your repository URL).git`

### Run the program

The original program converts image files to 500 pixel by 500 pixel grayscale copies.

From the terminal, run:
* Create the executable: run `go build`
* Run `./goroutines_pipeline.exe`

The input images were located in the /images folder:

![input image 1](/images/image1.jpeg)
![input image 2](/images/image2.jpeg)
![input image 3](/images/image3.jpeg)
![input image 4](/images/image4.jpeg)

The output images were saved in the /images/output folder:

![output image 1](/images/output/image1.jpeg)
![output image 2](/images/output/image2.jpeg)
![output image 3](/images/output/image3.jpeg)
![output image 4](/images/output/image4.jpeg)

## Add error checking for image file input and output.

One opportunity to avoid problems in the original code is handling goroutines that terminate early.  Each function in `main.go` calls a for loop that initiates a goroutine.  As an example, the original version of `loadImage` was:

```
func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}
```

After importing the `context` library the function was changed to detect and close terminated channels:

```
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
```

The remaing `main.go` functions `resize`, `covertToGrayscale`, and `saveImage` were similarly reconfigured.

## Replace the four input image files with files of your choosing.

Four public domain images of Pacific Northwest interest were added:

![my image 1, Ballard Locks](/images/ballard_locks.jpg)
![my image 2, Tahoma over Tacoma](/images/Mount_Rainier_over_Tacoma.jpg)
![my image 3, sockeye salmon in spawning colors](/images/sockeye.jpg)
![my image 4, Mount Saint Helens with steam plume, 1982](/images/st_helens_1982.jpg)

These images were selected in part because they have different width and height dimensions to test the adjustment to retain aspect ratio.  That adjustment was made to the `imageprocessing.Resize` function as follows:

```
func Resize(img image.Image) image.Image {
	maxSize := uint(500)
	b := img.Bounds()
	width := b.Dx()
	height := b.Dy()
	newWidth := uint(maxSize)
	newHeight := uint(maxSize)
	// resize.Resize will scale the image if the smaller dimension is passed as 0
	if width < height {
		newWidth = 0
	}
	if height < width {
		newHeight = 0
	}
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}
```

The imported function resize.Resize allows automatic scaling to the original aspect ratio if either the height or width is passed as 0.

Resuling images:

![output image 1, Ballard Locks](/images/output/ballard_locks.jpg)
![output image 2, Tahoma over Tacoma](/images/output/Mount_Rainier_over_Tacoma.jpg)
![output image 3, sockeye salmon in spawning colors](/images/output/sockeye.jpg)
![output image 4, Mount Saint Helens with steam plume, 1982](/images/output/st_helens_1982.jpg)

## Add unit tests to the code repository.

## Add benchmark methods for capturing pipeline throughput times. Design the program so it can be run with and without goroutines. 

## Make additional code modifications as you see fit.

## Build, test, and run the pipeline program with and without goroutines. Compare processing times with and without goroutines.

## Prepare a complete README.md file documenting your work.

## References

https://www.reddit.com/r/golang/comments/wwsclz/trying_to_understand_context_better_specifically/

ch 12 learning go 2nd ed

Wikimedia commons

each photo

https://stackoverflow.com/questions/21741431/get-image-size-with-golang

https://github.com/nfnt/resize

https://gist.github.com/sergiotapia/7882944 user https://gist.github.com/aprln reply
