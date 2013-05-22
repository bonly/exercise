struct A{
  int aInt;
  int bInt;
};

void func1(A a){
}

void func2(A &a){
}

void func3(const A a){
}

int main(){
   func1(A({3,4}));
   //func2(A({3,4}));
   func3(A({3,4}));
   return 0;
}

