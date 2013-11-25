#include <unistd.h>
#include <windows.h>
int main(int argc, char* argv[]){
	char buf[256];
	gethostname(buf, sizeof(buf));
	strcat(buf, " says, 'Hello world!'");
	MessageBox(NULL, buf, "Junk", MB_OK);
	return 0;
}

//wineg++ file