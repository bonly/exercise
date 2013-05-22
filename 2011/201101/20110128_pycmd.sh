#!/bin/sh
ARGS=""
for i in "$@" ; do
        ARGS="$ARGS '$i'"
done

#echo $ARGS
#echo `basename $0`
export LD_LIBRARY_PATH=$HOME/opt/lib:$HOME/opt/maria/lib:$LD_LIBRARY_PATH
exec $HOME/opt/bin/python  /wlmz/tlbb/scripts/py/`basename $0` $ARGS

