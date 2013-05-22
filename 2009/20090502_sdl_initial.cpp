//============================================================================
// Name        : try_sdl.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include "SDL.h"
using namespace std;
int main()
{
	try
	{
		if(SDL_Init(SDL_INIT_EVERYTHING)==-1)
			throw "Could not initialize SDL!";
	}
	catch(const char* s)
	{
		cerr << s << endl;
		return -1;
	}
	cout << "SDL initialized.\n";

	Uint32 subsystem_init;
  subsystem_init = SDL_WasInit(SDL_INIT_EVERYTHING);
  if(subsystem_init&SDL_INIT_VIDEO)
  	cout << "Video is initialized.\n";
  else
  	cout << "Video is not initialized.\n";

  if (SDL_WasInit(SDL_INIT_VIDEO)!=0)
  	cout << "Video is initialized.\n";
  else
  	cout << "Video is not initialized.\n";

  Uint32 subsystem_mask = SDL_INIT_VIDEO|SDL_INIT_AUDIO;
  if(SDL_WasInit(subsystem_mask)==subsystem_mask)
  	cout << "Video and Audio is initialized.\n";
  else
  	cout << "Video and Audio is not initialized.\n";
	SDL_Quit();
	return 0;
}
/*
 * -L c:/sdl/build/.libs -lSDL
 * cp c:/sdl/build/.libs/SDL.dll /cygdrive/c/WINDOWS/system
 * SDL_WasInit 可以在SDL_Init 之前调用
 */

