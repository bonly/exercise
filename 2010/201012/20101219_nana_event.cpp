#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>

int main()
{
	using namespace nana::gui;
	form fm;
	button btn(fm, nana::size(100, 20));
	btn.caption(STR("Quit"));
	btn.make_event<events::click>(API::exit);
	fm.show();
	exec();
}
