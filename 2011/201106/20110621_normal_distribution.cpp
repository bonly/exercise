#include <time.h>
#include <boost/random/normal_distribution.hpp>
#include <boost/random/mersenne_twister.hpp>
#include <boost/random/variate_generator.hpp>

double gen_normal(void)
{
  boost::variate_generator<boost::mt19937, boost::normal_distribution<> >
    generator(boost::mt19937(time(0)),
           boost::normal_distribution<>());

  double r = generator();
  return r;
}
int main(void)
{
  for(size_t i=0; i<10; ++i)
    std::cout<<gen_normal()
             <<std::endl;
}

// output:
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
// 0.121596
//http://www.bnikolic.co.uk/blog/cpp-boost-rand-normal.html
