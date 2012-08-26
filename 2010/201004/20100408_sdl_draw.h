/**
 *  @file sdl_draw.h
 *
 *  @date 2012-2-21
 *  @Author: Bonly
 */

#ifndef SDL_DRAW_H_
#define SDL_DRAW_H_

void put_pixel(SDL_Surface *surface, int x, int y, Uint32 pixel);
Uint32 get_pixel(SDL_Surface *surface, int x, int y);
void draw_point(SDL_Surface *screen,int x,int y,Uint32 color);
void plot_circle_points(SDL_Surface *screen,int xc,int yc,int x,int y,Uint32 c);
void bresenham_line(SDL_Surface *screen,int x1,int y1,int x2,int y2,Uint32 c);
void dda_line(SDL_Surface *screen,int xa,int ya,int xb,int yb,Uint32 c);
void seed_filling(SDL_Surface *screen,int i,int j,Uint32 color_fill,Uint32 boundary_color);


#endif /* SDL_DRAW_H_ */
