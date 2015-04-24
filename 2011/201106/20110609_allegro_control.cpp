

#include <allegro.h>

int x = 10;
int y = 10;

int main(){
 
    allegro_init();
    install_keyboard();
    set_gfx_mode( GFX_AUTODETECT, 640, 480, 0, 0);
    
    while ( !key[KEY_ESC] ){
    
        clear_keybuf();
        
        acquire_screen();
        
        textout_ex( screen, font, " ", x, y, makecol( 0, 0, 0), makecol( 0, 0, 0) );
        
        if (key[KEY_UP]) --y;        
        else if (key[KEY_DOWN]) ++y;    
        else if (key[KEY_RIGHT]) ++x;
        else if (key[KEY_LEFT]) --x;

        textout_ex( screen, font, "@", x, y, makecol( 255, 0, 0), makecol( 0, 0, 0) );
        
        release_screen();
        
        rest(50);

    }    
    
    return 0;
    
}   
END_OF_MAIN(); 

/*
      KEY_A - KEY_Z,
      KEY_0 - KEY_9,
      KEY_0_PAD - KEY_9_PAD,
      KEY_F1 - KEY_F12,

      KEY_ESC, KEY_TILDE, KEY_MINUS, KEY_EQUALS,
      KEY_BACKSPACE, KEY_TAB, KEY_OPENBRACE, KEY_CLOSEBRACE,
      KEY_ENTER, KEY_COLON, KEY_QUOTE, KEY_BACKSLASH,
      KEY_BACKSLASH2, KEY_COMMA, KEY_STOP, KEY_SLASH,
      KEY_SPACE,

      KEY_INSERT, KEY_DEL, KEY_HOME, KEY_END, KEY_PGUP,
      KEY_PGDN, KEY_LEFT, KEY_RIGHT, KEY_UP, KEY_DOWN,

      KEY_SLASH_PAD, KEY_ASTERISK, KEY_MINUS_PAD,
      KEY_PLUS_PAD, KEY_DEL_PAD, KEY_ENTER_PAD,

      KEY_PRTSCR, KEY_PAUSE,

      KEY_ABNT_C1, KEY_YEN, KEY_KANA, KEY_CONVERT, KEY_NOCONVERT,
      KEY_AT, KEY_CIRCUMFLEX, KEY_COLON2, KEY_KANJI,

      KEY_LSHIFT, KEY_RSHIFT,
      KEY_LCONTROL, KEY_RCONTROL,
      KEY_ALT, KEY_ALTGR,
      KEY_LWIN, KEY_RWIN, KEY_MENU,
      KEY_SCRLOCK, KEY_NUMLOCK, KEY_CAPSLOCK

      KEY_EQUALS_PAD, KEY_BACKQUOTE, KEY_SEMICOLON, KEY_COMMAND
http://www.cppgameprogramming.com/cgi/nav.cgi?page=allegio
*/
