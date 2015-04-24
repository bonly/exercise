//#include <codecvt>
#include <string>
#include <locale>
#include <string>
#include <cassert>

int main() {
  std::wstring_convert<std::codecvt_utf8<char32_t>, char32_t> convert;
  std::string utf8 = convert.to_bytes(0x5e9);
  assert(utf8.length() == 2);
  assert(utf8[0] == '\xD7');
  assert(utf8[1] == '\xA9');
}

/*
http://en.cppreference.com/w/cpp/locale/codecvt
*/

