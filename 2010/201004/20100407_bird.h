/**
 *  @file 20100407_bird.h
 *
 *  @date 2012-2-20
 *  @author Bonly
 */

#ifndef _BIRD_H_
#define _BIRD_H_
#include <cmath>
void process();
void display();
void disp_background();
void mouseEvn(SDL_Event evn);
void playground();
void bird();

Uint32 red = 0;
Uint32 blue = 0;
b2Body* getWorldBody(const char*);

template<typename C, typename T>
struct OBJ : public C
{
    T attr;
    void  operator=(const C& b)
    {
      memcpy(this, (const void*) &b, sizeof(b));
      return;
    }
    void  operator=(const T& b)
    {
      memcpy(this + sizeof(C), (const void*) &b, sizeof(b));
      return;
    }
};

template<typename KEY>
int  InButtonPic(int x, int y, int Xzero, int Yzero, KEY btn, int deviation = 0)
{
  if ((x + Xzero) >= (btn.x - deviation)
      && (x + Xzero) <= (btn.x + btn.w + deviation)
      && (y + Yzero) >= (btn.y - deviation)
      && (y + Yzero) <= (btn.y + btn.h + deviation))
    return 1;
  else
    return 0;
}

template<typename KEY>
int  InButtonPic(int x, int y, int Xzero, int Yzero, KEY* btn, int deviation = 0)
{
  if ((x + Xzero) >= (btn->x - deviation)
      && (x + Xzero) <= (btn->x + btn->w + deviation)
      && (y + Yzero) >= (btn->y - deviation)
      && (y + Yzero) <= (btn->y + btn->h + deviation))
    return 1;
  else
    return 0;

}


class myDraw : public b2Draw
{
  public:
  void DrawPolygon(const b2Vec2* vertices, int32 vertexCount, const b2Color& color);

  void DrawSolidPolygon(const b2Vec2* vertices, int32 vertexCount, const b2Color& color);

  void DrawCircle(const b2Vec2& center, float32 radius, const b2Color& color);

  void DrawSolidCircle(const b2Vec2& center, float32 radius, const b2Vec2& axis, const b2Color& color);

  void DrawSegment(const b2Vec2& p1, const b2Vec2& p2, const b2Color& color);

  void DrawTransform(const b2Transform& xf);
};

void dot(SDL_Surface *surface, int x, int y, Uint32 pixel)
{
  Uint8 *p = (Uint8 *)surface->pixels + y * surface->pitch + x * surface->format->BytesPerPixel;
  *(Uint16 *)p = pixel;
}
void line(SDL_Surface *surface, int xa, int ya, int xb, int yb, Uint32 pixel)
{
  Uint8 *p;
  Uint16 k, tmax = abs(xb - xa) > abs(yb - ya) ? abs(xb - xa) : abs(yb - ya);
  float tx = xa,ty = ya,dx = 1.0 * (xb - xa) / tmax,dy = 1.0 * (yb - ya) / tmax;
  for (k = 0; k < tmax; tx += dx, ty += dy, k++)
  {
    p = (Uint8 *) surface->pixels + ((int) ty) * surface->pitch
        + ((int) tx) * surface->format->BytesPerPixel;
    *(Uint16 *) p = pixel;
  }
}
#endif /* 20100407_BIRD_H_ */
