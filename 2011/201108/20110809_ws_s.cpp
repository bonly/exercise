#include <locale>
#include <assert.h>
#include <vector>
#include <stdexcept>
#include <string>
#include <iostream>

std::wstring 
mbs_to_wcs(std::string const& str, std::locale const& loc = std::locale()) {
    typedef std::codecvt<wchar_t, char, std::mbstate_t> codecvt_t;
    codecvt_t const& codecvt = std::use_facet<codecvt_t>(loc);
    std::mbstate_t state = std::mbstate_t();
    std::vector<wchar_t> buf(str.size() + 1);
    char const* in_next = str.c_str();
    wchar_t* out_next = &buf[0];
    std::codecvt_base::result r = codecvt.in(state, 
        str.c_str(), str.c_str() + str.size(), in_next, 
        &buf[0], &buf[0] + buf.size(), out_next);
    if (r == std::codecvt_base::error)
        throw std::runtime_error("can't convert string to wstring");   
    return std::wstring(&buf[0]);
}

std::string 
wcs_to_mbs(std::wstring const& str, std::locale const& loc = std::locale()) {
    typedef std::codecvt<wchar_t, char, std::mbstate_t> codecvt_t;
    codecvt_t const& codecvt = std::use_facet<codecvt_t>(loc);
    std::mbstate_t state = std::mbstate_t();
    std::vector<char> buf((str.size() + 1) * codecvt.max_length());
    wchar_t const* in_next = str.c_str();
    char* out_next = &buf[0];
    std::codecvt_base::result r = codecvt.out(state, 
        str.c_str(), str.c_str() + str.size(), in_next, 
        &buf[0], &buf[0] + buf.size(), out_next);
    if (r == std::codecvt_base::error)
       throw std::runtime_error("can't convert wstring to string");   
    return std::string(&buf[0]);
}

int main(){
// 전역 locale 설정
std::locale::global(std::locale(""));

std::string mbs1 = "abcdef가나다라";
std::wstring wcs1 = L"abcdef가나다라";
std::string mbs2 = wcs_to_mbs(wcs1);
assert(mbs1 == mbs2);
std::wstring wcs2 = mbs_to_wcs(mbs1);
assert(wcs1 == wcs2);

   return 0;
}

/*
wstring s2ws(const std::string str){
    typedef std::codecvt_utf8<wchar_t> convert_typeX;
    std::wstring_convert<convert_typeX, wchar_t> converterX;

    return converterX.from_bytes(str);
}

string ws2s(const std::wstring wstr){
    typedef std::codecvt_utf8<wchar_t> convert_typeX;
    std::wstring_convert<convert_typeX, wchar_t> converterX;

    return converterX.to_bytes(wstr);
}
*/
