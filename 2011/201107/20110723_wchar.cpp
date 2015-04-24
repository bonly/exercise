#include<string>
#include<iostream>
#include<locale>
using namespace std;
 
int main(){
locale prevloc;
locale loc("chs");
 
string str1("string class");
string str2("");
wstring wstr1(L"wstring class");          //L
wstring wstr2(L"");
 
prevloc = cout.imbue(locale(""));
cout<<"Default Locale: "<<prevloc.name()<<endl;
cout<<"System Locale: "<<locale("").name()<<endl;
//cout<<"C\n"<<L"w-string\n"<<str1<<\n<<str2<<\n<<endl;
 
prevloc = wcout.imbue(loc);   //wstr2
wcout<<"Default Locale: "<<prevloc.name().c_str()<<endl;    // .c_str() 
wcout<<"chs Locale Name: "<<loc.name().c_str()<<endl;
wcout<<"C-string\n"<<"C\n"<<L"\n"<<wstr1<<\n<<wstr2<<\n<<endl;
   return 0;
}


/*
        1.cout  string wcout  wstring ()
        2.wstring  L"xxx"  string  L 
        3.locale ("C") cout  Cstd::string
     L"xxx" 
          locale ("C") wcout Cstd::wstring
     locale ("chs")std::wstring C 
 
        string cout "C-style " 
                  wstring wcout L""  wcout  locale 
                  
*/                  