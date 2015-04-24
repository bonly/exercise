#include <locale>
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

int main(){
    std::string abc("这是一个测试abc");
    std::wcout << mbs_to_wcs(abc, std::locale("zh_CN.utf8")) << std::endl;
    std::wcout << mbs_to_wcs(abc, std::locale(""));
    //std::clog << mbs_to_wcs(abc);

   return 0;
}

