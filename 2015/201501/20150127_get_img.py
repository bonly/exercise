#!/usr/bin/env python
# coding: utf-8

# Download a random picture from Google image search.
#
# Usage:
# $ fetch_google_image.py cat cute   # Download a cute cat picture

import os
import sys
import urllib
import urllib2
import json
import random
import imghdr

if len(sys.argv) <= 1:
  print('Usage:')
  print('python fetch_google_image.py cat cute')
  exit()

q = ''

for arg in sys.argv[1:]:
  q += urllib.quote(arg) + '+'

f = urllib2.urlopen('http://ajax.googleapis.com/ajax/services/search/images?q=' + q + '&v=1.0&rsz=large&start=1')
data = json.load(f)
f.close()

results = data['responseData']['results']
url = results[random.randint(0, len(results) - 1)]['url']
urllib.urlretrieve(url, './image')

imagetype = imghdr.what('./image')
if not(type(imagetype) is None):
  os.rename('./image', './image.' + imagetype)