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

_Deferred to a later time_

## Add benchmark methods for capturing pipeline throughput times. Design the program so it can be run with and without goroutines.

Execution time was captured with a simple `start := time.Now()` / `time.Since(start)` technique.

For the program without goroutines, a function was built to apply the functions in `imageprocessing` in the correct order and to handle panics:

```
func processLinearly(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed!")
		}
	}()
	outPath := strings.Replace(path, "images/", "images/output/", 1)
	img := imageprocessing.ReadImage(path)
	img = imageprocessing.Resize(img)
	img = imageprocessing.Grayscale(img)
	imageprocessing.WriteImage(outPath, img)
	fmt.Println("Success!")
}
```

`processLinearly` was called at the end of `main.go` to process the images without use of goroutines:

```
// Without goroutines
start = time.Now()
for _, path := range imagePaths {
	processLinearly(path)
}
fmt.Println("Run Time Without Goroutines:", time.Since(start))
```

## Make additional code modifications as you see fit.

See above, where original image aspect ratio scaling was implemented.

## Build, test, and run the pipeline program with and without goroutines. Compare processing times with and without goroutines.

Execution through the terminal ran as follows:

```
PS C:\...\Data_Pipelines_Concurrency> .\goroutines_pipeline
Success!
Success!
Success!
Success!
Run Time With Goroutines: 124.0168ms
Success!
Success!
Success!
Success!
Run Time Without Goroutines: 169.1572ms
```

The pipeline is considerably faster with goroutines.  All images processed as expected.

## Prepare a complete README.md file documenting your work.

I hope you're enjoying it!

## References

aprln.  Reply to Sergio Tapia, "Golang - Getting the dimensions of an image. jpg, jpeg, png."  GitHub Gist, December 9, 2013, https://gist.github.com/sergiotapia/7882944.

Bodner, Jon.  _Learning Go, 2nd Edition._  O'Reilly Media Inc., 2024.

edap.  "Get image size with golang."  stackoverflow, February 12, 2014, https://stackoverflow.com/questions/21741431/get-image-size-with-golang.

InfectiouSoul.  "Trying to understand context better, specifically how ctx.Done() works."  Reddit, August 24, 2022, https://www.reddit.com/r/golang/comments/wwsclz/trying_to_understand_context_better_specifically/.

Menke, David.  _Red salmon or Sockeye salmon specimens_, digitized April 8, 2008, photograph, https://digitalmedia.fws.gov/digital/collection/natdiglib/id/2564.

Naseem, Ahsan.  "Recover and continue for loop if panic occur Golang."  stackoverflow, June 30, 2018, https://stackoverflow.com/questions/51113193/recover-and-continue-for-loop-if-panic-occur-golang.

Schlich, Jan.  resize, February 21, 2018, GitHub repository, https://github.com/nfnt/resize.

Singh, Amrit.  go_21_goroutines_pipeline, GitHub repository, February 7, 2024, https://github.com/code-heim/go_21_goroutines_pipeline.

Topinka, Lyn.  _Mount Rainier over Tacoma_, August 20, 1984, photograph, http://web.archive.org/web/20051103010250/http://vulcan.wr.usgs.gov/Imgs/Jpg/Rainier/Images/Rainier84_mount_rainier_and_tacoma_08-20-84.jpg via https://commons.wikimedia.org/wiki/File:Mount_Rainier_over_Tacoma.jpg.

Topinka, Lyn.  _St. Helens Plume from Harry's Ridge_, May 19, 1982, photograph, https://vulcan.wr.usgs.gov/Volcanoes/MSH/Images/MSH80/framework.html via https://commons.wikimedia.org/wiki/File:MSH82_st_helens_plume_from_harrys_ridge_05-19-82.jpg.

US Army Corps of Engineers, _Lake Washington Ship Canal, Hiram M. Chittenden Locks, 1995_, c. 1995, photograph, http://images.usace.army.mil/scripts/PortWeb.dll?query&field=Image%20name&opt=matches&value=1091-51.Jpg&template=Selected_Info&catalog=photoDVL via https://commons.wikimedia.org/wiki/File:Lake_Washington_ship_canal,_Hiram_M._Chittenden_Locks,_1995.jpg.
