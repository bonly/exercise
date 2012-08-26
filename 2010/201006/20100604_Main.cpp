/**
 *  @file 201004.cpp
 *
 *  @date 2012-2-24
 *  @author Bonly
 */
#include <sdl/SDL.h>


SDL_Surface *pSc = 0; ///< 全局主屏幕画布指针

/**
 * 创建
 * @param p
 */
void onCreate(void *p)
{
  CONF->X = 320;
  CONF->Y = 480;
}
/**
 * 渲染
 * @param p
 */
void onPaint(void *p)
{
  TASKS->onPaint(p);
}
/**
 * 逻辑调用
 * @param p
 */
void onTimer(void *p)
{
    TASKS->run(p);
    onPaint(0);
}

void onDestroy(void *p)
{
    DELTASK();
    CONF->destory();
}
/**
 * 主函数
 * @param argc 参数个数
 * @param argv 参数数组
 * @return
 */
int main(int argc, char* argv[])
{
  try
  {
    if (SDL_Init(SDL_INIT_EVERYTHING) == -1)
      throw SDL_GetError();
  }
  catch (const char* s)
  {
    printf("%s", s);
    return -1;
  }
  atexit(SDL_Quit);

  onCreate(0);
  int flags = SDL_SWSURFACE;//|SDL_FULLSCREEN;
  pSc = SDL_SetVideoMode(CONF->X, CONF->Y, 32, flags);

  bool run = true;
  while (run)
  {
    SDL_Event evn;
    if (SDL_PollEvent(&evn) != 0)
    {
      switch (evn.type)
      {
        case SDL_QUIT:
          run = false;
          break;
      }
    }

    onTimer(0);
    SDL_Flip(pSc);
    SDL_Delay(50); ///等待一段时间,以便控制帧数

    if(TASKS->current_page == 0)
    {
        TASKS->next_page = 1;
        TASKS->operation = NEXT_PAGE;
    }
    else
    {
        TASKS->next_page = 0;
        TASKS->operation = NEXT_PAGE;
    }

  }
  onDestroy(0);

  return 0;
}

