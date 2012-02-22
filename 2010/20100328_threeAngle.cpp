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
旋转轴经过原点，方向为(x,y,z),旋转角度为angle，方向满足右手定则。
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
将当前的用户坐标系的原点移到了屏幕中心：类似于一个复位操作
1.X坐标轴从左至右，Y坐标轴从下至上，Z坐标轴从里至外。
2.OpenGL屏幕中心的坐标值是X和Y轴上的0.0f点。
3.中心左面的坐标值是负值，右面是正值。
   移向屏幕顶端是正值，移向屏幕底端是负值。
   移入屏幕深处是负值，移出屏幕则是正值。
 */
