#ifndef __PASSPORT_H__
#define __PASSPORT_H__

// void cvShowManyImages(char* title, int nArgs, ...);
void process_key();
void Origin_img();
void Resize();
void Rgb2Gray();
void GaussianBlur(); //高斯模糊
void MorphologyEx_gradient(); //形态梯度
void Sobel();//边缘化
void MorphologyEx_close();//形态关闭
void Threshold_Bin();//二化
void Erode(); //侵蚀
void Find();
// void blur();
#endif