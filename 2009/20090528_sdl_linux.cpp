#include <iostream>
#include <SDL.h>

int
main (int argc, char* argv[])
{
  try
  {
    if(SDL_Init(SDL_INIT_EVERYTHING)==-1)
       throw "could not initialze SDL!";
  }
  catch(const char*s)
  {
    std::cerr << s << std::endl;
    return -1;
  }
  std::cout << "SDL initialized.\n";
  SDL_Quit();
  return 0;
}

