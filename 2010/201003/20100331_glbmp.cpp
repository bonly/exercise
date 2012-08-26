/**
 *  @file 20100331_glbmp.cpp
 *
 *  @date 2012-2-18
 *  @Author: Bonly
 *  @bref ����win�µ�  auxDIBImageLoad();
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

  fseek(img, 18, SEEK_SET);///�����ļ���Ϣͷ(seek_set)��18
  fread((void*)&bWidth, 4, 1, (FILE*)img);
  fread((void*)&bHeight, (size_t)4, 1, (FILE*)img);
  fseek(img, 0, SEEK_END);///��seek_end���ƶ�0λ,����λ���ļ�β
  size = ftell((FILE*)img) - 54; ///����ͼƬ���ݴ�С

  unsigned char *data = (unsigned char*) malloc(size);

  fseek(img, 54, SEEK_SET); // image data
  fread(data, size, 1, img);//��ÿ��һλ�ķ�ʽ��������

  fclose(img);

  glGenTextures(1, &Texture); ///����1������ID��texture��ַ��
  glBindTexture(GL_TEXTURE_2D, Texture);///���µĲ����󶨵�����IDΪtexture��
  gluBuild2DMipmaps(GL_TEXTURE_2D, 3, bWidth, bHeight, GL_BGR_EXT,
      GL_UNSIGNED_BYTE, data);
  /**@note GL_TEXTURE_2D˵����2D��ͼ��
            ����3�����ݵĳɷ�������Ϊͼ�����ɺ�ɫ���ݣ���ɫ���ݣ���ɫ��������������
            �����֪����ȣ����������������룬����������Ժ����׵�Ϊ��ָ����ֵ
     GL_UNSIGNED_BYTE��ζ�����ͼ����������޷����ֽ�����
     @todo �ɲ鿼glGenerateMipmapEXT���
   */

  /**@note ��������и���OpenGL����ʾͼ��ʱ��
   * �����ȷŴ��ԭʼ�������GL_TEXTURE_MAG_FILTER������С�ñ�ԭʼ������С��GL_TEXTURE_MIN_FILTER��ʱOpenGL���õ��˲���ʽ��
   * ͨ��������������Ҷ�����GL_LINEAR����ʹ������Ӻ�Զ��������Ļ�ܽ�ʱ��ƽ����ʾ��
   * ʹ��GL_LINEAR��ҪCPU���Կ�����������㡣������Ļ�����������Ҳ��Ӧ�ò���GL_NEAREST��
   * ���˵������ڷŴ��ʱ�򣬿������߲��ĺܣ����ߣ����������ˣ�����Ҳ���Խ���������˲���ʽ��
   * �ڽ���ʱʹ��GL_LINEAR��Զ��ʱGL_NEAREST��
   */
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);

  if (data)///data�ѱ����Ƶ������ڴ���,����������,����ɾ��
    free(data);

  return Texture;///��������ID
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
  glEnable(GL_TEXTURE_2D); ///�������ò�����ʾͼƬ
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
  ///�����ӽ�
  glViewport(0, 0, (GLsizei)w, (GLsizei)h);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  //glOrtho(-50.0, 50.0, -50.0, 50.0, -1.0, 1.0);///���ڶ����Ӿ�glOrtho,�ڴ˲��ʺ�
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
���� X, Y �� Z ���ƶ���
ע����glTranslatef(x, y, z)��,�����ƶ���ʱ���������������Ļ�����ƶ���
��������뵱ǰ���ڵ���Ļλ�á������þ��ǽ����������ԭ���ڵ�ǰԭ��Ļ�����ƽ��һ��(x,y,z)������
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
����������ʹ��glGenTexturesʱӦע���һ��(������ڴ�)[ת]
glGenTextures
����glGenTextures(GLsizei n, GLuint *textures)����˵��

n�������������������

textures���洢����������

glGenTextures�������������������n�������������������Ƽ��ϲ�����һ���������������ϡ�

��glGenTextures��������������Ҫ�������������������ģ����������OpenGL������Ҫ5��������������û���õ��������ﷵ��5�����㣩

glBindTextureʵ�����Ǹı���OpenGL�����״̬��������OpenGL�����������κβ������Ƕ������󶨵��������ģ�����glBindTexture(GL_TEXTURE_2D,1)����OpenGL��������ж�2D������κ����ö����������Ϊ1������ġ�

�����������ٶ�Ŀ��������������glBindTexture�������Ƶġ���ǰ����glGenTextures���������������������ɺ�����õ�glGenTextures�õ��������������ȱ�glDeleteTexturesɾ�����㲻��������ʾ�б��а���glGenTextures��
void glGenTextures(GLsizei n, GLuint *texture);
�ú������������������ơ�������������GLuint *texture�����͵ģ����Ҳ�������Ϊ�������Ϊ��n������ָ����n����ͬ��ID��
����GL��Ⱦ��ʱ�������Ǻܳ����Ķ�����ʹ������֮ǰ������ִ���������Ϊ���texture����һ��ID��Ȼ������������������ͼ����֮���������ſ���ʹ�á���������Ĵ������£�
BOOL LoadTextures(IplImage *pImage, GLuint *pTexture)
{
    int Status=FALSE;
    if(pImage != NULL)
    {
        Status=TRUE;
        glGenTextures(1, &pTexture[0]); //ע������
        glBindTexture(GL_TEXTURE_2D, pTexture[0]);
        glTexImage2D(GL_TEXTURE_2D, 0, 3,
                     pImage->width, pImage->height,
                     0, GL_BGR, GL_UNSIGNED_BYTE, (unsigned char *)pImage->imageData);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MIN_FILTER,GL_LINEAR);
        glTexParameteri(GL_TEXTURE_2D,GL_TEXTURE_MAG_FILTER,GL_LINEAR);
    }
    return Status;
}
    ʹ�������������ʱ��ҪС�ģ��������ֻ�ܷ���ѭ������ʹ�ã����������ѭ�����ظ��������texture[0]���������ز�ͬ���������磬�����ڴ�������ʾ����ͼ�񣩣����������������ѭ���ڲ����õĻ�����ô������ѭ���㹻���֮����ĵ��Խ���þ����ޱȣ���������������ԭ����Ƿ����ص���glGenTextures(1, &pTexture[0])�������������Ļ����Ҳ�������������ҽ���ʵʵ���ڵ������ˡ�
    ���ԣ������������һ�㶼�Ƿ���ѭ�����棬���ڳ�ʼ����ʱ�����ڸ���������������ô������ұ���Ҫ��ѭ������Ⱦ����֡�Ļ�������ô���أ����ǿ��Զ�����ĺ�����һ��СС�ĸı䣬���£�
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
�ڴ��ڳ�ʼ����ʱ����ִ��һ�飺
glGenTextures(1, &texture[0]);
Ȼ�������ѭ���ڲ����ã�
IplImage *videoFrame = cvQueryFrame(capture);
LoadTextures��videoFrame, texture[0]����
�����Ϳ�����ʾͼ��֡�ˣ�Ҳ�����ٳ��ֵ��������ٶȱ����������ˡ���֮��ǧ��Ҫ��һ��texture�ظ�����ID��
    ���Լ�д�����LoadTextures�����ṩ��ͼ��buffer�Ľӿڣ����Դ������ȡ��Ƶ֡���������������������, ʹ�������Ƚ���
*/

