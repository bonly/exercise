/*
 * double.cpp
 *
 *  Created on: 2012-2-16
 *      Author: Bonly
 */
//#define GLUT_DISABLE_ATEXIT_HACK
#include <GL/glut.h>

static GLfloat spin = 0.0;

void init();
void display();
void reshape(int w, int h);
void mouse(int button, int state, int x, int y);

void spinDisplay();

int main(int argc, char* argv[])
{
  glutInit(&argc, argv);
  glutInitDisplayMode(GLUT_DOUBLE|GLUT_RGB);
  glutInitWindowSize(250,250);
  glutInitWindowPosition(100,100);
  glutCreateWindow(argv[0]);

  init();
  glutDisplayFunc(display);
  glutReshapeFunc(reshape);
  glutMouseFunc(mouse);
  glutMainLoop();

  return 0;
}

void init()
{
  glClearColor(0.0, 0.0, 0.0, 0.0); ///设置清屏颜色
  glShadeModel(GL_FLAT);
}

void display()
{
  glClear(GL_COLOR_BUFFER_BIT); ///清屏
  glPushMatrix();
  glRotatef(spin, 0.0, 0.0, 1.0);
  glColor3f(0.0, 0.0, 0.5);
  glRectf(-25.0, -25.0, 25.0, 25.0);
  glPopMatrix();
  glutSwapBuffers(); ///交换显示buffer
}

void reshape(int w, int h)
{
  glViewport(0, 0, (GLsizei)w, (GLsizei)h); ///设置视角(可类似哈哈镜视角)
  glMatrixMode(GL_PROJECTION); ///修改矩阵模式
  glLoadIdentity();///使用特征生效
  glOrtho(-50.0, 50.0, -50.0, 50.0, -1.0, 1.0); ///设置截取图像的窗体大小(视景体)
  glMatrixMode(GL_MODELVIEW);///修改矩阵模式
  glLoadIdentity();///使用特征生效
}

void mouse(int button, int state, int x, int y)
{
  switch(button)
  {
    case GLUT_LEFT_BUTTON:
      if(state == GLUT_DOWN)
        glutIdleFunc(spinDisplay);
      break;
    case GLUT_RIGHT_BUTTON:
      if (state == GLUT_DOWN)
        glutIdleFunc(NULL);
      break;
    default:
      break;
  }
}

void spinDisplay()
{
  spin = spin + 2.0;
  if (spin > 360.0)
    spin = spin - 360.0;
  glutPostRedisplay();
}
