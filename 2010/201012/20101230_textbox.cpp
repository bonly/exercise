#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/textbox.hpp>

int main(){
	using namespace nana::gui;
	form frm;
	button btn(frm, nana::rectangle(5, 5, 94, 23));
	btn.caption(STR("开始"));
	
	textbox bx(frm, nana::rectangle(5, 40, 100, 100));
		
	btn.make_event<events::click>([&](){
		  bx.append (nana::string(STR("atext\n")), false);
		  }
		);
	frm.show();
	exec();	
}
