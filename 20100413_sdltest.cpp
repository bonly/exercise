/**
 *  @file 201004.cpp
 *
 *  @date 2012-2-24
 *  @Author: Bonly
 */
#include <SDL.h>
#include <cstdio>

SDL_Surface *pSc = 0;

int main(int argc, char* argv[])
{
  try
  {
    if (SDL_Init(SDL_INIT_EVERYTHING) == -1)
      throw SDL_GetError();
  }
  catch (const char* s)
  {
    printf("%s", s);
    return -1;
  }
  atexit(SDL_Quit);

  int flags = SDL_SWSURFACE;//|SDL_FULLSCREEN;
  pSc = SDL_SetVideoMode(400, 600, 32, flags);

  bool run = true;
  while (run)
  {
    SDL_Event evn;
    if (SDL_PollEvent(&evn) != 0)
    {
      switch (evn.type)
      {
        case SDL_QUIT:
          run = false;
          break;
      }
    }

    SDL_Flip(pSc);
    SDL_Delay(50); ///等待一段时间,以便控制帧数
  }
  return 0;
}

