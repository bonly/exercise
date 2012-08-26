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

void pressEsc();
void someLoop();

int
main (int argc, char* argv[])
{
	try{
		if(SDL_Init(SDL_INIT_VIDEO==-1))
			throw SDL_GetError();
	}
	catch(const char* s)
	{
		cerr << s << endl;
	  return -1;
	}
	atexit(SDL_Quit);

	SDL_Surface* screen=0;
	screen = SDL_SetVideoMode(480,272,24,SDL_SWSURFACE);
	cout << "Program is running, press Esc to quit\n";
	SDL_Surface* bmp=0;
	bmp = SDL_LoadBMP ("may.bmp");
	if (bmp==0)
	{
		cerr << SDL_GetError();
		exit(-1);
	}
	SDL_Rect* srcRect=0;
	SDL_Rect* dstRect=0;
	if(SDL_BlitSurface(bmp,srcRect,screen,dstRect)!=0)
	{
		cerr << SDL_GetError();
		exit(-1);
	}
	if (SDL_Flip(screen)!=0)
	{
		cerr << SDL_GetError();
		exit(-1);
	}
	pressEsc();
	cout << "Quit\n";

	return 0;
}

void pressEsc()
{
	cout << "Press Esc function begin...\n";
	bool over = false;
	while (over == false)
	{
		SDL_Event event;
		SDL_PollEvent(&event);
		if(&event!=0)
		{
			switch(event.type)
			{
				case SDL_QUIT:
				  over=true;break;
				case SDL_KEYDOWN:
					if(event.key.keysym.sym==SDLK_ESCAPE)
						over=true;
					break;
				default :
					someLoop();
					break;
			}
		}
	}
	return;
}

void someLoop()
{
	//cout << ".";
	return;
}


