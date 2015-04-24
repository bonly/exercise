#include <locale>
#include <vector>
#include <stdexcept>
#include <string>
#include <iostream>

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
    std::wstring abc(L"这是一个测试abc");
    std::cout << wcs_to_mbs(abc, std::locale("zh_CN.utf8")) << std::endl;
    std::cout << wcs_to_mbs(abc, std::locale(""));
    return 0;
}

