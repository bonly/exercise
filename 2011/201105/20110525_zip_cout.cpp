/**
 * @file 20110525_zip_cout.cpp
 * @brief 
 * @author bonly
 * @date 2013年11月18日 bonly Created
 */

#include <sstream>
#include <streambuf>
#include <boost/iostreams/filtering_streambuf.hpp>
#include <boost/iostreams/copy.hpp>
#include <boost/iostreams/filter/gzip.hpp>
//#include <boost/iostreams/filter/bzip2.hpp>
//#include <boost/iostreams/filter/zlib.hpp>

int main(){
    using namespace std;
    stringstream is;
    //is << "this is a test"; //要创建完后再放数据, 这里放的数据将没用
    boost::iostreams::filtering_streambuf< boost::iostreams::input > in;


    in.push(boost::iostreams::gzip_compressor());
    //in.push(boost::iostreams::zlib_compressor());
    //in.push(boost::iostreams::bzip_decompressor());
    in.push(is);

    is << "this is a test"; //要创建完后再放数据
    boost::iostreams::copy(in, cout);
}

