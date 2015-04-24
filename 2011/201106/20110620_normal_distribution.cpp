#include <boost/random.hpp>
#include <boost/random/normal_distribution.hpp>

int main(int argc, char* argv[])
{
  boost::mt19937 rng; // I don't seed it on purpouse (it's not relevant)

  boost::normal_distribution<> nd(atoi(argv[1]), atoi(argv[2]));

  boost::variate_generator<boost::mt19937&, boost::normal_distribution<> > var_nor(rng, nd);

  for (int i=0; i<10; ++i){
    int it = var_nor();
    std::cout << it << "\t" << ((int)(it*1000000))%11 << std::endl;
  }
  int i = 0; for (; i < 10; ++i)
  {
    double d = var_nor();
    std::cout << d << "\t" << ((int)(d*1000000))%11 << std::endl;
  }
}

