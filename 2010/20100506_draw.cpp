#include "draw.h"

static unsigned char gamma_table[NUM_LEVELS];
static double Draw_gamma = 2.35;
static int float_compare(double x, double y, double m);/*m is for ellpise*/
static void lock(SDL_Surface *surface);
static void unlock(SDL_Surface *surface);


void Draw_Pixel(SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 color)
{
    int bpp = surface->format->BytesPerPixel;
    /* Here p is the address to the color we want to set */
    Uint8 *p = (Uint8 *)surface->pixels + y * surface->pitch + x * bpp;

/*Try to optimize "case" when use different cpp*/
    switch(bpp) {
    case 1:
        *p = color;
        break;

    case 2:
        *(Uint16 *)p = color;
        break;

    case 3:
        if(SDL_BYTEORDER == SDL_BIG_ENDIAN) {
            p[0] = (color >> 16) & 0xff;
            p[1] = (color >> 8) & 0xff;
            p[2] = color & 0xff;
        } else {
            p[0] = color & 0xff;
            p[1] = (color >> 8) & 0xff;
            p[2] = (color >> 16) & 0xff;
        }
        break;

    case 4:
        *(Uint32 *)p = color;
        break;
    }
}

static void Draw_Line_In(SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 fgc,Uint32 bgc, int nlevels, int nbits)
{
	int i;
	int tmp;
	int dx;
	int dy;
	int xstep;
	int min_x;
	int min_y;
	int w;
	int h;

	Uint32 intshift, erracc,erradj;
	Uint32 erracctmp, wgt, wgtcompmask;
	Uint32 colors[nlevels];

	SDL_Rect rect;
	SDL_Color fg;
	SDL_Color bg;
	if (y2 > y3) {
		tmp = y2; y2 = y3; y3 = tmp;
		tmp = x2; x2 = x3; x3 = tmp;
	}
	
	dx = x3 - x2;
	dy = y3 - y2;

	xstep = (dx >= 0) ? 1 : -1;
	dx = (dx >= 0) ? dx : -dx;

	min_x = (x2 > x3) ? x3 : x2;
	min_y = y2;
	w = dx;
	h = dy;


#ifdef DEBUG
//	printf("x2=%d,y2=%d,x3=%d,y3=%d,dx=%d,dy=%d,xstep=%d",x2,y2,x3,y3,dx,dy,xstep);
#endif
	for (i=0; i<NUM_LEVELS; i++)
		gamma_table[i] = (int) (NUM_LEVELS-1)*pow((double)i/((double)\
					(NUM_LEVELS-1)), 1.0/Draw_gamma);

	SDL_GetRGB(fgc,surface->format,&fg.r,&fg.g,&fg.b);
	SDL_GetRGB(bgc,surface->format,&bg.r,&bg.g,&bg.b);

	for (i=0; i<nlevels; i++) {
		Uint8 r, g, b;

		r = gamma_table[fg.r - (i*(fg.r - bg.r))/(nlevels-1)];
		g = gamma_table[fg.g - (i*(fg.g - bg.g))/(nlevels-1)];
		b = gamma_table[fg.b - (i*(fg.b - bg.b))/(nlevels-1)];
		colors[i] = SDL_MapRGB(surface->format, r, g, b);
	}


	if (x3 > (Uint32)surface->w || y3 > (Uint32)surface->h) {
		printf("Out of screen!\n");
		exit(1);
	}

	if ((dx == 0)) {
		rect.x = x2;
		rect.y = y2;
		rect.w = 1;
		rect.h = dy;
		SDL_FillRect(surface,&rect,fgc);
		SDL_UpdateRect(surface,rect.x,rect.y,rect.w,rect.h);
		return;
	}
	if ((dy == 0)) {
		rect.x = (x2 < x3) ? x2 : x3;
		rect.y = y2;
		rect.w = dx;
		rect.h = 1;
		SDL_FillRect(surface,&rect,fgc);
		SDL_UpdateRect(surface,rect.x,rect.y,rect.w,rect.h);
		return;
	}
	if ((dx == dy)) {
			for (;dy!=0;dy--) {
			x2 += xstep;
			y2++;
			Draw_Pixel(surface,x2,y2,fgc);
		}
			SDL_UpdateRect(surface,min_x,min_y,dx,dx);
		return;
	}
	/* line is not horizontal, vertical, or diagonal */
	erracc = 0;			/* err. acc. is initially zero */
	/* # of bits by which to shift erracc to get intensity level */
	intshift = 32 - nbits;
	/* mask used to flip all bits in an intensity weighting */
	wgtcompmask = nlevels - 1;
	/* x-major or y-major? */
	if (dy > dx) {
		/* y-major.  Calculate 16-bit fixed point fractional part of a pixel that
		X advances every time Y advances 1 pixel, truncating the result so that
		we won't overrun the endpoint along the X axis */
		erradj = ((Uint64)dx << 32) / (Uint64)dy;
		/* draw all pixels other than the first and last */
		while (--dy) {
			erracctmp = erracc;
			erracc += erradj;
			if (erracc <= erracctmp) {
			/* rollover in error accumulator, x coord advances */
				x2 += xstep;
			}
			y2++;			/* y-major so always advance Y */
			/* the nbits most significant bits of erracc give us the intensity
			 weighting for this pixel, and the complement of the weighting for
			 the paired pixel. */
			wgt = erracc >> intshift;
			Draw_Pixel(surface, x2, y2, colors[wgt]);
			Draw_Pixel(surface, x2+xstep, y2, colors[wgt^wgtcompmask]);
		}
		/* draw the final pixel, which is always exactly intersected by the line
		and so needs no weighting */
		Draw_Pixel(surface, x3, y3,fgc);
		SDL_UpdateRect(surface,min_x,min_y,w,h);
		return;
	}
	/* x-major line.  Calculate 16-bit fixed-point fractional part of a pixel
	that Y advances each time X advances 1 pixel, truncating the result so
	that we won't overrun the endpoint along the X axis. */
	erradj = ((Uint64)dy << 32) / (Uint64)dx;
	/* draw all pixels other than the first and last */
	while (--dx) {
		erracctmp = erracc;
		erracc += erradj;
		if (erracc <= erracctmp) {
		/* accumulator turned over, advance y */
			y2++;
		}
		x2 += xstep;			/* x-major so always advance X */
		/* the nbits most significant bits of erracc give us the intensity
		weighting for this pixel, and the complement of the weighting for
		the paired pixel. */
		wgt = erracc >> intshift;
		Draw_Pixel(surface, x2, y2, colors[wgt]);
		Draw_Pixel(surface, x2, y2+1, colors[wgt^wgtcompmask]);
	}
	/* draw final pixel, always exactly intersected by the line and doesn't
	need to be weighted. */
	Draw_Pixel(surface, x3, y3, fgc);
	SDL_UpdateRect(surface,min_x,min_y,w,h);
}

void Draw_Line(SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 color)
{
	int nlevels = NUM_LEVELS;
	int nbits = 8;
	int bgc = 0;
	int fgc = color;
	lock(surface);
	Draw_Line_In(surface,x2,y2,x3,y3,fgc,bgc,nlevels,nbits);
	unlock(surface);

}
void Draw_Rect (SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 color)
{
	Draw_Line(surface,x2,y2,x3,y2,color);
	Draw_Line(surface,x2,y2,x2,y3,color);
	Draw_Line(surface,x3,y3,x3,y2,color);
	Draw_Line(surface,x3,y3,x2,y3,color);

}

void Draw_Rect_Fill (SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 color, Uint32 bcolor)
{
	SDL_Rect rect;
	int w;
	int h;

	w = (x3 > x2) ? x3-x2 : -(x3-x2);
	h = (y3 > y2) ? y3-y2 : -(y3-y2);

	rect.x = (x3 > x2) ? x2 : x3;
	rect.y = (y3 > y2) ? y2 : y3;
	rect.w = w;
	rect.h = h;

	lock(surface);
	
	if (color == bcolor) {			/*if no border*/
		SDL_FillRect(surface,&rect,color);
		SDL_UpdateRect(surface,rect.x,rect.y,w,h);
	}
	else {
		rect.x += 1;
		rect.y += 1;
		rect.w -= 1;
		rect.h -= 1;
//	printf("x2,y2,w,h %d,%d,%d,%d\n",x2,y2,w,h);
		SDL_FillRect(surface,&rect,color);
		Draw_Rect(surface,x2,y2,x3,y3,bcolor);
		SDL_UpdateRect(surface,rect.x,rect.y,w,h);
	}
	unlock(surface);
}

void Draw_Circle (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 r, Uint32 color)
{
	Uint32 i;
	Uint32 j;

	lock(surface);
	for (i=x-r;i<=(x+r);i++) {
		for (j=y-r;j<=(y+r);j++){
			if (float_compare(p2p(i,j,x,y),r,0)) {	/*if point to circle's 
							 *center is as far as r
							 */
				Draw_Pixel(surface,i,j,color);
			}
		}
	}
	SDL_UpdateRect(surface,x-r,y-r,2*r+1,2*r+1);
	unlock(surface);
}

void Draw_Circle_Fill (SDL_Surface *surface, Uint32 x, Uint32 y, Uint32 r, Uint32 color,Uint32 bcolor)
{
	Uint32 i;
	Uint32 j;
	lock(surface);
	if (color == bcolor) {
		for (i=x-r;i<=(x+r);i++) {
			for (j=y-r;j<=(y+r);j++){
				if (p2p(i,j,x,y) <= r) {		
					Draw_Pixel(surface,i,j,color);
				}
			}
		}

	}

	else {
		for (i=x-r+1;i<=(x+r-1);i++) {
			for (j=y-r+1;j<=(y+r-1);j++){
				if (p2p(i,j,x,y) <= r) {		
					Draw_Pixel(surface,i,j,color);
				}
			}
		}
		Draw_Circle(surface,x,y,r,bcolor);

	}
	SDL_UpdateRect(surface,x-r,y-r,2*r+1,2*r+1);
	unlock(surface);

}


void Draw_Ellipse (SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 r, Uint32 color)
{
	Uint32 i;
	Uint32 j;
	double c;
	Uint32 min_x;
	Uint32 min_y;

	c = p2p(x2,y2,x3,y3);			/*椭圆两焦点间距离*/
	if (r < c) {
		printf("Invalid r!\n");
		exit(1);
	}
	lock(surface);

	min_x = (Uint32)(((x3 + x2)/2) - r);
	min_y = (Uint32)(((y3 + y2)/2) - r);
					/*以椭圆中点圆心做正方形 边长为r 						 *在这个正方形内寻找匹配的点
					 */

	//printf("c=%g min_x=%d min_y=%d r=%d",c,min_x,min_y,r);
	for (i=min_x;i<=min_x+2*r;i++) {
		for (j=min_y;j<=min_y+2*r;j++) {
			if (float_compare((p2p(i,j,x2,y2) + 
					   p2p(i,j,x3,y3)),r,0.4)) {
				Draw_Pixel(surface,i,j,color);
		}
		}
	}
	SDL_UpdateRect(surface,min_x,min_y,min_x+2*r,min_y+2*r);
	unlock(surface);
}

void Draw_Ellipse_Fill (SDL_Surface *surface, Uint32 x2, Uint32 y2, Uint32 x3, Uint32 y3, Uint32 r, Uint32 color, Uint32 bcolor)
{
	Uint32 i;
	Uint32 j;
	double c;
	Uint32 min_x;
	Uint32 min_y;

	c = p2p(x2,y2,x3,y3);			/*椭圆两焦点间距离*/
	if (r < c) {
		printf("Invalid r!\n");
		exit(1);
	}
	lock(surface);

	min_x = (Uint32)(((x3 + x2)/2) - r);
	min_y = (Uint32)(((y3 + y2)/2) - r);
					/*以椭圆中点圆心做正方形 边长为r 						 *在这个正方形内寻找匹配的点
					 */

	//printf("c=%g min_x=%d min_y=%d r=%d",c,min_x,min_y,r);
	if (color == bcolor) {
		for (i=min_x;i<=min_x+2*r;i++) {
			for (j=min_y;j<=min_y+2*r;j++) {
				if ((p2p(i,j,x2,y2) + 
						   p2p(i,j,x3,y3)) <= r) {
					Draw_Pixel(surface,i,j,color);
			}
			}
		}
	}
	else {
		Draw_Ellipse(surface,x2,y2,x3,y3,r,bcolor);
		for (i=min_x;i<=min_x+2*r;i++) {
			for (j=min_y;j<=min_y+2*r;j++) {
				if ((p2p(i,j,x2,y2) + 
						   p2p(i,j,x3,y3)) < r-1) {
					Draw_Pixel(surface,i,j,color);
				}		/*如果r为1则边缘效果不太明显*/
			}
		}

	}
	SDL_UpdateRect(surface,min_x,min_y,min_x+2*r,min_y+2*r);
	unlock(surface);
}

/*calculate distants between two points*/
static int float_compare(double x, double y, double m)
{
	return ((x - y) >= -(EPSILON+m) && (x -y ) <= (EPSILON+m)) ? 1 : 0;
}
double p2p(Uint32 i, Uint32 j, Uint32 x, Uint32 y) 
{
	
	return sqrt(pow(abs(i-x),2) + pow(abs(j-y),2));

}

static void lock(SDL_Surface *surface)
{
if (SDL_MUSTLOCK(surface) < 0 ) {
	if (SDL_LockSurface(surface) < 0 ) {
		printf("Can't lock screen:%s\n",SDL_GetError());
		exit(1);
	}
	}
}

static void unlock(SDL_Surface *surface)
{
	if (SDL_MUSTLOCK(surface)) {
		SDL_UnlockSurface(surface);
	}

}
