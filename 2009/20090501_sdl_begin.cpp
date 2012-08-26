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
	SDL_Quit();
	return 0;
}
/*
 * -L c:/sdl/build/.libs -lSDL
 * cp c:/sdl/build/.libs/SDL.dll /cygdrive/c/WINDOWS/system
 * SDL_Init必需在调用其它SDL函数之前使用，除了SDL_WasInit /SDL_Quit
 */

