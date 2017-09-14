/*
A的变动不影响B，B与A的行为一致
*/
#include <iostream>

class Fa{
	public:
	   Fa(){};
	   virtual void Pr()=0;
};

class A : public Fa{
	public:
		void Pr(){
			std::cout << "ok" << std::endl;
		}
};

class B {
	private:
	    Fa* fa;

	public:
		void Pr(){
			fa->Pr();
  		}
  		B(Fa *f){
  			fa = f;
  		}

};

int main(){
   A aa;
   B bb(&aa);

   aa.Pr();
   return 0;
}
