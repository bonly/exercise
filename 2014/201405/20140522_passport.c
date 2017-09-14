#ifndef __PASSPORT_C__
#define __PASSPORT_C__

#include <cv.h>
#include <highgui.h>

#include <stdio.h>

#include "20140522_passport.h"

IplImage *img = 0;
IplImage *cut = 0;
IplImage *gray = 0;
IplImage *blur = 0;
IplImage *grad = 0;
IplImage *sobel = 0;
IplImage *mclose = 0;
IplImage *bin = 0;
IplImage *erode = 0;
IplImage *search = 0;

CvMemStorage * storage = 0;  
CvSeq * contour = 0;  

int Main(int argc, char** argv){
	printf("%s\n", __TIME__);

    char *args[argc]; //转换输入参数数组
    for (int idx=0; args[idx] = *argv++; idx++);
	
	//加载原始文件
    char *filename = args[1];
	if((img = cvLoadImage(filename, -1)) == 0){ //加载图片
    	fprintf(stderr, "读取文件[%s]失败\n", filename);
    	return 0;
	}

	Origin_img(); //原图
	// Resize(); //裁剪
	Rgb2Gray();//灰度
	GaussianBlur();//模糊
	MorphologyEx_gradient();//梯度
	Sobel();//边缘
	MorphologyEx_close();//关闭
	Erode();//侵蚀
	Threshold_Bin();//二化
	Find();//查找
	

	process_key(); //处理输入

	//释放资源
	cvReleaseImage(&img);
	cvReleaseImage(&cut);
	cvReleaseImage(&gray);
	cvReleaseImage(&grad);
	cvReleaseImage(&sobel);
	cvReleaseImage(&mclose);
	cvReleaseImage(&bin);
	cvReleaseImage(&erode);
	cvReleaseImage(&search);

	printf("%s\n", __TIME__);
	return 0;
}


void Origin_img(){
	cvNamedWindow("原图", 1);
	cvMoveWindow("原图", 0, 0);

	cvShowImage("原图", img);
}

void Resize(){
	cvNamedWindow("裁剪", 1);
	cvMoveWindow("裁剪", 480, 0);

	CvSize size;
	size.width = img->height * 2;
	size.height = img->height * 2;

	cut = cvCreateImage(size, img->depth, img->nChannels);
	cvResize(img, cut, CV_INTER_LINEAR );
	cvShowImage("裁剪", cut);
}

void Rgb2Gray(){
	cvNamedWindow("灰化", 1);
	cvMoveWindow("灰化", 0, 200);

	gray = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, 1); //先创建色深空间, 通道channel=1
	cvCvtColor(img, gray, CV_BGR2GRAY);
	cvShowImage("灰化", gray);	
}

void GaussianBlur(){
	blur = cvCreateImage(cvGetSize(gray), 8, 1);  
	cvSmooth(gray, blur, CV_GAUSSIAN, 7, gray->nChannels, 0, 0);//模糊
}

void MorphologyEx_gradient(){ 
	cvNamedWindow("梯度", 1);
	cvMoveWindow("梯度", 0, 300);


	//创建一个形态区
    IplConvKernel* morphKernel = cvCreateStructuringElementEx(
    		3, 3,  //3x3大小的画笔
    		0, 0, 
    		CV_SHAPE_ELLIPSE, 0); //椭圆形
    IplImage * temp = cvCreateImage(cvGetSize(gray), 8, 1);   //需与原图色深，通道一致
    grad = cvCreateImage(cvGetSize(gray), 8, 1);     
    cvMorphologyEx(blur, grad, temp, morphKernel, CV_MOP_GRADIENT, 6);	//形态梯度 最后值：次数
    // cvMorphologyEx(blur, grad, temp, NULL, CV_MOP_GRADIENT, 6);	//形态梯度 NULL:默认就是3x3的
    cvReleaseImage(&temp);

	cvShowImage("梯度", grad);	
}

void Sobel(){//边缘检测 
	cvNamedWindow("边缘", 1);
	cvMoveWindow("边缘", 0, 400);	
	sobel = cvCreateImage(cvGetSize(grad), IPL_DEPTH_8U, 1); 
	cvSobel(grad, sobel, 1, 1, 3); 

	cvShowImage("边缘", sobel);
}

void MorphologyEx_close(){//形态关闭 边缘检测没作用，不用的好
	cvNamedWindow("关闭", 1);
	cvMoveWindow("关闭", 0, 500);

	//创建一个形态区
    IplConvKernel* morphKernel = cvCreateStructuringElementEx(
    		3, 3,  //3x3大小的画笔
    		0, 0, 
    		CV_SHAPE_RECT, 0); //矩形
    IplImage * temp = cvCreateImage(cvGetSize(grad), 8, 1);   //需与原图色深，通道一致
    mclose = cvCreateImage(cvGetSize(grad), 8, 1);
    // cvMorphologyEx(sobel, mclose, temp, morphKernel, CV_MOP_CLOSE, 3);	//形态梯度 
    cvMorphologyEx(grad, mclose, temp, NULL, CV_MOP_CLOSE, 4);	//形态梯度 

	cvShowImage("关闭", mclose);	
}

void Threshold_Bin(){ //二值化
	cvNamedWindow("二值", 1);
	cvMoveWindow("二值", 0, 600);

	bin = cvCreateImage(cvGetSize(erode), 8, 1);
	cvThreshold(mclose, bin, 0, 255, CV_THRESH_BINARY|CV_THRESH_OTSU);
	cvShowImage("二值", bin);	
}

void Erode(){ //侵蚀
	cvNamedWindow("侵蚀", 1);
	cvMoveWindow("侵蚀", 0, 700);

	erode = cvCreateImage(cvGetSize(mclose), 8, 1);
	cvErode(mclose, erode, NULL, 4);
	cvShowImage("侵蚀", erode);
}

void Find(){
	cvNamedWindow("搜索", 1);
	cvMoveWindow("搜索", 200, 100);

	storage = cvCreateMemStorage(0);  
	search = cvCloneImage(gray);

	int Num = cvFindContours (bin, storage, //搜索
		&contour, sizeof(CvContour), 
		CV_RETR_EXTERNAL, 
		CV_CHAIN_APPROX_SIMPLE,
		cvPoint(0,0)); 

    printf("The number of Contours is: %d\n", Num);
    for(int idx=0; contour!=0; idx++,contour=contour->h_next){  
        // printf("***************************************************\n");  
        for(int i=0; i<contour->total; i++){  
            CvPoint* p=(CvPoint*)cvGetSeqElem(contour,i);  
            // printf("p->x=%d,p->y=%d\n",p->x,p->y);   
	        // CvRect* r = (CvRect*)cvGetSeqElem(contour,i);
	        // printf("x=%d, y=%d width=%d height=%d\n",
	        // 	r->x, r->y, r->width, r->height);   

	        CvRect rect = cvBoundingRect(contour, 0); //连点成面

	        if (rect.x < search->width/2 //开始点大于一半，说明是头像部分了
	        	&& rect.width < search->width * 3/4 //总长度大于3/4可能是长阴影
	        	&& rect.height < search->height * 3/4 //总高度大于3/4可能是高阴影
	        	&& rect.width > 25 && rect.height > 25){
	            IplImage *dst = cvCreateImage(cvSize(rect.width, rect.height), search->depth, search->nChannels);
	            cvSetImageROI(search, rect);//设置原图中的ROI部分
	            cvCopy(search, dst, NULL);    //复制
	            cvResetImageROI(search); //复原
	            char file_name[50];
	            sprintf(file_name, "%s%d%s", "roi", idx, ".jpg");
	            cvSaveImage(file_name, dst, 0); //保存
	            cvReleaseImage(&dst); //释放

	        	cvRectangle(search, //用框框住信息
	        		cvPoint(rect.x, rect.y),
	        		cvPoint(rect.x+rect.width, rect.y+rect.height),
	        		CV_RGB(255,0,0), 1, 8, 0);	            
            } 	
        }  
        // cvSetImageROI(search);//指回原图
        //将轮廓画出   
        // cvDrawContours(search, contour, CV_RGB(255,0,0), CV_RGB(0, 255, 0), 
        // 	           0, 2, 0, cvPoint(0,0));   
	}
	cvShowImage("搜索", search); 
}

/*
void blur_test(){
    IplImage *avgImg = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, img->nChannels);     
	IplImage *medianImg = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, img->nChannels);   
	IplImage *gaussianImg = cvCreateImage(cvGetSize(img), IPL_DEPTH_8U, img->nChannels);   

	cvSmooth(img, avgImg, CV_BLUR, 7,img->nChannels, 0, 0);  //采用7x7的窗口对图像进行均值滤波    
	cvSmooth(img, medianImg, CV_MEDIAN, 7, img->nChannels, 0, 0);  //采用7x7的窗口对图像进行中值滤波    
	cvSmooth(img, gaussianImg, CV_GAUSSIAN, 7, img->nChannels, 0, 0);  //  Gauss平滑滤波，核大小为7x7   
	 //高斯的核不同于上面两个，它实现了领域像素的加权平均，离中心越近的像素权重越高  
	cvShowImage("result1", avgImg);
	cvShowImage("result2", medianImg);
	cvShowImage("result3", gaussianImg);  	
}
*/

void process_key(){
	for(;;){
		int key = cvWaitKey(0);
		if(key == 27) break;
	}
}

/*
void cvShowManyImages(char* title, int nArgs, ...){
    // img - Used for getting the arguments 
    IplImage *img;

    // [[DispImage]] - the image in which input images are to be copied
    IplImage *DispImage;

    int size;
    int i;
    int m, n;
    int x, y;

    // w - Maximum number of images in a row 
    // h - Maximum number of images in a column 
    int w, h;

    // scale - How much we have to resize the image
    float scale;
    int max;

    // If the number of arguments is lesser than 0 or greater than 12
    // return without displaying 
    if(nArgs <= 0) {
        printf("Number of arguments too small....\n");
        return;
    }
    else if(nArgs > 12) {
        printf("Number of arguments too large....\n");
        return;
    }
    // Determine the size of the image, 
    // and the number of rows/cols 
    // from number of arguments 
    else if (nArgs == 1) {
        w = h = 1;
        size = 300;
    }
    else if (nArgs == 2) {
        w = 2; h = 1;
        size = 300;
    }
    else if (nArgs == 3 || nArgs == 4) {
        w = 2; h = 2;
        size = 300;
    }
    else if (nArgs == 5 || nArgs == 6) {
        w = 3; h = 2;
        size = 200;
    }
    else if (nArgs == 7 || nArgs == 8) {
        w = 4; h = 2;
        size = 200;
    }
    else {
        w = 4; h = 3;
        size = 150;
    }

    // Create a new 3 channel image
    DispImage = cvCreateImage( cvSize(100 + size*w, 60 + size*h), 8, 3 );

    // Used to get the arguments passed
    va_list args;
    va_start(args, nArgs);

    // Loop for nArgs number of arguments
    for (i = 0, m = 20, n = 20; i < nArgs; i++, m += (20 + size)) {

        // Get the Pointer to the IplImage
        img = va_arg(args, IplImage*);

        // Check whether it is NULL or not
        // If it is NULL, release the image, and return
        if(img == 0) {
            printf("Invalid arguments");
            cvReleaseImage(&DispImage);
            return;
        }

        // Find the width and height of the image
        x = img->width;
        y = img->height;

        // Find whether height or width is greater in order to resize the image
        max = (x > y)? x: y;

        // Find the scaling factor to resize the image
        scale = (float) ( (float) max / size );

        // Used to Align the images
        if( i % w == 0 && m!= 20) {
            m = 20;
            n+= 20 + size;
        }

        // Set the image ROI to display the current image
        cvSetImageROI(DispImage, cvRect(m, n, (int)( x/scale ), (int)( y/scale )));

        // Resize the input image and copy the it to the Single Big Image
        cvResize(img, DispImage, CV_INTER_LINEAR);

        // Reset the ROI in order to display the next image
        cvResetImageROI(DispImage);
    }

    // Create a new window, and show the Single Big Image
    cvNamedWindow( title, 1 );
    cvShowImage( title, DispImage);

    cvWaitKey(0);
    cvDestroyWindow(title);

    // End the number of arguments
    va_end(args);

    // Release the Image Memory
    cvReleaseImage(&DispImage);
}
*/
#endif