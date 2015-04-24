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
//#include <boost/iostreams/filter/gzip.hpp>
#include <boost/iostreams/filter/bzip2.hpp>
//#include <boost/iostreams/filter/zlib.hpp>

class mybuf : public std::streambuf{
public:
    mybuf(char* begin, char* cur, char* end){ //读的时候要用
        this->setg((char*)begin, (char*)cur, (char*)end);
    }   
    mybuf(char* begin, char* end){ //写的时候用
        this->setp(begin, end);
    }   
};


int main(){
    using namespace std;
    ifstream file("data.dat", ios_base::in | ios_base::binary);
    boost::iostreams::filtering_streambuf< boost::iostreams::input > in;
    //in.push(boost::iostreams::gzip_decompressor());
    in.push(boost::iostreams::bzip2_decompressor());
    //in.push(boost::iostreams::zlib_decompressor());
    in.push(file);
    
    boost::iostreams::copy(in, cout);

    return 0;
}

/*
 * zlib1g-dev:i386 -lz
 * boost_iostreams
 */
