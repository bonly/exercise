#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

int main( int argc, char **argv ) {
    pid_t child = 0;
    int pipes[2];

    if( pipe( pipes ) != 0 ) {
        fprintf( stderr, "cannot create pipe" );
        return 1;
    }

    child = fork();

    if (child < 0) {
        fprintf( stderr, "process failed to fork\n" );
        return 1;
    }
    if (child == 0) {
        close( pipes[1] );
        // read from pipes[0] here
        //wait(NULL);
    }
    else {
        dup2( pipes[1], 0 );        // 0 is stdout
        close( pipes[0] );
        execl( "/bin/ls", "ls");
    }
    return 0;
}

