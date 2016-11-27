#include <iostream>

int func(int x) {
	int countx = 0;
	while(x)
	{
		countx ++;
		x = x&(x-1);
	}
	return countx;
} 

int main(int argc, char* argv[]){
	int m = 9999;
	int k = func(m);
	std::cout << m << ":" << k << std::endl;

	m = 3111;
	k = func(m);
	std::cout << m << ":" << k << std::endl;

	m = 8;
	k = func(m);
	std::cout << m << ":" << k << std::endl;

	m = 1;
	k = func(m);
	std::cout << m << ":" << k << std::endl;

	m = 346;
	k = func(m);
	std::cout << m << ":" << k << std::endl;

	return 0;
}
