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

## Replace the four input image files with files of your choosing.

## Add unit tests to the code repository.

## Add benchmark methods for capturing pipeline throughput times. Design the program so it can be run with and without goroutines. 

## Make additional code modifications as you see fit.

## Build, test, and run the pipeline program with and without goroutines. Compare processing times with and without goroutines.

## Prepare a complete README.md file documenting your work.
