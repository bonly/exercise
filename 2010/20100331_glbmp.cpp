/**
 *  @file 20100331_glbmp.cpp
 *
 *  @date 2012-2-18
 *  @Author: Bonly
 *  @bref 代替win下的  auxDIBImageLoad();
 */

#include <GL/glut.h>
#include <GL/glext.h> //for GL_BGR_EXT
#include <cstdio>
#include <cstdlib>
#include <string>
#include <sys/types.h>
using namespace std;

typedef unsigned long DWORD;
typedef unsigned size_t;

unsigned int LoadTex(string Image)
{
  unsigned int Texture;

  FILE* img = NULL;
  img = fopen(Image.c_str(), "rb");

  unsigned long bWidth = 0;
  unsigned long bHeight = 0;
  DWORD size = 0;

  fseek(img, 18, SEEK_SET);///跳过文件信息头(seek_set)至18
  fread((void*)&bWidth, 4, 1, (FILE*)img);
  fread((void*)&bHeight, (size_t)4, 1, (FILE*)img);
  fseek(img, 0, SEEK_END);///到seek_end并移动0位,即定位在文件尾
  size = ftell((FILE*)img) - 54; ///计算图片数据大小

  unsigned char *data = (unsigned char*) malloc(size);

  fseek(img, 54, SEEK_SET); // image data
  fread(data, size, 1, img);//按每次一位的方式读出数据

  fclose(img);

  glGenTextures(1, &Texture); ///分配1个纹理ID到texture地址中
  glBindTexture(GL_TEXTURE_2D, Texture);///以下的操作绑定到纹理ID为texture中
  gluBuild2DMipmaps(GL_TEXTURE_2D, 3, bWidth, bHeight, GL_BGR_EXT,
      GL_UNSIGNED_BYTE, data);
  /**@note GL_TEXTURE_2D说明是2D的图像
            数字3是数据的成分数。因为图像是由红色数据，绿色数据，蓝色数据三种组分组成
            如果您知道宽度，您可以在这里填入，但计算机可以很容易的为您指出此值
     GL_UNSIGNED_BYTE意味着组成图像的数据是无符号字节类型
     @todo 可查考glGenerateMipmapEXT替代
   */

  /**@note 下面的两行告诉OpenGL在显示图像时，
   * 当它比放大得原始的纹理大（GL_TEXTURE_MAG_FILTER）或缩小得比原始得纹理小（GL_TEXTURE_MIN_FILTER）时OpenGL采用的滤波方式。
   * 通常这两种情况下我都采用GL_LINEAR。这使得纹理从很远处到离屏幕很近时都平滑显示。
   * 使用GL_LINEAR需要CPU和显卡做更多的运算。如果您的机器很慢，您也许应该采用GL_NEAREST。
   * 过滤的纹理在放大的时候，看起来斑驳的很（译者：就是马赛克）。您也可以结合这两种滤波方式。
   * 在近处时使用GL_LINEAR，远处时GL_NEAREST。
   */
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);

  if (data)///data已被复制到纹理内存中,此数据无用,可以删除
    free(data);

  return Texture;///返回纹理ID
}
namespace
{
  GLint width = 640;
  GLint height = 480;
  GLint mainWindow;
  GLfloat xrot,yrot,zrot;
  GLuint texture[1];
}

void DrawGLScene();
int LoadGLTextures();
void reshape(int w, int h);

int main(int argc, char* argv[])
{
  glutInit(&argc, argv);
  glutInitDisplayMode(GLUT_RGBA|GLUT_DOUBLE);
  glutInitWindowSize(width, height);
  mainWindow = glutCreateWindow("my test");

  //*
  if(!LoadGLTextures())
    return -1;
  glEnable(GL_TEXTURE_2D); ///开启设置才能显示图片
  glShadeModel(GL_SMOOTH);
  glClearColor(0.0, 0.0, 0.0, 0.0);
  glClearDepth(1.0f);
  glEnable(GL_DEPTH_TEST);
  glDepthFunc(GL_LEQUAL);
  glHint(GL_PERSPECTIVE_CORRECTION_HINT, GL_NICEST);
  //*/

  glutDisplayFunc(DrawGLScene);
  glutReshapeFunc(reshape);

  glutMainLoop();

  return 0;
}

void reshape(int w, int h)
{
  ///定义视角
  glViewport(0, 0, (GLsizei)w, (GLsizei)h);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  //glOrtho(-50.0, 50.0, -50.0, 50.0, -1.0, 1.0);///用于定义视景glOrtho,在此不适合
  gluPerspective(45.0f,(GLfloat)width/(GLfloat)height,0.1f,100.0f);
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();
}

void DrawGLScene(int notuse)
{
  glClear(GL_COLOR_BUFFER_BIT|GL_DEPTH_BUFFER_BIT);
  glPushMatrix();
  //glRotatef(spin, 0.0, 0.0, 1.0);
  glColor3f(1.0, 0.0, 0.0);
  glBegin(GL_POLYGON);
    glVertex2f(0.0, 0.0);
    glVertex2f(7.0, 7.0);
    glVertex2f(2.0, 13.0);
  glEnd();
  glPopMatrix();
  glutSwapBuffers();
}

/**
@note glTranslatef(x, y, z)
沿着 X, Y 和 Z 轴移动。
注意在glTranslatef(x, y, z)中,当您移动的时候，您并不是相对屏幕中心移动，
而是相对与当前所在的屏幕位置。其作用就是将你绘点坐标的原点在当前原点的基础上平移一个(x,y,z)向量。
 */
void DrawGLScene()
{
  glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
  glLoadIdentity();
  glTranslatef(0.0f, 0.0f, -5.0f);

  glRotatef(xrot, 1.0f, 0.0f, 0.0f);
  glRotatef(yrot, 0.0f, 1.0f, 0.0f);
  glRotatef(zrot, 0.0f, 0.0f, 1.0f);

  glBindTexture(GL_TEXTURE_2D, texture[0]);

  //glColor3f(1.0, 0.0, 0.0);
  glBegin(GL_QUADS);
    // Front Face
    glTexCoord2f(0.0f, 0.0f);  glVertex3f(-1.0f, -1.0f,  1.0f);
    glTexCoord2f(1.0f, 0.0f);  glVertex3f( 1.0f, -1.0f,  1.0f);
    glTexCoord2f(1.0f, 1.0f);  glVertex3f( 1.0f,  1.0f,  1.0f);
    glTexCoord2f(0.0f, 1.0f);  glVertex3f(-1.0f,  1.0f,  1.0f);
    // Back Face
    glTexCoord2f(1.0f, 0.0f);  glVertex3f(-1.0f, -1.0f, -1.0f);
    glTexCoord2f(1.0f, 1.0f);  glVertex3f(-1.0f,  1.0f, -1.0f);
    glTexCoord2f(0.0f, 1.0f);  glVertex3f( 1.0f,  1.0f, -1.0f);
    glTexCoord2f(0.0f, 0.0f);  glVertex3f( 1.0f, -1.0f, -1.0f);
    // Top Face
    glTexCoord2f(0.0f, 1.0f);  glVertex3f(-1.0f,  1.0f, -1.0f);
    glTexCoord2f(0.0f, 0.0f);  glVertex3f(-1.0f,  1.0f,  1.0f);
    glTexCoord2f(1.0f, 0.0f);  glVertex3f( 1.0f,  1.0f,  1.0f);
    glTexCoord2f(1.0f, 1.0f);  glVertex3f( 1.0f,  1.0f, -1.0f);
    // Bottom Face
    glTexCoord2f(1.0f, 1.0f);  glVertex3f(-1.0f, -1.0f, -1.0f);
    glTexCoord2f(0.0f, 1.0f);  glVertex3f( 1.0f, -1.0f, -1.0f);
    glTexCoord2f(0.0f, 0.0f);  glVertex3f( 1.0f, -1.0f,  1.0f);
    glTexCoord2f(1.0f, 0.0f);  glVertex3f(-1.0f, -1.0f,  1.0f);
    // Right face
    glTexCoord2f(1.0f, 0.0f); glVertex3f( 1.0f, -1.0f, -1.0f);
    glTexCoord2f(1.0f, 1.0f); glVertex3f( 1.0f,  1.0f, -1.0f);
    glTexCoord2f(0.0f, 1.0f); glVertex3f( 1.0f,  1.0f,  1.0f);
    glTexCoord2f(0.0f, 0.0f); glVertex3f( 1.0f, -1.0f,  1.0f);
    // Left Face
    glTexCoord2f(0.0f, 0.0f); glVertex3f(-1.0f, -1.0f, -1.0f);
    glTexCoord2f(1.0f, 0.0f); glVertex3f(-1.0f, -1.0f,  1.0f);
    glTexCoord2f(1.0f, 1.0f); glVertex3f(-1.0f,  1.0f,  1.0f);
    glTexCoord2f(0.0f, 1.0f); glVertex3f(-1.0f,  1.0f, -1.0f);
  glEnd();

  glutSwapBuffers();
  xrot+=0.3f;
  yrot+=0.2f;
  zrot+=0.4f;
  return;

}

int LoadGLTextures()
{
  texture[0] = LoadTex("res/NeHe.bmp");
  if (texture[0] == 0)
    return false;
  return true;
}
/**
加载纹理与使用glGenTextures时应注意的一点(解决吃内存)[转]
glGenTextures
　　glGenTextures(GLsizei n, GLuint *textures)函数说明

n：用来生成纹理的数量

textures：存储纹理索引的

glGenTextures函数根据纹理参数返回n个纹理索引。纹理名称集合不必是一个连续的整数集合。

（glGenTextures就是用来产生你要操作的纹理对象的索引的，比如你告诉OpenGL，我需要5个纹理对象，它会从没有用到的整数里返回5个给你）

glBindTexture实际上是改变了OpenGL的这个状态，它告诉OpenGL下面对纹理的任何操作都是对它所绑定的纹理对象的，比如glBindTexture(GL_TEXTURE_2D,1)告诉OpenGL下面代码中对2D纹理的任何设置都是针对索引为1的纹理的。

产生纹理函数假定目标纹理的面积是由glBindTexture函数限制的。先前调用glGenTextures产生的纹理索引集不会由后面调用的glGenTextures得到，除非他们首先被glDeleteTextures删除。你不可以在显示列表中包含glGenTextures。
void glGenTextures(GLsizei n, GLuint *texture);
该函数用来产生纹理名称。这里纹理名称GLuint *texture是整型的，因此也可以理解为这个函数为这n个纹理指定了n个不同的ID。
在用GL渲染的时候，纹理是很常见的东西。使用纹理之前，必须执行这句命令为你的texture分配一个ID，然后绑定这个纹理，加载纹理图像，这之后，这个纹理才可以使用。加载纹理的代码如下：
BOOL LoadTextures(IplImage *pImage, GLuint *pTexture)
{
    int Status=FALSE;
    if(pImage != NULL)
    {
        Status=TRUE;
        glGenTextures(1, &pTexture[0]); //注意这里
        glBindTexture(GL_TEXTURE_2D, pTexture[0]);
        glTexImage2D(GL_TEXTURE_2D, 0, 3,
                     pImage->width, pImage->height,
                     0, GL_BGR, GL_UNSIGNED_BYTE, (unsigned char *)pImage->imageData);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MIN_FILTER,GL_LINEAR);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MAG_FILTER,GL_LINEAR);
    }
    return Status;
}
    使用上面这个函数时需要小心，这个函数只能放在循环外面使用！如果你想在循环中重复利用这个texture[0]，给它加载不同的纹理（比如，你想在窗口中显示序列图像），而把这个函数放在循环内部调用的话，那么当程序循环足够多次之后，你的电脑将变得巨慢无比，甚至导致死机。原因就是反复地调用glGenTextures(1, &pTexture[0])。这个问题产生的机制我并不清楚，但是我今天实实在在的遇到了。
    所以，上面这个函数一般都是放在循环外面，窗口初始化的时候，用于给背景加载纹理。那么，如果我必须要在循环中渲染序列帧的话，该怎么做呢？我们可以对上面的函数加一点小小的改变，如下：
BOOL LoadTextures(IplImage *pImage, GLuint texture)
{
    int Status=FALSE;
    if(pImage != NULL)
    {
        Status=TRUE;
        glBindTexture(GL_TEXTURE_2D, texture);
        glTexImage2D(GL_TEXTURE_2D, 0, 3,
                     pImage->width, pImage->height,
                     0, GL_BGR, GL_UNSIGNED_BYTE, (unsigned char *)pImage->imageData);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MIN_FILTER,GL_LINEAR);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MAG_FILTER,GL_LINEAR);
    }
    return Status;
}
在窗口初始化的时候先执行一遍：
glGenTextures(1, &texture[0]);
然后在你的循环内部调用：
IplImage *videoFrame = cvQueryFrame(capture);
LoadTextures（videoFrame, texture[0]）；
这样就可以显示图像帧了，也不会再出现电脑运行速度变慢的问题了。总之，千万不要给一个texture重复分配ID。
    我自己写的这个LoadTextures函数提供了图像buffer的接口，可以从外面读取视频帧并传给这个函数，绑定纹理, 使用起来比较灵活。
*/

