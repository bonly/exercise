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
  glClearColor(0.0, 0.0, 0.0, 0.0); ///����������ɫ
  glShadeModel(GL_FLAT);
}

void display()
{
  glClear(GL_COLOR_BUFFER_BIT); ///����
  glPushMatrix();
  glRotatef(spin, 0.0, 0.0, 1.0);
  glColor3f(0.0, 0.0, 0.5);
  glRectf(-25.0, -25.0, 25.0, 25.0);
  glPopMatrix();
  glutSwapBuffers(); ///������ʾbuffer
}

void reshape(int w, int h)
{
  glViewport(0, 0, (GLsizei)w, (GLsizei)h); ///�����ӽ�(�����ƹ������ӽ�)
  glMatrixMode(GL_PROJECTION); ///�޸ľ���ģʽ
  glLoadIdentity();///ʹ��������Ч
  glOrtho(-50.0, 50.0, -50.0, 50.0, -1.0, 1.0); ///���ý�ȡͼ��Ĵ����С(�Ӿ���)
  glMatrixMode(GL_MODELVIEW);///�޸ľ���ģʽ
  glLoadIdentity();///ʹ��������Ч
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
