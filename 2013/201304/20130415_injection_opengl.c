/* 
 * eve.c: Simple LD_PRELOAD injection for OpenGL applications.
 * 
 * gcc -std=c99 -Wall -Werror -m32 -O0 -fpic -shared -ldl -lGL -o eve.so eve.c
 * wineg++  -Wall   -O0 -fpic -shared -ldl -lGL  -o eve.so 20130415_injection_opengl.c
 * -m32 是32位上用的，一般需要这个
 * How to use with WINE:
 *   wineserver -k
 *   export LD_PRELOAD=/path/to/eve.so
 *   wine Wow.exe
 *
 */

/* These libraries are necessary for the hook */
#include <dlfcn.h>
#include <stdlib.h>
#include <GL/gl.h>

/* "Injected" stuff */
#include <stdio.h>
#include <stdint.h>
#include <string.h>
void doevil();

/* Hook function */
void glClear(GLbitfield mask) {
	typedef void(*func)(GLbitfield mask);
	static void (*lib_glClear)(GLbitfield mask) = NULL;
	void* handle;
	char* errorstr;

	if(!lib_glClear) {
		/* Load real libGL */
		handle = dlopen("/usr/lib32/libGL.so", RTLD_LAZY);
		if(!handle) {
			fputs(dlerror(), stderr);
			exit(1);
		}

		/* Fetch pointer of real glClear() func */
		lib_glClear = (func)dlsym(handle, "glClear");
		if( (errorstr = dlerror()) != NULL ) {
			fprintf(stderr, "dlsym fail: %s\n", errorstr);
			exit(1);
		}
	}

	/* Woot */
	doevil();

	/* Call real glClear() */
	lib_glClear(mask);
}


uint32_t read_uint( uint32_t addr ) { 
	return *((uint32_t*) addr); 
}

float read_float( uint32_t addr ) { 
	return *((float*) addr); 
}
	
/* Here be dragons */
void doevil() {
	static int framecnt = 0;
	framecnt++;

	// calling game functions works too! (WoW 3.3.0a)
	typedef int (*pt)();
	static int (*ClntObjMgrGetActivePlayer)() = (pt) 0x0047A2B0;

	printf("doevil(), frame %d... ", framecnt);
	if( ClntObjMgrGetActivePlayer() == 0 ) {
		printf("not logged in.\n");
	} else {
		char p_name[16];
		uint32_t p_base;
		float p_x, p_y;
		
		strncpy( p_name, (char*) 0x00C923F8, 16 );
		if( read_uint(0x00CF7C00) ) {
			p_base = read_uint(read_uint( read_uint( 0x00CF7C00 ) + 0x34 ) + 0x24);
			p_x    = read_float( p_base + 0x798 );
			p_y    = read_float( p_base + 0x79C );
		}

		printf("p_name: %s, x/y: %.2f/%.2f\n", p_name, p_x, p_y);
	}
}