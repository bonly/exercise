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

int main(int argc, char* argv[])
{
  try
  {
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

  SDL_SetVideoMode(640,480,32,SDL_SWSURFACE);
  /** @note SDL_SWSURFACE=0
   * 函数返回SDL_Surface*,所指的结构数据储存在系统内存中,
   * SDL_HWSURFACE=1 数据存在显存中
   */
  std::cout << "Program is running, press Esc to quit.\n";
  pressEsc();
  std::cout << "Game over"<< std::endl;

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
  }
  return;
}

void doSomeLoop()
{
  static int a=0;
  for(int i=0; i<a; ++i)
  {
     std::cout << ".";
  }
  std::cout << endl;
}
/*
 * -L c:/sdl/build/.libs -lSDLmain -lSDL ///SDLmain必须在SDL库前面
 * cp c:/sdl/build/.libs/SDL.dll /cygdrive/c/WINDOWS/system
 * SDL_Init必需在调用其它SDL函数之前使用，除了SDL_WasInit /SDL_Quit
 */

