#include <cstdio>
#include <fcntl.h>
#include <unistd.h>
#include <cstdlib>
extern int eof (int &stream);

int main( void )
{
    int filedes , len;
    char buffer[100];

    filedes = open( "file", O_RDONLY );
    if( filedes != -1 ) {
        while( ! eof( filedes ) ) {
            len = read( filedes , buffer, sizeof(buffer) - 1 );
            buffer[ len ] = '\0';
            printf( "%s", buffer );
        }
        close( filedes );
        
        return EXIT_SUCCESS;
    }
    
    return EXIT_FAILURE;
}
