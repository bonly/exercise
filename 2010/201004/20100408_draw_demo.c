#include "draw_demo.h"
#include "draw.h"

int init(void);
SDL_Surface *screen;
int main(void)
{
	Uint32 yellow;
	Uint32 red;

	if(!init()) {
		printf("Init error:%s\n",SDL_GetError());
		return 1;
	}
	yellow = SDL_MapRGB(screen->format, 0xff, 0xff, 0x00);
	red = SDL_MapRGB(screen->format, 0xff, 0x00, 0x00);
	
	Draw_Line(screen,100,100,50,50,yellow);
	Draw_Line(screen,110,100,10,200,yellow);
	Draw_Line(screen,120,120,400,400,yellow);
	Draw_Line(screen,1,1,1,50,yellow);
	Draw_Rect(screen,50,50,120,120,yellow);
	Draw_Line(screen,20,30,460,200,yellow);
	Draw_Line(screen,480,10,10,400,yellow);
	Draw_Rect_Fill(screen,300,300,380,380,red,red);
	Draw_Rect_Fill(screen,300,300,50,360,yellow,red);
	Draw_Rect_Fill(screen,300,300,360,60,yellow,red);
	Draw_Circle_Fill(screen,200,200,50,red,yellow);
	Draw_Circle(screen,400,400,10,red);
	Draw_Circle_Fill(screen,300,300,50,yellow,red);
	Draw_Ellipse_Fill(screen,160,160,200,200,120,red,yellow);

	SDL_Delay(3000);
	return 0;
}

int init(void)
{
	if(SDL_Init(SDL_INIT_EVERYTHING) == -1 ) {
		return FALSE;    
	}

	screen = SDL_SetVideoMode(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_BPP, SDL_SWSURFACE); 
	if(screen == NULL) {
		return FALSE;    
	}

	SDL_WM_SetCaption( "SDL Draw  by cocobear 2007.10", NULL );

	return TRUE;
}
