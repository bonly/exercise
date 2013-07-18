package main 

import "code.google.com/p/go-opencv/trunk/opencv" 

func main() { 
        img := new(opencv.Image); 
        win := new(opencv.Window); 

        img.Load("lena.jpg"); 
        defer img.Release(); 

        win.Create("winName"); 
        defer win.Release(); 

        win.ShowImage(img); 
        opencv.WaitKey(0); 
} 
