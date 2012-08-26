FilePre=20100526

#export CPATH=/usr/local/include:$(HOME)/ACE_wrappers/:/cygdrive/d/boost_1_36_0:$(HOME)/src/include:/cygdrive/d/oracle/oci/include
#export LIBRARY_PATH=/cygdrive/d/boost_1_37_0/stage/lib:$HOME/usr/w32api/lib:/usr/local/lib:$(HOME)/ACE_wrappers/lib:/cygdrive/d/oracle/oci/lib/msvc

MHOME=E:/mingw/msys/1.0/home/Bonly
#if use wssocket add -l ws2_32 -l mswsock -D__USE_W32_SOCKETS
WIN_SOCK_LIB=-l ws2_32 -l mswsock -l boost_system-gcc34-mt-1_36.dll

LIBS= wtsapi32 advapi32 netapi32
#mingw32 
#glu32 glut32 opengl32 (for opengl)
#boost_thread-gcc34-mt-1_37 boost_date_time-gcc34-mt-1_37
#boost_regex-gcc34-mt-1_37 
#boost_system-gcc34-mt-1_37 

INC = 
#"E:\mingw\lib\gcc\mingw32\4.6.2\include\c++" e:/mingw/include e:/workspace/box D:\BREW\inc "C:\Program Files\Microsoft Visual Studio 8\VC\PlatformSDK\Include\gl"
CFLAGS=-O0 -g3 -Wall -fmessage-length=0
#-console -mwindows 
#-Wl,--whole-archive
#-enable-auto-import 

CXXFLAGS= $(CFLAGS) -D_WIN32_WINNT=0x0501
#-D__USE_W32_SOCKETS
#-m32 32bit
#-m64 64bit
CXX = g++ -fno-omit-frame-pointer

WIN_LIB=-l kernel32 -l user32 -l gdi32 -lwinmm -ldxguid

LIBPATH_BOOST=
#-L /cygdrive/d/boost_1_37_0/stage/lib

LIBPATH= /e/mingw/lib E:/temp/Box2D D:/SDL-1.2.15/build/.libs 
#E:/mingw/lib e:/workspace/Box2D/Debug E:/workspace/glui/Debug E:/workspace/freeglut/Debug
#-L $(HOME)/usr/w32api/lib

#SDLLIBS=$(shell sdl-config --libs)
#SDLFlAGS=$(shell sdl-config --cflags)
SDLFLAGS=-I$(MHOME)/opt/include/SDL -D_GNU_SOURCE=1 -Dmain=SDL_main 
SDLLIBS=-L$(MHOME)/opt/lib -lmingw32 -lSDLmain -lSDL  -liconv -lm -luser32 -lgdi32 -lwinmm

gtk_inc=-mms-bitfields -ID:/GTK/include/gtkmm-2.4 -ID:/GTK/lib/gtkmm-2.4/include -ID:/GTK/include/glibmm-2.4 -ID:/GTK/lib/glibmm-2.4/include -ID:/GTK/include/gdkmm-2.4 -ID:/GTK/lib/gdkmm-2.4/include -ID:/GTK/include/pangomm-1.4 -ID:/GTK/include/atkmm-1.6 -ID:/GTK/include/gtk-2.0 -ID:/GTK/include/sigc++-2.0 -ID:/GTK/lib/sigc++-2.0/include -ID:/GTK/include/glib-2.0 -ID:/GTK/lib/glib-2.0/include -ID:/GTK/lib/gtk-2.0/include -ID:/GTK/include/cairomm-1.0 -ID:/GTK/include/pango-1.0 -ID:/GTK/include/cairo -ID:/GTK/include/freetype2 -ID:/GTK/include -ID:/GTK/include/atk-1.0  
gtk_lib=-user32 -Wl,-luuid -LD:/GTK/lib -lgtkmm-2.4 -lgdkmm-2.4 -latkmm-1.6 -lgtk-win32-2.0 -lpangomm-1.4 -lcairomm-1.0 -lglibmm-2.4 -lsigc-2.0 -lgdk-win32-2.0 -lgdi32 -limm32 -lshell32 -lole32 -latk-1.0 -lgdk_pixbuf-2.0 -lpangowin32-1.0 -lpangocairo-1.0 -lcairo -lpangoft2-1.0 -lfontconfig -lfreetype -lz -lpango-1.0 -lm -lgobject-2.0 -lgmodule-2.0 -lglib-2.0 -lintl -liconv  

FILES = $(wildcard $(FilePre)*.cc) $(wildcard $(FilePre)*.cpp) 
SOURCES = $(notdir $(FILES)) 
OBJS = $(patsubst %.cpp,%.o,$(SOURCES))
EXE = $(patsubst %.cpp,%.exe,$(SOURCES))

all: $(EXE)

$(EXE): $(SOURCES)
	@echo ========= building $^ ... =======
	$(CXX) $(CXXFLAGS) \
	       $(addprefix -I, $(INC))  $^ -o $@ \
	       $(addprefix -L, $(LIBPATH)) $(addprefix -l, $(LIBS)) 
	@echo ========= Finish. $^ =======
	mv $(EXE) Debug
	
#           $(SDLFLAGS)\
#	       $(SDLLIBS) \
#$(OBJS): %.o: %.cc

.PHONY:all clean test RM_O RM_EXE 
RM_O:
	-@rm -f $(OBJS) 2>/dev/null 2>/dev/null;

RM_EXE:			  
	-@rm -f Debug/$(EXE) 2>/dev/null 2>/dev/null;

clean:RM_O RM_EXE

test:
	@echo $(EXE)
	@echo $(OBJS) 

