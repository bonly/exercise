#ifndef __DRAW_H__
#define __DRAW_H__
#include <math.h>
#include <SDL.h>

#define NUM_LEVELS 256
#define ABS(x) ((x) > 0 ? (x) : (-x))
#define EPSILON 0.5
void Draw_Pixel(SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 color);
void Draw_Line (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 x1, Uint32 y1, Uint32 color);
void Draw_Rect (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 x1, Uint32 y1, Uint32 color);
void Draw_Rect_Fill (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 x1, Uint32 y1, Uint32 color, Uint32 bcolor);
void Draw_Circle (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 r, Uint32 color);
void Draw_Circle_Fill (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 r, Uint32 color,Uint32 bcolor);
void Draw_Ellipse (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 x1, Uint32 y1, Uint32 r, Uint32 color);
void Draw_Ellipse_Fill (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 x1, Uint32 y1, Uint32 r, Uint32 color, Uint32 bcolor);

double p2p(Uint32 i, Uint32 j, Uint32 x, Uint32 y);

#endif
