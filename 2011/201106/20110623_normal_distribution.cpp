#include <time.h>
#include <boost/random/normal_distribution.hpp>
#include <boost/random/mersenne_twister.hpp>
#include <boost/random/variate_generator.hpp>
template<class T>
double gen_normal_3(T &generator)
{
  return generator();
}

// Version that fills a vector
template<class T>
void gen_normal_3(T &generator,
              std::vector<double> &res)
{
  for(size_t i=0; i<res.size(); ++i)
    res[i]=generator();
}

int main(void)
{
  boost::variate_generator<boost::mt19937, boost::normal_distribution<> >
    generator(boost::mt19937(time(0)),
              boost::normal_distribution<>());

  for(size_t i=0; i<10; ++i)
    std::cout<<gen_normal_3(generator)
             <<std::endl;
}
// Output:
// -0.643475
// 0.144729
// 0.439714
// 0.481678
// 0.402485
// 0.416421
// -1.59029
// -0.800964
// -0.621854
// -0.150999
