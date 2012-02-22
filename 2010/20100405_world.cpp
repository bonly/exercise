/**
 *  @file 20100405_world.cpp
 *
 *  @date 2012-2-20
 *  @Author: Bonly
 */
#include "20100405_world.h"

BirdWorld::BirdWorld()
{
  b2Vec2 gravity;  ///定义重力
  gravity.Set(0.0f, -10.0f);
  _world = new b2World(gravity); ///创建世界
  _world->SetContactListener(this); ///设置关系监听器
  _world->SetDebugDraw(&_Draw); ///设置绘图监听器

  b2Body* ground = NULL;
  {
    b2BodyDef bd; ///定义地面所用的刚体
    ground = _world->CreateBody(&bd); ///创建地面

    b2EdgeShape shape; ///定义角度形状物
    shape.Set(b2Vec2(-40.0f, 0.0f), b2Vec2(40.0f, 0.0f));
    ground->CreateFixture(&shape, 0.0f); ///地面上创建角度形状物
  }

  ///定义圆形物
  b2CircleShape shape;
  shape.m_radius = 0.5f;

  b2FixtureDef fd;
  fd.shape = &shape;
  fd.density = 1.0f;

  b2BodyDef bd;
  bd.type = b2_dynamicBody;
  bd.position.Set(-6.0f + 6.0f, 10.0f);
  bird = _world->CreateBody(&bd);
  bird->CreateFixture(&fd);
}

void BirdWorld::Step()
{
  _world->Step(1, 1, 1);
  //_world->DrawDebugData();
}
