/**
 * @file 20110523_zip.cpp
 * @brief 
 * @author bonly
 * @date 2013年11月17日 bonly Created
 */

#include <fstream>
#include <iostream>
#include <boost/iostreams/filtering_streambuf.hpp>
#include <boost/iostreams/copy.hpp>
#include <boost/iostreams/filter/gzip.hpp>
//#include <boost/iostreams/filter/bzip2.hpp>
//#include <boost/iostreams/filter/zlib.hpp>

int main(){
    using namespace std;
    ifstream file("hello.gz", ios_base::in | ios_base::binary);
    boost::iostreams::filtering_streambuf< boost::iostreams::input > in;
    in.push(boost::iostreams::gzip_decompressor());
    //in.push(boost::iostreams::zlib_compressor());
    //in.push(boost::iostreams::bzip_decompressor());
    in.push(file);
    boost::iostreams::copy(in, cout);
    return 0;
}

/*
 * zlib1g-dev:i386 -lz
 * boost_iostreams
 */
