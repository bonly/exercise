/**
# Function: transmit euid and egid to other scripts
#   since shell/python/... scripts can't get suid permission in Linux
#   usage: transeuid xxx.sh par1 par2 par3
#          xxx.sh will get the euid and egid from transeuid
# ******************************************************************** */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#define BUFFSIZE 1024

/*
 * usually euid is the uid who run the program
 * but when stick is setted to the program
 * euid is the uid or the program's owner
 */
int main(int argc, char *argv[]) {
    char *cmd = (char*)malloc(BUFFSIZE);
    // set uid and gid to euid and egid
    setuid(geteuid());
    setgid(getegid());
    cmd = argv[1];
    int i = 0;
    for(i = 0;i < argc - 1;i++) {
        argv[i] = argv[i+1];
    }
    argv[argc-1] = NULL;
    // search $PATH find this cmd and run it with pars:argv
    if (execvp(cmd, argv)) {
        printf("error");
        free(cmd);
        exit(1);
    }
    free(cmd);
}
