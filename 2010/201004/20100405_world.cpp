/**
 *  @file 20100405_world.cpp
 *
 *  @date 2012-2-20
 *  @Author: Bonly
 */
#include "20100405_world.h"

BirdWorld::BirdWorld()
{
  b2Vec2 gravity;  ///��������
  gravity.Set(0.0f, -10.0f);
  _world = new b2World(gravity); ///��������
  _world->SetContactListener(this); ///���ù�ϵ������
  _world->SetDebugDraw(&_Draw); ///���û�ͼ������

  b2Body* ground = NULL;
  {
    b2BodyDef bd; ///����������õĸ���
    ground = _world->CreateBody(&bd); ///��������

    b2EdgeShape shape; ///����Ƕ���״��
    shape.Set(b2Vec2(-40.0f, 0.0f), b2Vec2(40.0f, 0.0f));
    ground->CreateFixture(&shape, 0.0f); ///�����ϴ����Ƕ���״��
  }

  ///����Բ����
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
