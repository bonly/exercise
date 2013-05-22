//============================================================================
// Name        : spirit.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <boost/spirit/include/qi.hpp>
#include <iostream>
#include <cstring>
#include <cstdlib>
using namespace std;
using namespace boost;

template <class P, class Attr>
bool myparser(P const &p, const char* input, Attr const& excepted)
{
    char const* f(input);
    const char* l(f + strlen(f));
    Attr attr;
    return spirit::qi::parse(f, l, p, attr) && f==l && attr == excepted;
}

int main()
{
	assert(myparser(spirit::qi::int_[cout << spirit::qi::_1], "1234", 1234));
	assert(myparser(spirit::int_, "1234", 12345));
	return 0;
}
