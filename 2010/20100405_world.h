/**
 *  @file 20100405_world.h
 *
 *  @date 2012-2-20
 *  @Author: Bonly
 */

#ifndef _WORLD_H_
#define _WORLD_H_
#include <Box2D/Box2D.h>
#include "20100406_Render.h"

class BirdWorld : public b2ContactListener
{
  public:
    BirdWorld();

  public:
    b2World* _world;
    Render _Draw;
    b2Body* bird;

    void Step();

};

#endif /* 20100405_WORLD_H_ */
