export OPT=~/opt
PATH=.:/bin:/usr/bin:/sbin:/usr/sbin:/usr/local/bin:/usr/local/sbin:$OPT/bin:

#export XDG_CONFIG_HOME=$HOME/.config

#git
#GIT=$OPT/git
#PATH=$GIT/bin:$GIT/share:$PATH

#mercurial hg
#python setup.py install --force --home=$HOME
#HG=$OPT/mercurial
#PATH=$HG/bin:$PATH
#LD_LIBRARY_PATH=$HG/lib:$LD_LIBRARY_PATH

#graphviz/bin
#GR=$OPT/graphviz
#PATH=$GR/bin:$GR/share:$PATH
#LD_LIBRARY_PATH=$GR/lib:$LD_LIBRARY_PATH

#node.js
#npm config set prefix /home/opt/node_modules
#NODE_JS=$OPT/node.js
#MY_NODE=$OPT/node_modules
#NODE_PATH=$OPT/node_modules
#PATH=$NODE_JS/bin:$NODE_JS/share:$MY_NODE/bin:$MY_NODE/.bin:$PATH
#LD_LIBRARY_PATH=$NODE_JS/lib:$LD_LIBRARY_PATH

#peazip
PATH=$OPT/peazip:$PATH

#ruby
PATH=$HOME/.gem/ruby/2.3.0/bin/:$PATH

#xmms2
#XMMS2=$OPT/xmms
#PATH=$XMMS2/bin:$XMMS2/share:$PATH
#LD_LIBRARY_PATH=$XMMS2/lib:$XMMS2/lib/xmms2:$LD_LIBRARY_PATH

#bluefish
BLUEFISH=$OPT/bluefish
PATH=$BLUEFISH/bin:$BLUEFISH/share:$PATH
LD_LIBRARY_PATH=$BLUEFISH/lib:$LD_LIBRARY_PATH

#netease music
NMUSIC=$OPT/netease
PATH=$NMUSIC/usr/bin:$NMUSIC/usr/lib/netease-cloud-music:$NMUSIC/usr/share/doc
LD_LIBRARY_PATH=$NMUSIC/usr/lib/netease-cloud-music:$LD_LIBRARY_PATH

#wine
WINE_ROOT=$OPT/wine
PATH=$WINE_ROOT/usr/bin:$PATH
LD_LIBRARY_PATH=$WINE_ROOT/usr/lib:$LD_LIBRARY_PATH

#diffuse
DIFFUSE=$OPT/diffuse
PATH=$DIFFUSE/bin:$PATH

#gimp
#GIMP=$OPT/gimp
#PATH=$GIMP/bin:$PATH
#LD_LIBRARY_PATH=$GIMP/lib:$LD_LIBRARY_PATH

#bouml
#export BOUML_CHARSET=UTF8
#BOUML_ID

#chrome
export CHROME_PATH=/opt/google/chrome

#tmux
TMUX_PATH=$OPT/tmux
PATH=$TMUX_PATH/bin:$PATH

#java
JAVA_HOME=$OPT/jdk
PATH=$OPT/jdk/bin:$PATH 
export JAVA_HOME
#export _JAVA_OPTIONS='-Dawt.useSystemAAFontSettings=lcd'

#aria2
ARIA2=$OPT/aria
PATH=$ARIA2/bin:$ARIA2/share:$PATH

#openvpn
#OPENVPN=$OPT/openvpn
#PATH=$OPENVPN/sbin:$OPENVPN/share:$PATH
#LD_LIBRARY_PATH=$OPENVPN/lib/openvpn/plugins/:$LD_LIBRARY_PATH

#wireshark
#WIRESHARK=$OPT/wireshark
#PATH=$WIRESHARK/bin:$WIRESHARK/share:$PATH
#LD_LIBRARY_PATH=$WIRESHARK/lib:$LD_LIBRARY_PATH

#php
PHP=$OPT/php
PATH=$PHP/bin:$PHP/sbin:$PHP/php:$PATH
LD_LIBRARY_PATH=$PHP/lib:$PHP/lib/php:$LD_LIBRARY_PATH

#snav
#SNAV=$OPT/snav
#PATH=$SNAV/bin:$SNAV/libexec/snavigator:$PATH
#LD_LIBRARY_PATH=$SNAV/lib:$LD_LIBRARY_PATH
#alias snav=snavigator

#vlc
#VLC=$OPT/vlc
#PATH=$VLC/bin:$VLC/share:$PATH
#LD_LIBRARY_PATH=$VLC/lib:$LD_LIBRARY_PATH

#nginx
NGINX=$OPT/nginx
PATH=$NGINX/sbin:$PATH
LD_LIBRARY_PATH=$NGINX/lib:$LD_LIBRARY_PATH

#meld
#MELD=$OPT/meld-1.6.0
#PATH=$MELD/bin:$PATH

#p4v
#P4V=$OPT/p4v
#PATH=$P4V/bin:$PATH
LD_LIBRARY_PATH=$P4V/lib:$P4V/lib/p4v/qt4/lib:$LD_LIBRARY_PATH

#cflow
#PATH=/home/opt/cflow/bin:$PATH

#scrot
#PATH=$OPT/scrot/bin:$OPT/scrot/share:$PATH

#wayland
#WLD=$OPT/wayland
##PATH=$WLD/share:$WLD/bin:$WLD/libexec:$WLD/client:$PATH
##LD_LIBRARY_PATH=$WLD/lib:$WLD/syslib:$LD_LIBRARY_PATH
#export XDG_RUNTIME_DIR=~/weston

#adb
ADB=$OPT/adb/usr
PATH=$ADB/bin:$ADB/share:$PATH

#qt-creator
#QT=$OPT/qt
#LD_LIBRARY_PATH=$QT/Tools/QtCreator/lib/qtcreator/plugins/QtProject:$LD_LIBRARY_PATH
#PATH=$QT/Tools/QtCreator/bin/:$QT/5.2.0/gcc/bin:$PATH

#golang
#export GOOS=linux/freebsd/darwin/nacl
#export GOARCH=amd64/386/arm
export GOHOSTOS=linux
export GOHOSTARCH=amd64
export GOROOT=$OPT/go
export GOBIN=$GOROOT/bin
export GOPATH=$OPT/go.my
export GOROOT_BOOTSTRAP=/usr/lib/go
PATH=$GOPATH/bin:$GOROOT/bin:$PATH

#mongodb
PATH=$OPT/mongodb/bin:$PATH

#mysql-workbench
MYBEN=$OPT/mysql_workbench
PATH=$MYBEN/usr/bin:$PATH
LD_LIBRARY_PATH=$MYBEN/usr/lib/mysql-workbench:$LD_LIBRARY_PATH

#xbindkeys
PATH=$OPT/xbindkeys/bin:$OPT/xbindkeys/share:$PATH

#cscope
PATH=$OPT/cscope/bin:$OPT/cscope/share:$PATH

#understand
export STIHOME=$OPT/scitools
PATH=$OPT/scitools/bin:$PATH

#privoxy
export PRIVOXY=$OPT/privoxy
PATH=$PRIVOXY/sbin:$PATH

#gradle
#export GRADLE_OPTS=
#export JAVA_OPTS=
export GRADLE_HOME=$OPT/gradle
export PATH=$GRADLE_HOME/bin:$PATH

#ctags
PATH=$OPT/ctags/bin:$OPT/ctags/share:$PATH


#scanmem
PATH=$OPT/scanmem/bin:$PATH
LD_LIBRARY_PATH=$OPT/scanmem/lib:$LD_LIBRARY_PATH

#bazel
PATH=$OPT/bazel/lib/bazel/bin:$PATH

#neovim
PATH=$OPT/neovim/bin:$PATH


export LD_LIBRARY_PATH=.:$LD_LIBRARY_PATH
export PATH=.:$PATH 

v(){
   nvim "$@"
}

m() {
    #LD_LIBRARY_PATH=/home/opt/xmms/lib:$LD_LIBRARY_PATH;xmms2  "$@"
    mocp  "$@"
}
ms() {
    #LD_LIBRARY_PATH=/home/opt/xmms/lib:$LD_LIBRARY_PATH;xmms2  status
    mocp -s
}
mr(){
    LD_LIBRARY_PATH=/home/opt/xmms/lib:$LD_LIBRARY_PATH;xmms2d  --yes-run-as-root
}
mn(){
    #LD_LIBRARY_PATH=/home/opt/xmms/lib:$LD_LIBRARY_PATH;xmms2d  --yes-run-as-root
    mocp -p
}

# Set colors for man pages
man() {
        env \
                LESS_TERMCAP_mb=$(printf "\e[1;31m") \
                LESS_TERMCAP_md=$(printf "\e[1;31m") \
                LESS_TERMCAP_me=$(printf "\e[0m") \
                LESS_TERMCAP_se=$(printf "\e[0m") \
                LESS_TERMCAP_so=$(printf "\e[1;44;33m") \
                LESS_TERMCAP_ue=$(printf "\e[0m") \
                LESS_TERMCAP_us=$(printf "\e[1;32m") \
                        man "$@"
}

source $OPT/bin/.git-completion.bash

alias docker='sudo docker '
[ -f ~/bin/bashrc_docker ] && . ~/bin/bashrc_docker

vi(){
    eval pth=$PWD
    GOPATH=${pth%src*}:${GOPATH} vim "$@"
}

