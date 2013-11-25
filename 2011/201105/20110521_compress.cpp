#include <boost/iostreams/filtering_stream.hpp>
#include <boost/iostreams/device/file_descriptor.hpp>
#include <boost/iostreams/device/file.hpp>
#include <boost/iostreams/filter/bzip2.hpp>
#include <boost/iostreams/filter/gzip.hpp>
#include <boost/iostreams/copy.hpp>
#include <iostream>
#include <sstream>


int main(){
 try{
  std::string dest;
  boost::iostreams::filtering_ostream out;
  out.push(boost::iostreams::gzip_compressor());
  //out.push(boost::iostreams::file_sink("test.txt"));
  out.push(boost::iostreams::back_inserter(dest));
  //boost::iostreams::write(out, "hello", 5);
  boost::iostreams::copy(std::stringstream("hello"), out);
  std::cout << "dest: " << dest << std::endl;
  boost::iostreams::filtering_istream in;
  in.push(boost::iostreams::gzip_decompressor());
  //in.push(boost::iostreams::file_source("test.txt"));
  std::stringstream ss(dest);
  in.push(ss);
  char text[100]={0};
  boost::iostreams::read(in, text, 4);
  //boost::iostreams::copy(in, std::stringstream(text));
  std::cout << "text: "<< text << std::endl;
 }catch(std::exception& e){
   std::cout << "exception: " << e.what() << std::endl;
 }catch(...){
   std::cout << "unknown exception. " << std::endl;
 }
 system("pause");
 return 0;
}

/*
iostreams主要有两类东西组成，一个是device，另一个是filter，可以到源码目录下找，iostreams目录下有这两个目录可以找到相关类。
device像是一种设备，不能单独使用，要配合普通流stream或stream_buffer来使用，可将流中的数据输入/输出到这个设备上，可分为
Source，它以读取的方式访问字符序列，如：file_source 做文件输入。
Sink，它以写的方式访问字符序列，如：file_sink 做文件输出。
stream<file_source> 那么这个就是一个文件输入流，类似于ifilestream，而stream<file_sink>就是一个文件输出流，类似于ofilestream。
filter像一种过滤器，和device一样，也是不能单独使用的，要配合过滤流filtering_stream或filtering_streambuf来使用，将流中的数据按一种规则过滤，可分为：
InputFilter，过滤通过Source读取的输入，如：gzip_decompressor 按gzip算法解压流中的数据。
OutputFilter，过滤向Sink写入的输出，如：gzip_compressor 按gzip算法压缩流中的数据。
但filtering_stream是要维护一个filter的链表的，以device为结束。输出过滤流filtering_ostream，是按顺序执行filter，然后再输出到devic上，如：
压缩时
filtering_ostream out;
out.push(gzip_compressor()); //gzip OutputFilter
out.push(bzip2_compressor());//bzip2 OutputFilter
out.push(boost::iostreams::file_sink("test.txt"));//以file_sink device结束
这就会将流的数据先按gzip压缩，然后再按bzip2压缩之后，才输出到text.txt文件中。

解压时
filtering_istream in;
in.push(gzip_decompressor());/gzip InputFilter
in.push(bzip2_decompressor());/bzip2 InputFilter
in.push(file_source("test.txt"));
这时候先将test.txt文件中数据读出，然后按bzip2解压，然后再按gzip解压，存入in流中，正好是压缩的逆序。
*/
