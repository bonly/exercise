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
		SDL_putenv("SDL_VIDEODRIVER=directx");//默认为windib驱动，必须在SDL_Init()前设置使用directx，才能硬加速
		//putenv();SDL_putenv();getenv();SDL_getenv();cygwin下无法测试使用directx成功
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
	screen = SDL_SetVideoMode(480,272,24,SDL_HWSURFACE);
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

	char driverName[20];
	cout << "SDL_VideoDriverName = " << SDL_VideoDriverName(driverName,20) << endl;

	const SDL_VideoInfo* myInfo = SDL_GetVideoInfo();
	cout << "hardware surface? " << myInfo->hw_available<<endl;
	cout << "available window manager? "<<myInfo->wm_available<<endl;
	cout << "hardware to hardware blits accelerated? "<<myInfo->blit_hw<<endl;
	cout << "hardware to hardware colorkey blits accelerated? "<<myInfo->blit_hw_CC<<endl;
	cout << "hardware to hardware alpha blits accelerated? "<<myInfo->blit_hw_A<<endl;
	cout << "software to hardware blits accelerated? "<< myInfo->blit_sw<<endl;
	cout << "software to hardware colorkey blits accelerated? "<< myInfo->blit_sw_CC << endl;
	cout << "software to hardware alpha blits accelerated? "<< myInfo->blit_sw_A<<endl;
	cout << "color fills accelerated? "<<myInfo->blit_fill<<endl;
	cout << "Total amount of video memory in K? "<< myInfo->video_mem<<endl;
	cout << "width: "<<myInfo->current_w<<endl;
	cout << "Height: "<<myInfo->current_h<<endl;
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


