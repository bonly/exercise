/**
 *  @file 20100425_GetObj.h
 *
 *  @date 2012-3-13
 *  @Author: Bonly
 */

#ifndef _GETOBJ_H_
#define _GETOBJ_H_

#define SafeDelete(X) if(X){delete X;X=0;}
#define SafeNew new
#define TRUE true

struct JImage
{
    JImage(const char* i, bool b){};
    int myint;
};

typedef void(*Tcreate)(void*);
struct Obj
{
    int  ID;
    Tcreate create;
};

template<typename C, typename T>
struct OBJ : public C
{
  T  attr;
};

//////////////////////////////////////////
extern OBJ<Obj,void*>* List[];
#define IMGSEG
enum RES{
  #define RES(O,F) ID_##O,
  #include "20100425_Resource.inc"
  #undef RES
  RES_MAX
};
//////////////////////////////////////////
class ResPool
{
  public:
    static ResPool* resp;

  public:
    void fixRelation();
    ResPool();
    virtual ~ResPool();

    template<typename T>
    T* getObj(int ID)
    {
      if (ID > RES_MAX)
        return 0;

      return (List[ID]);
    }

    static ResPool& instance();
    static void destory();
};

#define RESP ResPool::instance()
#define GETOBJ(X,Y) RESP.getObj<Y>(X)
JImage* GETIMG(int idx);
void DELIMG(int idx);
/////////////////////////////////////////////////////////////



#define RES(O,F) extern OBJ<Obj,void*> V_##O;
#include "20100425_Resource.inc"
#undef RES

#define RES(O,F) typedef OBJ<Obj,void*> T_##O;
#include "20100425_Resource.inc"
#undef RES

#define RES(X,F) \
 void load_##X(void*);
#include "20100425_Resource.inc"
#undef RES

#undef IMGSEG

#endif /* 20100425_GETOBJ_H_ */
