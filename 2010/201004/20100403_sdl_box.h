/**
 *  @file 20100403_sdl_box.h
 *
 *  @date 2012-2-20
 *  @Author: Bonly
 */

#ifndef _SDL_BOX_H_
#define _SDL_BOX_H_
#include <SDL.h>
//#include <SDL_image.h>
#include <iostream>
class b2Body;
using namespace std;

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

struct Goods
{
    virtual int  Init(void*)=0;
    virtual int  Display()=0;
    virtual ~Goods()
    {
    };
};

typedef Goods* CreateFcn(void*);

struct Task
{
    const char *name; ///ÈÎÎñÃû
    CreateFcn *create;
    Goods *goods;
};

class Background : public Goods
{
    SDL_Surface *pDc;
    SDL_Surface* bmp;
  public:
    static Goods* Create(void *task);

    virtual int  Init(void* p)
    {
      pDc = (SDL_Surface*) p;

      try
      {
        bmp = SDL_LoadBMP("res/bg.bmp");
        if (bmp == 0)
          throw SDL_GetError();
      }
      catch (const char* s)
      {
        cerr << "load bmp failed: " << s << endl;
        return -1;
      }
      return 0;
    }
    virtual int Display();
};

class Bird : public Goods
{
  public:
    SDL_Surface *pDc;
    SDL_Surface* bmp;
    b2Body* wbird;
  public:
    static Goods* Create(void *task);

    virtual int Init(void* p)
    {
      pDc = (SDL_Surface*) p;

      try
      {
        bmp = SDL_LoadBMP("res/bird1.bmp");
        //bmp = IMG_Load("res/bird1.png");
        if (bmp == 0)
          throw SDL_GetError();
      }
      catch (const char* s)
      {
        cerr << "load bmp failed: " << s << endl;
        return -1;
      }
      return 0;
    }
    virtual int Display();
};

Goods* getGoods(const char* name);



#endif /* 20100403_SDL_BOX_H_ */
