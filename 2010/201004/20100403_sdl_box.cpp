//============================================================================
// Name        : sdl_box.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include "20100403_sdl_box.h"
#include "20100405_world.h"

void process();
void display();

namespace
{
  SDL_Surface* pSc = 0;
  int taskCount = 0;
  BirdWorld bw;
}

Task gtask[] ={
{ "背景", Background::Create },
{ "小红鸟", Bird::Create },
{ NULL, NULL }
};

Goods* Background::Create(void *task)
{
  Background* bg = new Background;
  reinterpret_cast<Task*>(task)->goods = bg;
  bg->Init(pSc);
  return bg;
}

int Background::Display()
{
  SDL_BlitSurface(bmp, 0, pSc, 0); ///移去另一个图层合并,SDL_Rect都为0表示左上角为0重合
  //SDL_UpdateRect(pSc, 0, 0, bmp->w, bmp->h);
  return 0;
}

int main(int argc, char* argv[])
{
  try
  {
    if (SDL_Init(SDL_INIT_EVERYTHING) == -1)
      throw SDL_GetError();
  }
  catch (const char* s)
  {
    cerr << s << endl;
    return -1;
  }
  cout << "SDL initialized.\n";
  atexit(SDL_Quit);

  pSc = SDL_SetVideoMode(240, 400, 32, SDL_SWSURFACE);

  while (gtask[taskCount].create != NULL)
  {
    ++taskCount;
  }

  for (int i = 0; i < taskCount; ++i)
  {
    gtask[i].create(&gtask[i]);
  }

  process();
  cout << "Game over" << endl;

  return 0;
}

Goods* Bird::Create(void *task)
{
  OBJ<Bird,SDL_Rect>* bird = new OBJ<Bird,SDL_Rect>;
  bird->attr.x = 0;
  bird->attr.y = 0;
  reinterpret_cast<Task*>(task)->goods = bird;
  bird->Init(pSc);
  bird->attr.h = bird->bmp->h;
  bird->attr.w = bird->bmp->w;
  bird->wbird = bw.bird;
  return bird;
}

int Bird::Display()
{
  b2Vec2 ps = wbird->GetPosition();
  SDL_Rect rec;
  rec.x = ps.x;
  rec.y = ps.y;
  //OBJ<Bird,SDL_Rect> *bird = dynamic_cast <OBJ<Bird,SDL_Rect>* >(this);
  SDL_BlitSurface(bmp, 0, pSc, &rec); ///移去另一个图层合并,SDL_Rect都为0表示左上角为0重合
  //SDL_UpdateRect(pSc, 0, 0, bmp->w, bmp->h);
  return 0;
}

Goods* getGoods(const char* name)
{
  for (int i=0; i<taskCount; ++i)
  {
    if(strcmp(name, gtask[i].name)==0)
      return gtask[i].goods;
  }
  return 0;
}

void mouseEvn(SDL_Event evn)
{
  if (evn.button.button == SDL_BUTTON_LEFT)
  {
      OBJ<Bird,SDL_Rect>* redbird = (OBJ<Bird,SDL_Rect>*)getGoods("小红鸟");
      if(1==InButtonPic(evn.button.x, evn.button.y, 0, 0, redbird->attr))
      {
         redbird->attr.x = evn.motion.x - (redbird->attr.w/2);
         redbird->attr.y = evn.motion.y - (redbird->attr.h/2);
      }

  }

}

void process()
{
  bool gameOver = false;
  while (!gameOver)
  {
    SDL_Event evn;

    while (SDL_PollEvent(&evn) != 0)
    {
      switch (evn.type)
      {
        case SDL_QUIT:
          gameOver = true;
          break;
        case SDL_KEYDOWN:
          if (evn.key.keysym.sym == SDLK_ESCAPE)
            gameOver = true;
          break;
        case SDL_MOUSEMOTION:
          mouseEvn(evn);
          break;
      }
    }
    display();
  }
  return;
}

void display()
{
  for (int i = 0; i < taskCount; ++i)
  {
    gtask[i].goods->Display();
  }
  bw.Step();
  if (0 != SDL_Flip(pSc)) ///刷新
  {
    clog << "flip faild: " << SDL_GetError();
  }
}

