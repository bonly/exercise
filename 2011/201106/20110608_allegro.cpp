#include <allegro.h>

int main(){
  allegro_init();
  install_keyboard();
  set_gfx_mode(GFX_AUTODETECT,  640, 480, 0, 0);
  readkey();
  return 0;
}

END_OF_MAIN();
//g++ 20110608_allegro.cpp  -l alleg
