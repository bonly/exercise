/*
 * 20100611_mpi.cpp
 *
 *  Created on: 2012-7-22
 *      Author: bonly
 */

#include <iostream>
#include <boost/mpi.hpp>
#include <string>
#include <boost/serialization/string.hpp>

namespace mpi=boost::mpi;

int main(int argc, char* argv[])
{
    mpi::environment env(argc, argv);
    mpi::communicator world;

    std::string str;
    world.recv(0, 0,str);
    std::clog << "recv from 0: " << str << std::endl;
    world.send(0, 0, std::string("hello from 1\n"));
    return 0;
}


