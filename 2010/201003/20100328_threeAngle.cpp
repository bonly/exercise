/*
 * double.cpp
 *
 *  Created on: 2012-2-16
 *      Author: Bonly
 */
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
  glClearColor(0.0, 0.0, 0.0, 0.0);
  glShadeModel(GL_FLAT);
}

/**
@note glRotatef(angle, x, y, z)
��ת�ᾭ��ԭ�㣬����Ϊ(x,y,z),��ת�Ƕ�Ϊangle�������������ֶ���
 */
void display()
{
  glClear(GL_COLOR_BUFFER_BIT);
  glPushMatrix();
  glRotatef(spin, 0.0, 0.0, 1.0);
  glColor3f(1.0, 0.0, 0.0);
  glBegin(GL_POLYGON);
    glVertex2f(0.0, 0.0);
    glVertex2f(7.0, 7.0);
    glVertex2f(2.0, 13.0);
  glEnd();
  glPopMatrix();
  glutSwapBuffers();
}

void reshape(int w, int h)
{
  glViewport(0, 0, (GLsizei)w, (GLsizei)h);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  glOrtho(-50.0, 50.0, -50.0, 50.0, -1.0, 1.0);
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();
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

/**
@note glLoadIdentity
����ǰ���û�����ϵ��ԭ���Ƶ�����Ļ���ģ�������һ����λ����
1.X������������ң�Y������������ϣ�Z������������⡣
2.OpenGL��Ļ���ĵ�����ֵ��X��Y���ϵ�0.0f�㡣
3.�������������ֵ�Ǹ�ֵ����������ֵ��
   ������Ļ��������ֵ��������Ļ�׶��Ǹ�ֵ��
   ������Ļ��Ǹ�ֵ���Ƴ���Ļ������ֵ��
 */
