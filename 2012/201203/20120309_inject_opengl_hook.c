#pragma once
// http://www.parashift.com/c++-faq-lite/mixing-c-and-cpp.html

/* In your C++ source:
 *
 *	void hookfunc() { ... }
 *
 */

#ifdef __cplusplus

extern "C" void hookfunc();

#else

void hookfunc();

#endif

//#include "hook-glx.h"

#define _GNU_SOURCE
#include <dlfcn.h>

#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include <GL/gl.h>
#include <GL/glx.h>

void hookfunc();

/* glXSwapBuffers() is very appropriate as an "end scene" hook because
 * WoW calls it exactly once, at the end of each frame.
 *
 * Unfortunately WINE needs to be patched to allow hooking this via
 * LD_PRELOAD.
 *
 * 1) Instead of using the default RTLD loader, WINE retrieves the address
 *    for glXGetProcAddressARB directly from the OpenGL library.
 *    This is done to enable binary compatibility to systems with or without
 *    OpenGL.
 *
 *    Solution: patch WINE to use the default loader
 *
 * 2) WINE gets the pointers to glx functions via glXGetProcAddressARB
 *    
 *    => I'm hooking glXGetProcAddressARB and return my own
 *       pointer to glXSwapBuffers.
 *
 */

typedef __GLXextFuncPtr (*fp_glXGetProcAddressARB) (const GLubyte*);
typedef __GLXextFuncPtr (*fp_glXSwapBuffers)(Display* dpy, GLXDrawable drawable);

// glXSwapBuffers
fp_glXSwapBuffers real_glXSwapBuffers;

void my_glXSwapBuffers(Display* dpy, GLXDrawable drawable) {
	real_glXSwapBuffers(dpy, drawable);
	hookfunc();
}

// glXGetProcAddressARB
__GLXextFuncPtr glXGetProcAddressARB (const GLubyte* procName)
{
	__GLXextFuncPtr result;
	printf("* hook-glx.c: glXGetProcAddressARB(\"%s\")\n", procName);

	// Fetch pointer of actual glXGetProcAddressARB() function
	static fp_glXGetProcAddressARB lib_getprocaddr = NULL;
	if(!lib_getprocaddr)
	{
		char* errorstr;
		lib_getprocaddr = (fp_glXGetProcAddressARB) 
			dlsym(RTLD_NEXT, "glXGetProcAddressARB");
		if( (errorstr = dlerror()) != NULL )
		{
			fprintf(stderr, "dlsym fail: %s\n", errorstr);
			exit(1);
		}
	}
	result = lib_getprocaddr(procName);

	// Return our own function pointers
	if( strcmp( (const char*) procName, "glXSwapBuffers" ) == 0 )
	{
		real_glXSwapBuffers = (fp_glXSwapBuffers) result;
		return (__GLXextFuncPtr) my_glXSwapBuffers;
	}
	
	// Return default function pointer
	return lib_getprocaddr(procName);
}

//http://www.ownedcore.com/forums/world-of-warcraft/world-of-warcraft-bots-programs/wow-memory-editing/276206-linux-simple-injection-ld_preload.html
//http://www.linuxjournal.com/article/7795
