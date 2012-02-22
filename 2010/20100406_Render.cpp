/**
 *  @file Render.cpp
 *
 *  @date 2012-2-20
 *  @Author: Bonly
 */

#include "20100406_Render.h"
#include "20100403_sdl_box.h"

extern Task gtask[];
Render::Render()
{
  // @TODO Auto-generated constructor stub
  
}

Render::~Render()
{
  // @TODO Auto-generated destructor stub
}

/// Draw a closed polygon provided in CCW order.
void Render::DrawPolygon(const b2Vec2* vertices, int32 vertexCount, const b2Color& color)
{

}

/// Draw a solid closed polygon provided in CCW order.
void Render::DrawSolidPolygon(const b2Vec2* vertices, int32 vertexCount, const b2Color& color)
{

}

/// Draw a circle.
void Render::DrawCircle(const b2Vec2& center, float32 radius, const b2Color& color)
{

}

/// Draw a solid circle.
void Render::DrawSolidCircle(const b2Vec2& center, float32 radius, const b2Vec2& axis, const b2Color& color)
{

}

/// Draw a line segment.
void Render::DrawSegment(const b2Vec2& p1, const b2Vec2& p2, const b2Color& color)
{

}

/// Draw a transform. Choose your own length scale.
/// @param xf a transform.
void Render::DrawTransform(const b2Transform& xf)
{
  OBJ<Bird,SDL_Rect>* bird = dynamic_cast<OBJ<Bird, SDL_Rect>* >(getGoods("Ð¡ºìÄñ"));
  bird->Display();

}

