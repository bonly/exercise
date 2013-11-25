/**
 * @file 20110528_unzip_from_buf.cpp
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

template <class cha_type, class iterator_type>
struct my_source {
    typedef cha_type char_type;
    typedef boost::iostreams::source_tag category;

    iterator_type& it;
    iterator_type end;

    my_source(iterator_type& it, iterator_type end = {}) : it(it), end(end)
    { }

    std::streamsize read(char* s, std::streamsize n) {
        std::streamsize result = 0;
        while ((it!=end) && n--) {
            ++result;
            *s++ = *it++;
        }
        return result;
    }
};

class mybuf : public std::streambuf{
public:
    mybuf(char* begin, char* cur, char* end){
        this->setg(begin, cur, end); //读入用的需指定开始/当前/结束指针
    }
    mybuf(char* begin, char* end){
        this->setp(begin, end); //写出时需指定开始/结束指针
    }
};
int main(){
    using namespace std;
    {
    stringstream is;
    //is << "this is a test"; //要创建完后再放数据, 这里放的数据将没用
    boost::iostreams::filtering_streambuf< boost::iostreams::input > in;


    in.push(boost::iostreams::gzip_compressor());
    //in.push(boost::iostreams::zlib_compressor());
    //in.push(boost::iostreams::bzip_decompressor());
    in.push(is);

    is << "this is a test"; //要创建完后再放数据
    //boost::iostreams::copy(in, cout);

    stringstream out;
    boost::iostreams::copy(in, out);
    out.seekp(0, ios::end);
    clog << out.tellp();

    }

    { ///把char[]转换成streambuf再操作
        char bus[]="this is a test";
        stringstream is;

        boost::iostreams::filtering_streambuf< boost::iostreams::input> in;
        in.push(boost::iostreams::gzip_compressor());
        in.push(is);

        is.rdbuf()->pubsetbuf(bus, sizeof(bus));

        stringstream out;
        boost::iostreams::copy(in, out); //return value std::streamsize
        out.seekp(0, ios::end);
        clog << out.tellp();
    }

    {
        std::string const rawdata {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' };
        std::istringstream iss(rawdata, std::ios::binary);

        boost::iostreams::filtering_streambuf<boost::iostreams::input> def;
        def.push(boost::iostreams::zlib_decompressor());
        def.push(iss);
        boost::iostreams::copy(def, std::cout);
    }

    {
        std::string const rawdata {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' };
        std::istringstream iss(rawdata, std::ios::binary); //必须是istringstream,stringstream不行

        auto start = std::istreambuf_iterator<char>(iss);
        my_source<char, decltype(start)> data(start);

        boost::iostreams::filtering_istreambuf def;
        def.push(boost::iostreams::zlib_decompressor());
        def.push(data);

        boost::iostreams::copy(def, std::cout);
    }

    {  //必须是stirng _data,不能是char data[]
        //char  _data[] = {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' }; //加\0结束符也没有用
        std::string  _data {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' }; //加\0结束符也没有用
        std::istringstream is(_data, std::ios::binary);

        boost::iostreams::filtering_streambuf< boost::iostreams::input > steam;
        steam.push(boost::iostreams::zlib_decompressor());
        steam.push(is);

        //boost::iostreams::copy(steam, std::cout);

        std::stringstream out;
        char tmp[1024]="";
        out.rdbuf()->pubsetbuf((char*)tmp, 1024);

        std::streamsize slen = boost::iostreams::copy(steam, out);
        std::clog << tmp;
    }
    {
        char  _data[] = {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' }; //加\0结束符也没有用
        std::string  dat(_data,sizeof(_data));
        std::istringstream is(dat, std::ios::binary);

        boost::iostreams::filtering_streambuf< boost::iostreams::input > steam;
        steam.push(boost::iostreams::zlib_decompressor());
        steam.push(is);

        //boost::iostreams::copy(steam, std::cout);

        std::stringstream out;
        char tmp[1024]="";
        out.rdbuf()->pubsetbuf((char*)tmp, 1024);

        std::streamsize slen = boost::iostreams::copy(steam, out);
        std::clog << tmp;
    }
    {
        char  _data[] = {'x', '\234', '\313', 'H', '\315', '\311', '\311', 'W', '(', '\317', '/', '\312', 'I', '\341', '\002', '\0', '\036', 'r', '\004', 'g' }; //加\0结束符也没有用

        mybuf dat(_data, _data, _data+sizeof(_data));
        std::istream is(&dat);

        boost::iostreams::filtering_streambuf< boost::iostreams::input > steam;
        steam.push(boost::iostreams::zlib_decompressor());
        steam.push(is);

        //boost::iostreams::copy(steam, std::cout);

        char tmp[1024]="";
        mybuf out(tmp, tmp+1024);

        std::streamsize slen = boost::iostreams::copy(steam, out);
        std::clog <<"len["<<slen<< "]time: " << tmp << endl;
    }

    {
        char bus[]="this is a test";
        mybuf data(bus, bus, bus+sizeof(bus));
        std::istream is(&data);

        boost::iostreams::filtering_streambuf< boost::iostreams::input> in;
        in.push(boost::iostreams::gzip_compressor());
        in.push(is);

        char ot[1024];
        mybuf out(ot, ot+1024);
        int len = boost::iostreams::copy(in, out);
        clog << len << ":" << ot << endl;

        // 反过来
        mybuf out2in(ot, ot, ot+len);
        std::istream ins(&out2in);

        boost::iostreams::filtering_streambuf< boost::iostreams::input> in2;
        in2.push(boost::iostreams::gzip_decompressor());
        in2.push(ins);

        char str[1024];
        mybuf re(str, str+1024);
        int len2 = boost::iostreams::copy(in2, re, len);
        clog << len2 << ":" << str << endl;
    }
    return 0;
}

/*
 * zlib1g-dev:i386 -lz
 * boost_iostreams
 */

