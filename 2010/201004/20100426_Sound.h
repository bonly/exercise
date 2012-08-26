#ifndef __SOUND_H__
#define __SOUND_H__

struct Sound
{
  enum{MP3,FRONT,BACK};
  int ID;
  int type;
  int size;
  char* buffer;
};

struct DMusic
{
  DMusic(void);
  ~DMusic(void);
  bool PreLoad();
  bool play(int id);
  void* Handle;
  Sound* sndf;
  Sound* sndb;
  static DMusic& instance();
  void destory();
};
#define MUSIC DMusic::instance()
#define PLAY(X) DMusic::instance().play(X)

#define SNDSEG

#define RES(K,F,T) extern Sound V_##K;
#include "Resource.inc"
#undef RES

extern Sound* SND_List[];

#define RES(K,F,T) ID_##K,
enum SND_RES{
  #include "Resource.inc"
  SND_MAX
};
#undef RES
 
#define RES(K,F,T) bool load_##K(void*);
#include "Resource.inc"
#undef RES

#undef SNDSEG

#endif
