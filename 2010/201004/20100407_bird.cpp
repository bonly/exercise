/**
 *  @file 20100407_bird.cpp
 *
 *  @date 2012-2-20
 *  @author Bonly
 */
#include <SDL.h>
#include <cstdlib>
#include <iostream>
#include <freeglut.h>
#include <Box2D/Box2D.h>
//#include <SDL_image.h>
#include "20100407_bird.h"
#include "draw.h"

using namespace std;
namespace
{
  SDL_Surface* pSc = 0;
  b2World *world;
  SDL_Surface *bg;
  SDL_Surface *redbird;
  SDL_Surface *pground;
  SDL_Surface *sling;
  myDraw *dr;
  b2Vec2 slingPoint;
  const float ptm = 30.0f;
}

struct UserData
{
    char name[20];
    bool mouseAble;
};

UserData ud[10];

int main(int argc, char* argv[])
{
  try
  {
    if (SDL_Init(SDL_INIT_EVERYTHING) == -1)
      throw SDL_GetError();
  }
  catch (const char* s)
  {
    cerr << s << endl;
    return -1;
  }
  cout << "SDL initialized.\n";
  atexit(SDL_Quit);

  int flags = SDL_SWSURFACE;//SDL_OPENGL;//|SDL_FULLSCREEN;
  pSc = SDL_SetVideoMode(400, 600, 32, flags);
  /*
  SDL_GL_SetAttribute(SDL_GL_DOUBLEBUFFER,1);
  glViewport(0,0,240,400);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  gluPerspective(60.0,240/400,1.0,1024.0);
  //*/

  b2Vec2 gravity;  ///定义重力
  gravity.Set(-10.0f/ptm, 0.0f);
  world = new b2World(gravity); ///创建世界
  world->SetAllowSleeping(true);

  bg = SDL_LoadBMP("res/bg.bmp");
  //bg = IMG_Load("res/bg.png");
  if(bg==0)
    cerr << SDL_GetError() << endl;
  {
    SDL_Surface* img = SDL_LoadBMP("res/bird1.bmp");
    redbird = SDL_DisplayFormat(img);
    SDL_FreeSurface(img);
    Uint32 colorkey = SDL_MapRGB(redbird->format, 0, 0, 0xFF);
    SDL_SetColorKey(redbird, SDL_RLEACCEL|SDL_SRCCOLORKEY, colorkey);
  }
  redbird = SDL_LoadBMP("res/bird1.bmp");
  sling = SDL_LoadBMP("res/sling.bmp");

  pground = SDL_LoadBMP("res/stick1.bmp");
  red = SDL_MapRGB(pSc->format, 0xff, 0x00, 0x00);
  blue = SDL_MapRGB(pSc->format, 0x00, 0x00, 0xff);
  dr = new myDraw;
  world->SetDebugDraw(dr);
  playground();
  bird();
  process();
  delete world;
  delete dr;
  return 0;
}

void process()
{
  bool gameOver = false;
  while (!gameOver)
  {
    SDL_Event evn;

    while (SDL_PollEvent(&evn) != 0)
    {
      switch (evn.type)
      {
        case SDL_QUIT:
          gameOver = true;
          break;
        case SDL_KEYDOWN:
          if (evn.key.keysym.sym == SDLK_ESCAPE)
            gameOver = true;
          break;
        case SDL_MOUSEMOTION:
        case SDL_MOUSEBUTTONDOWN:
        case SDL_MOUSEBUTTONUP:
          mouseEvn(evn);
          break;
      }
    }
    uint32 flags = 0;
    flags += b2Draw::e_shapeBit;
    //flags += b2Draw::e_aabbBit;
    dr->SetFlags(flags);
    world->Step(1.0f/20.0f,20,10);
    disp_background();
    world->DrawDebugData();

    //bresenham_line(pSc,0,0,70,89,128);
    //display();
    SDL_Flip(pSc);
    //SDL_GL_SwapBuffers();
  }
  return;
}

b2Body* getWorldBody(const char*)
{
  for(b2Body *currentBody = world->GetBodyList();
      currentBody;
      currentBody=currentBody->GetNext()
      )
  {
       UserData* tud;
       if((tud =(UserData*)currentBody->GetUserData()))
       {
         if(strcmp(tud->name, "redbird")==0)
           return currentBody;
       }
  }
  return 0;
}
void display()
{
  for(b2Body *currentBody = world->GetBodyList();
      currentBody;
      currentBody=currentBody->GetNext()
      )
  {
       UserData* tud;
       if((tud = (UserData*)currentBody->GetUserData()))
       {
         if(strcmp(tud->name,"redbird")==0)
         {
           int x = (currentBody->GetPosition()).x;
           int y = (currentBody->GetPosition()).y;
           SDL_Rect rec;
           rec.x = x; rec.y = y;
           rec.h = redbird->h;
           rec.w = redbird->w;
           SDL_BlitSurface(redbird,0,pSc,&rec);
         }
         if(strcmp(tud->name,"playground")==0)
         {
           int x = (currentBody->GetPosition()).x;
           int y = (currentBody->GetPosition()).y;
           SDL_Rect rec;
           rec.x = x; rec.y = y;
           rec.h = pground->h;
           rec.w = pground->w;
           SDL_BlitSurface(pground,0,pSc,&rec);
         }
       }
  }
  //SDL_Flip(pSc);
}

void mouseEvn(SDL_Event evn)
{
  static bool drag = false;
  if(evn.type == SDL_MOUSEMOTION )
  {
    //if (evn.button.button == SDL_BUTTON_LEFT)
    //if (SDL_GetMouseState(NULL,NULL)&SDL_BUTTON(1))
    if (SDL_GetMouseState(NULL,NULL)&&SDL_BUTTON_LMASK&&drag)
    {
        b2Body* redbirdBody = getWorldBody("redbird");
        if(redbirdBody==0)
          return;
        SDL_Rect rc;
        rc.x = (redbirdBody->GetPosition().x - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
        rc.y = (redbirdBody->GetPosition().y - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
        rc.h = redbird->h;
        rc.w = redbird->w;
        if(1==InButtonPic(evn.button.x, evn.button.y, 0, 0, rc))
        {
          /*
          b2Vec2 touchPos(evn.button.x, evn.button.y);
          b2Vec2 fromBodyToTouch = touchPos - redbirdBody->GetPosition();
          float distanceFromBodyToTouch = fromBodyToTouch.Normalize();//makes the vector length 1, returns the original length
          float power = 30.0;//... whatever ...; //maybe use the distance to alter the power?
          b2Vec2 forceToApply = power * fromBodyToTouch;
          redbirdBody->ApplyForce( forceToApply, redbirdBody->GetWorldCenter() );//or ApplyLinearImpulse, or SetLinearVelocity
          redbirdBody->SetAwake(true);
          redbirdBody->SetTransform(redbirdBody->GetPosition(),0);
          //*/
          //*
          b2Vec2 ps;
          ps.x = evn.motion.x/ptm ;//- (redbird->w/2);
          ps.y = evn.motion.y/ptm ;//- (redbird->h/2);
          redbirdBody->SetTransform(ps,0);
          //b2Vec2 f = redbirdBody->GetWorldVector(b2Vec2(0.0f, -200.0f));
          //b2Vec2 p = redbirdBody->GetWorldPoint(b2Vec2(0.0f, 2.0f));
          //redbirdBody->ApplyForce(f, p);
          //redbirdBody->ApplyForceToCenter(p);
          //redbirdBody->SetAwake(true);
          //*/
        }
    }
  }
  if ((evn.type == SDL_MOUSEBUTTONDOWN)&&SDL_BUTTON_LMASK)
  {
    SDL_Rect rc;
    b2Body* redbirdBody = getWorldBody("redbird");
    rc.x = (redbirdBody->GetPosition().x - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
    rc.y = (redbirdBody->GetPosition().y - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
    rc.h = redbird->h;
    rc.w = redbird->w;
    if(1==InButtonPic(evn.button.x, evn.button.y, 0, 0, rc))
    {
      redbirdBody->SetAwake(false);
      drag = true;
    }
  }
  if((evn.type == SDL_MOUSEBUTTONUP )&& SDL_BUTTON_LMASK)
  {
    SDL_Rect rc;
    b2Body* redbirdBody = getWorldBody("redbird");
    rc.x = (redbirdBody->GetPosition().x - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
    rc.y = (redbirdBody->GetPosition().y - redbirdBody->GetFixtureList()[0].GetShape()->m_radius)*ptm;
    rc.h = redbird->h;
    rc.w = redbird->w;
    if(1==InButtonPic(evn.button.x, evn.button.y, 0, 0, rc))
    {
      float32 distanceX= (float32)evn.button.x - slingPoint.x;
      float32 distanceY=(float32)evn.button.y - slingPoint.y;
      float32 distance=b2Sqrt((float32)(distanceX*distanceX+distanceY*distanceY));
      float32 birdAngle=b2Atan2((float32)distanceY, (float32)distanceX);
      b2Vec2 birdStrlen = b2Vec2(-distance*cos(birdAngle)/4.0f,-distance*sin(birdAngle)/4.0f);

      redbirdBody->SetActive(true);
      redbirdBody->SetLinearVelocity(birdStrlen ); //设置时会自动设置redbirdBody->SetAwake(true);
      //redbirdBody->ApplyForce(birdStrlen, redbirdBody->GetWorldCenter());
      //redbirdBody->ApplyAngularImpulse(-5);
      drag = false;
    }
  }
}

void disp_background()
{
  SDL_BlitSurface(bg, 0, pSc, 0);
  SDL_Rect slingRec;
  slingRec.x = 3;
  slingRec.y = 50;
  SDL_BlitSurface(sling, 0, pSc, &slingRec);
}

void playground()
{
  //*
  b2Body* ground;

  {///创建地面四方框
    b2BodyDef bd;
    //bd.position.Set(200.0f*ptm, 120.0f*ptm);
    bd.position.Set((200.f/ptm), (300.0f/ptm));
    ground = world->CreateBody(&bd);

    b2EdgeShape shape;

    b2FixtureDef sd;
    sd.shape = &shape;
    sd.density = 0.0f;
    sd.restitution = 0.4f;

    // Left vertical
    //shape.Set(b2Vec2(-200.0f*ptm, -120.0f*ptm), b2Vec2(-200.0f*ptm, 120.0f*ptm));
    shape.Set(b2Vec2(-200.0f/ptm, -300.f/ptm), b2Vec2(-200.0f/ptm, 300.0f/ptm));
    ground->CreateFixture(&sd);

    // Bottom horizontal
    shape.Set(b2Vec2(-200.0f/ptm, 300.0f/ptm), b2Vec2(200.0f/ptm, 300.0f/ptm));
    //shape.Set(b2Vec2(-200.0f*ptm, 120.0f*ptm), b2Vec2(200.0f*ptm, 120.0f*ptm));
    ground->CreateFixture(&sd);

  // Right vertical
    shape.Set(b2Vec2(200.0f/ptm, 300.0f/ptm), b2Vec2(200.0f/ptm, -300.0f/ptm));
    //shape.Set(b2Vec2(200.0f*ptm, -120.0f*ptm), b2Vec2(200.0f*ptm, 120.0f*ptm));
    ground->CreateFixture(&sd);

    // Top horizontal
    shape.Set(b2Vec2(200.0f/ptm, -300.0f/ptm), b2Vec2(-200.0f/ptm, -300.0f/ptm));
    //shape.Set(b2Vec2(-200.0f*ptm, -120.0f*ptm), b2Vec2(200.0f*ptm, -120.0f*ptm));
    ground->CreateFixture(&sd);
  }
  //*/
  /*
  b2BodyDef playgroundDef;
  playgroundDef.position.Set(2.0/ptm, (200)/ptm);
  memcpy(&ud[1].name,"playground",20);
  ud[1].mouseAble = false;
  playgroundDef.userData = &ud[1].name;
  b2Body *ground = world->CreateBody(&playgroundDef);

  b2PolygonShape groundShape;
  groundShape.SetAsBox(0.5/ptm,200./ptm);

  ground->CreateFixture(&groundShape, 0.0f);
  //*/

  /*
  b2Body* m_middle;
  b2Body* ground = NULL;
  {
    b2BodyDef bd; ///定义地面所用的刚体
    bd.position.Set(124.0f, 300.0f);
    ground = world->CreateBody(&bd); ///创建地面


    b2EdgeShape shape; ///定义角度形状物
    shape.Set(b2Vec2(-40.0f, 30.0f), b2Vec2(40.0f, 20.0f));
    ground->CreateFixture(&shape, 0.0f); ///地面上创建角度形状物
  }

  {
    b2PolygonShape shape;
    shape.SetAsBox(0.5f, 0.125f);

    b2FixtureDef fd;  ///定义像皮筋所用的材质
    fd.shape = &shape;
    fd.density = 20.0f; ///设置密度
    fd.friction = 0.2f;///设置摩擦

    b2RevoluteJointDef jd; ///定义旋转关节

    b2Body* prevBody = ground;
    for (int32 i = 0; i < 30; ++i)
    {///创建30个动态物体连接成像皮筋
      b2BodyDef bd;
      bd.type = b2_dynamicBody;
      bd.position.Set(-14.5f + 1.0f * i, 5.0f);
      b2Body* body = world->CreateBody(&bd);
      body->CreateFixture(&fd);

      b2Vec2 anchor(-15.0f + 1.0f * i, 5.0f);
      jd.Initialize(prevBody, body, anchor);
      world->CreateJoint(&jd);

      if (i == (30 >> 1))
      {
        m_middle = body;
      }
      prevBody = body;
    }

    b2Vec2 anchor(-15.0f + 1.0f * 30, 5.0f);
    jd.Initialize(prevBody, ground, anchor);
    world->CreateJoint(&jd);///创建连接体
  }
  //*/
  /*
  b2BodyDef bodyDef;
  bodyDef.position.Set(0, 0);

  // Tell the physics world to create the body
  b2Body *body = world->CreateBody(&bodyDef);

  b2ChainShape shape;
  //Make a list of points for the ground.
  b2Vec2 list[] = {b2Vec2(0,100), b2Vec2(50,100),b2Vec2(100,100)};
  shape.CreateLoop(list, 3);

  b2FixtureDef loopShapeDef;
  loopShapeDef.shape = &shape;
  body->CreateFixture(&loopShapeDef);
  */
}

void bird()
{
  b2CircleShape cirShape;
  //cirShape.m_radius = (redbird->w+(redbird->w/2))/ptm;
  cirShape.m_radius = (10.0f)/ptm;

  b2FixtureDef  birdFix;
  birdFix.density = 1;
  birdFix.friction = 3;
  birdFix.restitution = 0.7; ///弹性
  birdFix.shape = &cirShape;

  b2BodyDef birdDef;
  birdDef.type = b2_dynamicBody;
  //birdDef.position.Set((60.f+redbird->w/2.f)/ptm,(100.f+redbird->h/2.f)/ptm);
  birdDef.position.Set((90.f)/ptm,(60.f)/ptm);
  slingPoint.x = 90.f;
  slingPoint.y = 60.f;
  memcpy(&ud[0].name, "redbird",20);
  ud[0].mouseAble = true;
  birdDef.userData = &ud[0];

  b2Body *rbird = world->CreateBody(&birdDef);
  rbird->CreateFixture(&birdFix);

  //b2Vec2 force = b2Vec2(64 * rbird->GetMass(), (100 * rbird->GetMass()));
  //b2Vec2 force = b2Vec2((60 * rbird->GetPosition().x+50)*rbird->GetMass(), (120 * rbird->GetPosition().y+50)*rbird->GetMass());
  //rbird->ApplyLinearImpulse(force, rbird->GetWorldCenter());
  rbird->SetActive(false);
}

void myDraw::DrawPolygon(const b2Vec2* vertices, int32 vertexCount,
        const b2Color& color)
{
  if (vertexCount < 3)
    return;
  for (int32 i = 0; i < vertexCount; ++i)
  {
    Draw_Line(pSc,max(0.0f,vertices[i].x*ptm),max(0.0f,vertices[i].y*ptm),
        max(0.0f,vertices[i+1].x*ptm),max(0.0f,vertices[i+1].y*ptm),red);
  }

}

void myDraw::DrawSolidPolygon(const b2Vec2* vertices, int32 vertexCount,
    const b2Color& color)
{
  if (vertexCount < 3)
    return;
  for (int32 i = 0; i < vertexCount; ++i)
  {
    Draw_Line(pSc,(vertices[i].x*ptm),(vertices[i].y*ptm),
        (vertices[i+1].x*ptm),(vertices[i+1].y*ptm),red);
  }
}

void myDraw::DrawCircle(const b2Vec2& center, float32 radius, const b2Color& color)
{

}

void myDraw::DrawSolidCircle(const b2Vec2& center, float32 radius, const b2Vec2& axis,
    const b2Color& color)
{
  /*
  SDL_Rect rc;
  rc.x = center.x *ptm;
  rc.y = center.y *ptm;
  rc.h = (redbird->h);
  rc.w = (redbird->w);
  SDL_BlitSurface(redbird, 0, pSc, &rc);
  //*/
  Draw_Circle_Fill(pSc, center.x*ptm, center.y*ptm, radius*ptm, red, blue);
}

void myDraw::DrawSegment(const b2Vec2& p1, const b2Vec2& p2, const b2Color& color)
{
  Draw_Line(pSc,(Uint32)p1.x*ptm,(Uint32)p1.y*ptm,(Uint32)p2.x*ptm,(Uint32)p2.y*ptm,(Uint32)red);
}

void myDraw::DrawTransform(const b2Transform& xf)
{

}

