#!/usr/bin/python
#coding:utf-8
#zz from http://www.exonsoft.com/~kochin/TinyP2P/tinyp2p.html
# tinyp2p.py 1.0 (documentation at http://freedom-to-tinker.com/tinyp2p.html)
# (C) 2004, E.W. Felten
# license: http://creativecommons.org/licenses/by-nc-sa/2.0


# Annotated by Kochin Chang, Jan. 2005


# Usage:
#   Server - python tinyp2p.py password server hostname  portnum [otherurl]
#   Client - python tinyp2p.py password client serverurl pattern
#                   ar[0]      ar[1]    ar[2]  ar[3]     ar[4]   ar[5]


import sys, os, SimpleXMLRPCServer, xmlrpclib, re, hmac
# Import libraries used in the program.
# sys : system variables and functions.
# os : portable OS dependent functionalities.
# SimpleXMLRPCServer : basic XML-RPC server framework.
# xmlrpclib : XML-RPC client support.
# re : regular expression support.
# hmac : RFC 2104 Keyed-Hashing Message Authentication.


ar,pw,res = (sys.argv,lambda u:hmac.new(sys.argv[1],u).hexdigest(),re.search)
# A multiple assignment.
# ar <- sys.argv : the argument list.
# pw <- lambda u:hmac.new(sys.argv[1],u).hexdigest() :
#   a function makes an HMAC digest from a URL.
#   INPUT: a string, u, which is a URL here.
#   OUTPUT: a hexdecimal HMAC digest.
#   DESCRIPTION:
#     1. Creates a HMAC object from the URL using network's password,
#        sys.argv[1], as the key.
#     2. Returns a hexdecimal digest of the HMAC object.
# res <- re.search : alias for the regular expression search function.

pxy,xs = (xmlrpclib.ServerProxy,SimpleXMLRPCServer.SimpleXMLRPCServer)
# A multiple assignment.
# pxy <- xmlrpclib.ServerProxy : alias for the ServerProxy class.
# xs <- SimpleXMLRPCServer.SimpleXMLRPCServer : alias for the SimpleXMLRPCServer class.

def ls(p=""):return filter(lambda n:(p=="")or res(p,n),os.listdir(os.getcwd()))
# a function lists directory entries.
# INPUT: a string, p, which is a regular expression pattern.
# OUTPUT: a list of directory entries matched the pattern.
# DESCRIPTION:
#   1. Creates a function using lambda expression that takes a pathname as its
#      parameter. The function returns true if the pattern is empty or the
#      pathname matches the pattern.
#   2. Finds out what is the current working directory.
#   3. Retrieves a list of directory entries of current working directory.
#   4. Filters the list using the lambda function defined.

if ar[2]!="client":
# Running in server mode...

  myU,prs,srv = ("http://"+ar[3]+":"+ar[4], ar[5:],lambda x:x.serve_forever())
  # A multiple assignment.
  # myU <- "http://"+ar[3]+":"+ar[4] : server's own URL.
  # prs <- ar[5:] : URL's of other servers in the network.
  # srv <- lambda x:x.serve_forever() :
  #   a function to start a SimpleXMLRPCServer.
  #   INPUT: a SimpleXMLRPCServer object, x.
  #   OUTPUT: (none)
  #   DESCRIPTION:
  #     Calls the server's serve_forever() method to start handling request.

  def pr(x=[]): return ([(y in prs) or prs.append(y) for y in x] or 1) and prs
  # a function returns the server list.
  # INPUT: a list, x, of servers' URLs to be added to the server list.
  # OUTPUT: the updated server list.
  # DESCRIPTION:
  #   1. For each URL in x, checks whether it's already in the server list.
  #      If it's not in the list, appends in onto the list.
  #   2. Returns the updated server list.

  def c(n): return ((lambda f: (f.read(), f.close()))(file(n)))[0]
  # a function returns content of the specified file.
  # INPUT: a string, n, which is a filename.
  # OUTPUT: the content of the file in a string.
  # DESCRIPTION:
  #   1. Creates a function using lambda expression that takes a file object, f,
  #      as its parameter. The function reads the content of the file, then
  #      closes it. The results of the read and close are put into a tuple, and
  #      the tuple is returned.
  #   2. Creates a file object with the filename. Passes it to the lambda
  #      function.
  #   3. Retrieves and returns the first item returned from the lambda function.

  f=lambda p,n,a:(p==pw(myU))and(((n==0)and pr(a))or((n==1)and [ls(a)])or c(a))
  #   a request handling function, depending on the mode, returns server list,
  #   directory entries, or content of a file.
  #   INPUT: a string, p, which is a hexdecimal HMAC digest.
  #          a mode number, n.
  #          if n is 0, a is a list of servers to be added to server list.
  #          if n is 1, a is a pattern string.
  #          if n is anything else, a is a filename.
  #   OUTPUT: if n is 0, returns the server list.
  #           if n is 1, returns directory entries match the pattern.
  #           if n is anything else, returns content of the file.
  #   DESCRIPTION:
  #     1. Verifies the password by comparing the HMAC digest received and the
  #        one created itself. Continues only when they match.
  #     2. If n is 0, calls pr() to add list, a, and returns the result.
  #        If n is 1, calls ls() to list entries match pattern a, and returns
  #        the result enclosed in a list.
  #        If n is any other value, retreives and return content of the file
  #        with filename specified in a.

  def aug(u): return ((u==myU) and pr()) or pr(pxy(u).f(pw(u),0,pr([myU])))
  # a function augments the network.
  # INPUT: a string, u, which is a URL.
  # OUTPUT: a list of URL's of servers in the network.
  # DESCRIPTION:
  #   1. If the URL, u, equals to my own URL, just returns the server list.
  #   2. Otherwise, creates a ServerProxy object for server u. Then calls its
  #      request handling function f with a HMAC digest, mode 0, and server
  #      list with myself added.
  #   3. Calls pr() with the result returned from server u to add them to my
  #      own list.
  #   4. Returns the new list.

  pr() and [aug(s) for s in aug(pr()[0])]
  # 1. Checks the server list is not empty.
  # 2. Takes the first server on the list. Asks that server to augment its
  #    server list with my URL.
  # 3. For each server on the returned list, asks it to add this server to its
  #    list.

  (lambda sv:sv.register_function(f,"f") or srv(sv))(xs((ar[3],int(ar[4]))))
  # Starts request processing.
  # 1. Defines a function with lambda expression that takes a SimpleXMLRPCServer
  #    object, registers request handling function, f, and starts the server.
  # 2. Creates a SimpleXMLRPCServer object using hostname (ar[3]) and portnum
  #    (ar[4]). Then feeds the object to the lambda function.

# Running in client mode...
for url in pxy(ar[3]).f(pw(ar[3]),0,[]):
# 1. Create a ServerProxy object using the serverurl (ar[3]).
# 2. Calls the remote server and retrieves a server list.
# 3. For each URL on the list, do the following:

  for fn in filter(lambda n:not n in ls(), (pxy(url).f(pw(url),1,ar[4]))[0]):
  # 1. Create a ServerProxy object using the URL.
  # 2. Calls the remote server to return a list of filenames matching the
  #    pattern (ar[4]).
  # 3. For each filename doesn't exist locally, do the following:

    (lambda fi:fi.write(pxy(url).f(pw(url),2,fn)) or fi.close())(file(fn,"wc"))
    # 1. Define a lambda function that takes a file object, calls remote server
    #    for the file content, then closes the file.
    # 2. Create a file object in write and binary mode with the filename. (I
    #    think the mode "wc" should be "wb".)
    # 3. Passes the file object to the lambda function.
