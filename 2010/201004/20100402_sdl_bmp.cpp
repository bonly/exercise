//============================================================================
// Name        : try_sdl.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <SDL.h>
using namespace std;

void pressEsc();
void doSomeLoop();

namespace{
  SDL_Surface* pSc = 0;
}

int main(int argc, char* argv[])
{
  try
  {
    //SDL_putenv("SDL_VIDEODRIVER=directx");//Ĭ��Ϊwindib������������SDL_Init()ǰ����ʹ��directx������Ӳ����
    //putenv();SDL_putenv();getenv();SDL_getenv();cygwin���޷�����ʹ��directx�ɹ�
    if(SDL_Init(SDL_INIT_EVERYTHING)==-1)
      throw SDL_GetError();
  }
  catch(const char* s)
  {
    cerr << s << endl;
    return -1;
  }
  cout << "SDL initialized.\n";
  atexit(SDL_Quit);

  pSc = SDL_SetVideoMode(640,480,32,SDL_SWSURFACE);
  cout << "Program is running, press Esc to quit.\n";
  pressEsc();
  cout << "Game over"<< endl;

  return 0;
}

void pressEsc()
{
  bool gameOver = false;
  while(!gameOver)
  {
    SDL_Event evn;

    while(SDL_PollEvent(&evn) != 0)
    {
      switch(evn.type)
      {
        case SDL_QUIT:
          gameOver = true;
          break;
        case SDL_KEYDOWN:
          if(evn.key.keysym.sym == SDLK_ESCAPE)
            gameOver = true;
          break;
      }
      doSomeLoop();
    }
    clog << "after poll" << endl;
  }
  return;
}

void doSomeLoop()
{
   static SDL_Surface* bmp = 0;
   static SDL_Rect pint ;

   if (bmp==0)
   {
     try
     {
        if((bmp = SDL_LoadBMP("res/NeHe.bmp"))==0)
          throw SDL_GetError();
     }
     catch(const char* s)
     {
       cerr << "load bmp failed: " << s << endl;
       SDL_Quit();
     }
   }

   SDL_BlitSurface(bmp, 0, pSc, &pint);///��ȥ��һ��ͼ��ϲ�,SDL_Rect��Ϊ0��ʾ���Ͻ�Ϊ0�غ�
   ++pint.x;
   SDL_Flip(pSc); ///ˢ��
}
/*
 * -L c:/sdl/build/.libs -lSDLmain -lSDL ///SDLmain������SDL��ǰ��
 * cp c:/sdl/build/.libs/SDL.dll /cygdrive/c/WINDOWS/system
 * SDL_Init�����ڵ�������SDL����֮ǰʹ�ã�����SDL_WasInit /SDL_Quit
 */

