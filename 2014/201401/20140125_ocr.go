package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	// This is the simlest way :)
	// out := gosseract.Must(gosseract.Params{Src: "/tmp/dow.jpg",Languages:"chi_sim"});
	// fmt.Println(out);

	image := opencv.LoadImage("/tmp/dow.jpg");
	if image == nil {
		panic("LoadImage fail");
	}
	defer image.Release();

	w := image.Width();
	h := image.Height();

	// Create the output image
	// cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3);
	// defer cedge.Release();

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	// edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1);
	defer gray.Release();
	// defer edge.Release();

	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY);

	opencv.SaveImage("/tmp/dow1.jpg", gray, 0);

	// Using client digits
	client, _ := gosseract.NewClient();
	client = client.Digest("/tmp/dig");
	out, _ := client.Src("/tmp/dow1.jpg").Out();
	fmt.Println(out);
}
