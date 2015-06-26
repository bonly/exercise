#include <boost/scope_exit.hpp>
#include <cstdlib>
#include <cstdio>
#include <cassert>

int main(){
	std::FILE* f = std::fopen("example_file.txt", "w");
	assert(f);
	BOOST_SCOPE_EXIT(f){
		std::fclose(f);
	}BOOST_SCOPE_EXIT_END
}

