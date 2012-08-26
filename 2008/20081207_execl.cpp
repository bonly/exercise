#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <boost/thread.hpp>

int main( int argc, char **argv ) {
    pid_t child = 0;
    child = fork();
    if (child < 0) {
        fprintf( stderr, "process failed to fork\n" );
        return 1;
    }
    if (child == 0) {
        //wait(NULL);
        boost::thread::sleep(boost::get_system_time()+boost::posix_time::milliseconds(10000)); 
    }
    else {
        execl( "/bin/ls", "ls");
    }
    return 0;
} 



