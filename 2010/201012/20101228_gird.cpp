#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/layout.hpp>

int main(){
	using namespace nana::gui;
	form fm;
	button btn(fm);
	btn.caption(STR("Button"));
	
	gird gobj(fm);
	
	gird *child_gird_a = gobj.add(0,0);
	
	//距离a 10，宽80
	gird *child_gird_b = gobj.add(10, 80);
	
	gird *child_gird_c = gobj.add(0, 10);
	
	//离顶10piels,20pixels高
	child_gird_b->push(btn, 10, 20);
	fm.show();
	exec();
}
