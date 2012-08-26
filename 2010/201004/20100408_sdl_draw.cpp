/**
 *  @file sdl_draw.cpp
 *
 *  @date 2012-2-21
 *  @Author: Bonly
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <SDL.h>
#include "sdl_draw.h"

void seed_filling(SDL_Surface *screen,int i,int j,Uint32 color_fill,Uint32 boundary_color)
{
 //没有成功
 Uint32 c = get_pixel(screen,i,j);//得到i,j点的颜色
 if((c != boundary_color)&&(c != color_fill))
 {
  //四邻法填充
  put_pixel(screen,i,j,color_fill);
 /* seed_filling(screen,i+1,j,color_fill,boundary_color);
  seed_filling(screen,i-1,j,color_fill,boundary_color);
  seed_filling(screen,i,j+1,color_fill,boundary_color);
  seed_filling(screen,i,j-1,color_fill,boundary_color);*/
 }
}
void plot_circle_points(SDL_Surface *screen,int xc,int yc,int x,int y,Uint32 c)
{
 put_pixel(screen,xc+x,yc+y,c);
 put_pixel(screen,xc-x,yc+y,c);
 put_pixel(screen,xc+x,yc-y,c);
 put_pixel(screen,xc-x,yc-y,c);
 put_pixel(screen,xc+y,yc+x,c);
 put_pixel(screen,xc-y,yc+x,c);
 put_pixel(screen,xc+y,yc-x,c);
 put_pixel(screen,xc-y,yc-x,c);
}
void bresenham_circle(SDL_Surface *screen,int xc,int yc,int radius,Uint32 c)
{
 int x,y,p;
 x = 0;
 y = radius;
 p = 3-2*radius;
 while(x<y)
 {
  plot_circle_points(screen,xc,yc,x,y,c);
  if(p<0)
   p=p+4*x+6;
  else
  {
   p=p+4*(x-y)+10;
   y-=1;
  }
  x+=1;
 }
 if(x == y)
  plot_circle_points(screen,xc,yc,x,y,c);
}
void bresenham_line(SDL_Surface *screen,int x1,int y1,int x2,int y2,Uint32 c)
{
 int dx,dy,x,y,p,const1,const2,inc,tmp;
 dx = x2-x1;
 dy = y2-y1;
 if(dx*dy >= 0)
  inc = 1;
 else
  inc = -1;
 if(abs(dx)>abs(dy))
 {
  if(dx<0)
  {
   tmp = x1;
   x1 = x2;
   x2 = tmp;
   tmp = y1;
   y1 = y2;
   y2 = tmp;
   dx = -dx;
   dy = -dy;
  }
  p = 2*dy-dx;
  const1 = 2*dy;
  const2 = 2*(dy-dx);
  x = x1;
  y = y1;
  put_pixel(screen,x,y,c);
  while(x<x2)
  {
   x++;
   if(p<0)
    p += const1;
   else
   {
    y += inc;
    p += const2;
   }
   put_pixel(screen,x,y,c);
  }
 }
 else
 {
  if(dy<0)
  {
   tmp = x1;
   x1 = x2;
   x2 = tmp;
   tmp = y1;
   y1 = y2;
   y2 = tmp;
   dx = -dx;
   dy = -dy;
  }
  p = 2*dy-dx;
  const1 = 2*dy;
  const2 = 2*(dy-dx);
  x = x1;
  y = y1;
  put_pixel(screen,x,y,c);
  while(y<y2)
  {
   y++;
   if(p<0)
    p += const1;
   else
   {
    x+=inc;
    p+=const2;
   }
   put_pixel(screen,x,y,c);
  }
 }
}
void dda_line(SDL_Surface *screen,int xa,int ya,int xb,int yb,Uint32 c)
{
 float delta_x,delta_y,x,y;
 int dx,dy,steps,k;
 dx = xb-xa;
 dy = yb-ya;

 if(abs(dx)>abs(dy))
   steps = abs(dx);
 else
   steps = abs(dy);

 delta_x = (float)dx/(float)steps;
 delta_y = (float)dy/(float)steps;
 x = xa;
 y = ya;
 put_pixel(screen,x,y,c);
 for(k = 0;k<steps;k++)
 {
  x+=delta_x;
  y+=delta_y;
  put_pixel(screen,x,y,c);
 }
 return ;
}
/*
 * Return the pixel value at (x, y)
 * NOTE: The surface must be locked before calling this!
 */
Uint32 get_pixel(SDL_Surface *surface, int x, int y)
{
    int bpp = surface->format->BytesPerPixel;
    /* Here p is the address to the pixel we want to retrieve */
    Uint8 *p = (Uint8 *)surface->pixels + y * surface->pitch + x * bpp;

    switch(bpp) {
    case 1:
        return *p;

    case 2:
        return *(Uint16 *)p;

    case 3:
        if(SDL_BYTEORDER == SDL_BIG_ENDIAN)
            return p[0] << 16 | p[1] << 8 | p[2];
        else
            return p[0] | p[1] << 8 | p[2] << 16;

    case 4:
        return *(Uint32 *)p;

    default:
        return 0;       /* shouldn't happen, but avoids warnings */
    }
}

/*
 * Set the pixel at (x, y) to the given value
 * NOTE: The surface must be locked before calling this!
 */
void put_pixel(SDL_Surface *surface, int x, int y, Uint32 pixel)
{
    int bpp = surface->format->BytesPerPixel;
    /* Here p is the address to the pixel we want to set */
    Uint8 *p = (Uint8 *)surface->pixels + y * surface->pitch + x * bpp;

    switch(bpp) {
    case 1:
        *p = pixel;
        break;

    case 2:
        *(Uint16 *)p = pixel;
        break;

    case 3:
        if(SDL_BYTEORDER == SDL_BIG_ENDIAN) {
            p[0] = (pixel >> 16) & 0xff;
            p[1] = (pixel >> 8) & 0xff;
            p[2] = pixel & 0xff;
        } else {
            p[0] = pixel & 0xff;
            p[1] = (pixel >> 8) & 0xff;
            p[2] = (pixel >> 16) & 0xff;
        }
        break;

    case 4:
        *(Uint32 *)p = pixel;
        break;
    }
}

void draw_point(SDL_Surface *screen,int x,int y,Uint32 color)
{
    if ( SDL_MUSTLOCK(screen) ) {
        if ( SDL_LockSurface(screen) < 0 ) {
            fprintf(stderr, "Can't lock screen: %s\n", SDL_GetError());
            return;
        }
    }

 put_pixel(screen, x, y, color);

    if ( SDL_MUSTLOCK(screen) ) {
        SDL_UnlockSurface(screen);
    }
    /* Update just the part of the display that we've changed */
    SDL_UpdateRect(screen, x, y, 1, 1);

    return;

}

