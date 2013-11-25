#include "cv.h"
#include "highgui.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <math.h>
#include <float.h>
#include <limits.h>
#include <time.h>
#include <ctype.h>
using namespace std;

static CvMemStorage* storage = 0;
static CvHaarClassifierCascade* cascade = 0;

void detect_and_draw( IplImage* image );

const char* cascade_name =
"/home/bonly/opt/opencv/share/OpenCV/haarcascades/haarcascade_frontalface_alt.xml";
/* "haarcascade_profileface.xml";*/

int main(int argc, char* argv[])
{
        CvCapture* capture = 0;

        cascade = (CvHaarClassifierCascade*)cvLoad( cascade_name, 0, 0, 0 );

        if( !cascade )
        {
                fprintf( stderr, "ERROR: Could not load classifier cascade/n" );
                //fprintf( stderr,
                        //"Usage: facedetect --cascade=/"<cascade_path>"/[filename|camera_index]/n" );
                return -1;
        }
        storage = cvCreateMemStorage(0);


        cvNamedWindow( "result", 1 );


        const char* filename = "emp.jpg";
        //const char* filename = argv[1];
        IplImage* image = cvLoadImage(filename );

        if( image )
        {
                detect_and_draw( image );
                cvWaitKey(0);
                cvReleaseImage( &image );
        }

        cvDestroyWindow("result");
        cvWaitKey(0);
        return 0;
}

void detect_and_draw( IplImage* img )
{
        static CvScalar colors[] = 
        {
                {{0,0,255}},
                {{0,128,255}},
                {{0,255,255}},
                {{0,255,0}},
                {{255,128,0}},
                {{255,255,0}},
                {{255,0,0}},
                {{255,0,255}}
        };

        double scale = 1.3;
        IplImage* gray = cvCreateImage( cvSize(img->width,img->height), 8, 1 );
        IplImage* small_img = cvCreateImage( cvSize( cvRound (img->width/scale),
                cvRound (img->height/scale)),
                8, 1 );
        int i;

        cvCvtColor( img, gray, CV_BGR2GRAY );
        cvResize( gray, small_img, CV_INTER_LINEAR );
        cvEqualizeHist( small_img, small_img );
        cvClearMemStorage( storage );

        if( cascade )
        {
                double t = (double)cvGetTickCount();
                CvSeq* faces = cvHaarDetectObjects( small_img, cascade, storage,
                        1.1, 2, 0/*CV_HAAR_DO_CANNY_PRUNING*/,
                        cvSize(30, 30) );
                t = (double)cvGetTickCount() - t;
                printf( "detection time = %gms/n", t/((double)cvGetTickFrequency()*1000.) );
                for( i = 0; i < (faces ? faces->total : 0); i++ )
                {
                        CvRect* r = (CvRect*)cvGetSeqElem( faces, i );
                        CvPoint center;
                        int radius;
                        center.x = cvRound((r->x + r->width*0.5)*scale);
                        center.y = cvRound((r->y + r->height*0.5)*scale);
                        radius = cvRound((r->width + r->height)*0.25*scale);
                        cvCircle( img, center, radius, colors[i%8], 3, 8, 0 );
                }
        }

        cvShowImage( "result", img );
        cvReleaseImage( &gray );
        cvReleaseImage( &small_img );
} 
/*
g++ 20110230_opencv_face.cpp -I /home/opt/opencv/include/opencv -L /home/opt/opencv/lib/ -lopencv_core -lopencv_photo -lopencv_calib3d -lopencv_imgproc -lopencv_highgui  -lopencv_objdetect


g++ 20110230_opencv_face.cpp -I /usr/include/opencv/ -L /usr/lib -lopencv_core -lopencv_calib3d -lopencv_imgproc -lopencv_highgui  -lopencv_objdetect
*/

