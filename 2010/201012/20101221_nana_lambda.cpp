#include <nana/gui/wvl.hpp>
#include <iostream>

int main()
{
	using namespace nana::gui;

	form fm;
	fm.make_event<events::click>(
			[]{ std::cout<<"form is clicked"<<std::endl; }
		);
	fm.show();
	exec();
}
