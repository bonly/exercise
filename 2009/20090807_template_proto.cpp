template <class T>
struct PrototypeCreator
{
   PrototypeCreator(T* pObj=0)
   {}
   T* Create()
   {
      return pPrototype_?pPrototype_->Clone():0;
   }
   T* GetPrototype(){return pPrototype_;}
   void SetPrototype(T* pObj){pPrototype_=pObj;}
private:
   T* pPrototype_;
};

class Widget
{
};

template<template<class Creator> class CreationPolicy>
class WidgetManager: public CreationPolicy<Widget>
{
};



int main()
{
  Widget *pPrototype=new Widget;
  WidgetManager<PrototypeCreator> mgr;
  mgr.SetPrototype(pPrototype);
  return 0;
}