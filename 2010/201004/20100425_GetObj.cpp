/**
 *  @file 20100425_GetObj.cpp
 *
 *  @date 2012-3-13
 *  @Author: Bonly
 */

#include "20100425_GetObj.h"

#define IMGSEG
#define RES(O,F) OBJ<Obj,void*> V_##O;
#include "20100425_Resource.inc"
#undef RES

#define RES(O,F) (OBJ<Obj,void*>*)&V_##O,
OBJ<Obj, void*>* List[] = {
#include "20100425_Resource.inc"
            };
#undef RES

#define RES(X,F) \
 void load_##X(void*) { \
 if(V_##X.attr!=0) SafeDelete(V_##X.attr); \
 V_##X.attr =  SafeNew JImage(#F, TRUE); \
}
#include "20100425_Resource.inc"
#undef RES

#undef IMGSEG

void DELIMG(int idx)
{
    if (idx >= RES_MAX)
       return;
    SafeDelete(List[idx]->attr);
}
ResPool* ResPool::resp = 0;
ResPool::ResPool()
{
    fixRelation();
}

ResPool::~ResPool()
{
}

void ResPool::destory()
{
    if (resp)
    {
        for (int i = 0; i < RES_MAX; ++i)
        {
            SafeDelete(List[i]->attr);
        }SafeDelete(resp);
    }
}

void ResPool::fixRelation()
{
    int i = 0;
#define IMGSEG
#define RES(X,F) \
  List[i]->ID = i; \
  List[i]->create = load_##X; \
  ++i;
#include "20100425_Resource.inc"
#undef RES
#undef IMGSEG
}
ResPool& ResPool::instance()
{
    if (0 == resp)
    {
        resp = new ResPool;

    }
    return *resp;
}

JImage* GETIMG(int idx)
{
    if (idx > RES_MAX)
        return 0;
    JImage* res = 0;
    if ((res = (JImage*) RESP.getObj<OBJ<Obj, void*> >(idx)->attr) == 0)
    {
        List[idx]->create(0);
    }
    res = (JImage*) RESP.getObj<OBJ<Obj, void*> >(idx)->attr;
    return res;
}
int main()
{
    JImage* a = GETIMG(ID_floor_1_20_1);
    JImage* b = GETIMG(ID_floor_1_20_1);
    JImage* c = GETIMG(ID_floor_1_1_1);
    DELIMG(ID_floor_1_1_1);
    //ResPool::destory();
    int* k = new int;
    *k = 10;

    return 0;
}
