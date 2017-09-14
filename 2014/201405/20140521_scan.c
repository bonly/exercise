#include <cv.h>
#include <highgui.h>
#include <opencv2/photo/photo_c.h>
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>
#include <locale.h>
#include <stdlib.h>  

#define min(x,y) ((x) > (y) ? (y) :(x))

IplImage *img0 = 0, *img = 0;
IplImage *gray = 0, *gray2 = 0;
IplImage *tmp = 0, *tmp2 = 0;

CvPoint prev_pt = {-1, -1};
CvPoint pt;
int level = 128;
int level2 = 50;
int thr1 = 100;
int thr2 = 200;
int count = 1;

void binarization(IplImage* image, int mode){
  cvCvtColor(img, image, CV_BGR2GRAY);
  switch(mode){
	  case 1:{
	    cvThreshold(image, image, level, 255, CV_THRESH_BINARY);
	    cvShowImage("灰度", image);
	    break;
	  }
	  case 2:{
	    cvThreshold(image, image, level2, 255, CV_THRESH_BINARY);
	    cvShowImage("二值化", image);
	    break;
	  }
  }
}

void on_change(int pos){ //灰度值参数变化响应
  binarization(gray2, 1);  
}


void binarization_canny(IplImage* image, int t1, int t2){
  cvCvtColor(img, image, CV_BGR2GRAY);
  cvCanny(image, image, t1, t2, 3 );
  cvShowImage("二值化", image);
}

void on_change2(int pos){//2值化参数变化响应
  binarization_canny(gray, thr1, thr2);
}

int Scan(){
// int Scan(int argc, char *argv[]){
	// setlocal(LC_CTYPE, "zh_CN.UTF-8");
	// if (argc < 2){
	// 	fprintf(stderr, "%s [file.jpg]\n", argv[0]);
	// 	return 1;
	// }

	// char *filename = argv[1];
	char *filename = "id.jpg";  //取得文件名

	if((img0 = cvLoadImage(filename, -1)) == 0){ //加载图片
    	fprintf(stderr, "读取文件[%s]失败\n", filename);
    	return 0;
	}

	//命名窗口
	cvNamedWindow("二值化", 1);
	cvMoveWindow("二值化", 640, 0 );
	cvNamedWindow("灰度", 1);
	cvMoveWindow("灰度", 0, 480);
  	cvNamedWindow("证件", 1);	
	cvMoveWindow("证件", 0, 0); //设置窗口位置
	
	
	
	//
	img = cvCloneImage(img0);
	tmp = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, 3);
  
  	//2值化
  	gray = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, 1);
  	gray2 = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, 1);

  	//显示原始图
  	cvShowImage("证件", img); //创建窗口显示
  	// cvShowImage("gray", gray); //每个cvShowImage都会创建一个窗口
  	cvCreateTrackbar("灰度设置", "灰度", &level, 255, on_change); //绑定窗口
	binarization(gray2, 1); //创建后初始化调用

	//二值化设置调节
	cvCreateTrackbar( "二值化设置", "二值化", &thr1, 255, on_change2);
	cvCreateTrackbar( "二值化设置", "二值化", &thr2, 255, on_change2);
	binarization_canny(gray, thr1, thr2);

	//保证初始化后显示
	// cvShowImage("二值化", gray);
	// cvShowImage("灰度", gray2);	

  	for (;;){
  		int c = cvWaitKey(0);
  		if (c == 27) break; //Esc退出
  	}

  	cvSaveImage("test.jpg", gray, 0);

	cvReleaseImage(&img);
	cvReleaseImage(&img0);
	cvReleaseImage(&tmp);
	cvReleaseImage(&tmp2);
	cvReleaseImage(&gray);
	cvReleaseImage(&gray2);  	

  	return 0;
}

int Main(){
	return Scan();
}
