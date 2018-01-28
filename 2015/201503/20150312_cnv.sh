#!/bin/bash
#for file in $(find . -name "*.dds")
#do 
#	echo "Processing $file"
#	nvdecompress $file $file.tga
#done
find . -name "*.dds" -exec nvdecompress {} {}.tga \;
