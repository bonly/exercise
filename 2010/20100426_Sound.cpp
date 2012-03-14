#include "Sound.h"
#include "jport.h"

extern int sxm_dbgprintf(const char *fmt, ...);

DMusic* dmusic = 0;

DMusic::DMusic(void)
:sndf(0),sndb(0)
{
  Handle = JSoundPlayerDoubleMusicCreate();
}

DMusic::~DMusic(void)
{
  J_SoundPlayer_StopBackgrdMusic_Mid(Handle);
  J_SoundPlayer_StopFrontMusic_Wav(Handle);
  JSoundPlayerDoubleMusicDestory(Handle);
}

DMusic& DMusic::instance()
{
  if(0 == dmusic)
    dmusic = SafeNew DMusic;
  return *dmusic;
}

void DMusic::destory()
{
  for (int i=0; i<SND_MAX; ++i)
  {
    FREE(SND_List[i]->buffer);
  }
  SafeDelete(dmusic);
}

bool DMusic::play(int id)
{
  sxm_dbgprintf("in play: %d type %d\n",id, SND_List[id]->type);
  switch(SND_List[id]->type)
  {
    case Sound::FRONT:
		  if(sndf)
		     J_SoundPlayer_StopFrontMusic_Wav(Handle);
      J_SoundPlayer_PlayFrontMusic_Wav(Handle, SND_List[id]->buffer, SND_List[id]->size);
      break;
    case Sound::BACK:
      if(sndb)
        J_SoundPlayer_StopBackgrdMusic_Mid(Handle);
      J_SoundPlayer_PlayBackgrdMusic_Mid(Handle, SND_List[id]->buffer, SND_List[id]->size);
      break;
  }
  return true;
}

#define SNDSEG

bool DMusic::PreLoad()
{
#define RES(K,F,T) if(!load_##K(0)) return false;
#include "Resource.inc"
#undef RES
  return true;
}

#define RES(K,F,T) Sound V_##K;
#include "Resource.inc"
#undef RES

#define RES(K,F,T) (Sound*)&V_##K,
Sound* SND_List[]={
  #include "Resource.inc"
};
#undef RES

#define RES(K,F,T) \
bool load_##K(void*) \
{\
  if (V_##K.buffer) return true; \
  IFile* pIFileRead = IFILEMGR_OpenFile(GETAPPBASIC()->pIFileMgr, #F , _OFM_READWRITE);	\
  if (!pIFileRead) return false; \
  FileInfo f_info; \
  if(IFILE_GetInfo((IFile *)pIFileRead, &f_info)==SUCCESS) \
  V_##K.size = f_info.dwSize; \
  else return false; \
  V_##K.buffer = (char*)MALLOC(V_##K.size); \
  int rc = IFILE_Read(pIFileRead, V_##K.buffer, V_##K.size);  \
  IFILE_Release(pIFileRead); \
  V_##K.type = T; \
  sxm_dbgprintf("T is: %d %d\n",T,V_##K.type); \
  V_##K.ID = ID_##K; \
  return true;  \
}
#include "Resource.inc"
#undef RES

#undef SNDSEG
