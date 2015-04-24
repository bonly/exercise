#include <time.h>
#include <boost/random/normal_distribution.hpp>
#include <boost/random/mersenne_twister.hpp>
#include <boost/random/variate_generator.hpp>

template<class T>
T gen_normal_4(T generator,
            std::vector<double> &res)
{
  for(size_t i=0; i<res.size(); ++i)
    res[i]=generator();
  // Note the generator is returned back
  return  generator;
}

int main(void)
{
  boost::variate_generator<boost::mt19937, boost::normal_distribution<> >
    generator(boost::mt19937(time(0)),
              boost::normal_distribution<>());

  std::vector<double> res(10);
  // Assigning back to the generator ensures the state is advanced
  generator=gen_normal_4(generator, res);

  for(size_t i=0; i<10; ++i)
    std::cout<<res[i]
             <<std::endl;

  generator=gen_normal_4(generator, res);

  for(size_t i=0; i<10; ++i)
    std::cout<<res[i]
             <<std::endl;
}
// Output:
// -2.01509
// -1.87061
// 1.35597
// -0.434768
// -0.0233686
// -0.0525241
// -0.521989
// -0.362958
// -0.557344
// -1.02654
// -0.506174
// 0.299165
// 0.847568
// 0.0126109
// -1.33701
// 0.892904
// 0.612492
// 1.01212
// 1.33039
// 0.487829
