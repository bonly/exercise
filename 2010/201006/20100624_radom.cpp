/*
 * 20100624_radom.cpp
 *
 *  Created on: 2012-8-7
 *      Author: bonly
 */

#include<ctime>
#include<cmath>
#include<iostream>
#include <fstream>
#include <boost/random.hpp>

#define PI 3.1416

using namespace std;
//using namespace boost;
using namespace boost::random;

inline double unif_rand()
{ //生成(0,1)的实数均匀分布
    //rand();
    return (rand() + 0.5) / (RAND_MAX + 1.0);
}
;

inline int unif_int(int a, int b)
{ //生成a到b-1的整数均匀分布
    return int(floor(a + (b - a) * unif_rand()));
}
;

inline double gaussien()
{ //标准正态分布的生成
    return sqrt(-2.0 * log(unif_rand())) * cos(2.0 * PI * unif_rand());
}
;
int main()
{
    srand(static_cast<unsigned int>(time(0)));
    rand(); //因为第一次的随机值不好，我们不要
    boost::uniform_int<> distribution(1, 100);
    boost::random::mt19937 engine;
    boost::random::mt19937::result_type random_seed =
                static_cast<boost::random::mt19937::result_type>(time(0));
    engine.seed(random_seed);
    engine.seed(static_cast<boost::random::mt19937::result_type>(random_seed));
    variate_generator<boost::random::mt19937, boost::uniform_int<> > random_choice(engine,
                distribution);
    boost::random::normal_distribution<> normal(0, 1);
    variate_generator<boost::random::mt19937, boost::random::normal_distribution<> > gaussian(engine, normal);

    std::ofstream flux("test.txt");
    //flux.precision(5);
    int N = 1000000;
    flux << "N = " << N << endl;
    double start, finish, duration;
    start = clock();
    for (int i = 0; i < N; i++)
    {
        random_choice();
    }
    finish = clock();
    duration = (double) (finish - start);
    flux << "boost的uniform_int(1,100)需要" << duration << "毫秒..." << endl << endl;

    start = clock();
    for (int i = 0; i < N; i++)
    {
        unif_int(1, 101);
    }
    finish = clock();
    duration = (double) (finish - start);
    flux << "自己的unif_int(1,100)需要" << duration << "毫秒..." << endl << endl;

    start = clock();
    for (int i = 0; i < N; i++)
    {
        gaussian();
    }
    finish = clock();
    duration = (double) (finish - start);
    flux << "boost的标准正态分布需要" << duration << "毫秒..." << endl << endl;

    start = clock();
    for (int i = 0; i < N; i++)
    {
        gaussien();
    }
    finish = clock();
    duration = (double) (finish - start);
    flux << "自己的标准正态分布需要" << duration << "毫秒..." << endl << endl;

    flux.close();
    return 0;
}

