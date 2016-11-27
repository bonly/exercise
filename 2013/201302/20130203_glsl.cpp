#include <GL/glew.h>
#include <GL/glut.h>
#include <cstdio>

int main(int argc, char **argv)
{
    glutInit(&argc, argv);
    // glutInitDisplayMode(GLUT_DEPTH | GLUT_DOUBLE | GLUT_RGBA);
    // glutInitWindowPosition(100,100);
    // glutInitWindowSize(320,320);
    glutCreateWindow("MM 2004-05"); //need to create window,and then success

    glewInit();

    if (glewIsSupported("GL_VERSION_2_0"))
        printf("Ready for OpenGL 2.0\n");
    else
    {
        printf("OpenGL 2.0 not supported\n");
        exit(1);
    }
    // setShaders();

    glutMainLoop();
    return 0;
}

/*
g++ -lX11 -lXi -lXmu -lglut -lGL -lGLU -lm -lGLEW
*/