/*
 * 20100610_mpi.cpp
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

    world.send(1, 0, std::string("hello from 0\n"));
    std::string str;
    world.recv(1, 0,str);
    std::clog << "recv from 1: " << str << std::endl;
    return 0;
}

