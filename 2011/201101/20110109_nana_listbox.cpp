#include <nana/gui/wvl.hpp> 
#include <nana/gui/widgets/listbox.hpp> 
int main() { 
	  nana::gui::form form(nana::gui::API::make_center(400, 300)); 
    nana::gui::listbox listbox(form, nana::rectangle(10, 10, 380, 200)); 
		
		listbox.append_header(STR("Column 1"), 100); 
		listbox.append_header(STR("Column 2"), 100); 
		listbox.append_header(STR("Column 3"), 100); 
		listbox.append_item(STR("Item 0")); 
		listbox.append_item(STR("Item 1")); 
		listbox.set_item_text(1, 1, STR("subitem")); 
		listbox.set_item_text(1, 2, STR("行1位置2")); 
		
		listbox.append_categ(STR("Test Category")); 
		listbox.append_item(1, STR("Item 0")); 
		listbox.append_item(1, STR("Item 1")); 
		listbox.append_item(1, STR("Item 2")); 
		listbox.append_item(1, STR("Item 3")); 
		listbox.append_item(1, STR("Item 4")); 
		listbox.append_item(1, STR("Item 5")); 
		listbox.set_item_text(1, 1, 1, STR("索引1位置1")); 
		listbox.set_item_text(1, 3, 2, STR("索引3位置2")); 
		
		form.show(); 
		nana::gui::exec(); 
}
