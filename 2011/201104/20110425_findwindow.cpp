#include <windows.h>
int main(int argc, char* argv[]){
	if (FindWindow("notepad", NULL) != NULL){
		MessageBox(NULL, "find windows", "OK", MB_OK);
	}else{
		MessageBox(NULL, "not find windows", "error", MB_OK);
	}
}