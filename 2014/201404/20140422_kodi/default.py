# -*- coding: utf8 -*-
from ctypes import cdll
import xbmcgui
import xbmc
# import xbmcplugin
import xbmcaddon
import os
import sys


if __name__ == '__main__':
    # print sys.path;
    # print "hello";
    # xbmcgui.Dialog().notification("ok","path:".join(sys.path))
    # xbmcgui.Dialog().ok("ok","".join(sys.path))
    # xbmc.log(msg="\n".join(sys.path));
    addon = xbmcaddon.Addon();
    so_file = addon.getAddonInfo('path')+("/libadd.so");
    xbmc.log(so_file);


    lib = cdll.LoadLibrary(so_file);
    result = lib.add(2, 3);
    xbmc.log(msg="ret:"+str(result));
    xbmcgui.Dialog().ok("计算结果是:", str(result));


