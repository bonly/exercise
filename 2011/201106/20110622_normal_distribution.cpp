template<class T>
double gen_normal_2(T generator)
{
  return generator();
}

int main(void)
{
  boost::variate_generator<boost::mt19937, boost::normal_distribution<> >
    generator(boost::mt19937(time(0)),
              boost::normal_distribution<>());

  for(size_t i=0; i<10; ++i)
    std::cout<<gen_normal_2(generator)
             <<std::endl;
}

// Output:
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
// 0.898862
