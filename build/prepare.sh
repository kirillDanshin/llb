#!/bin/sh
mkdir ./$BUILDDIR/builds/
for dir in $(find ./$BUILDDIR/*/* -type d)
do
	# {{.Arch}}_{{.Platform}} pattern
	tar -czf ./$BUILDDIR/builds/$(basename $(dirname $dir))_$(basename $dir).tar.gz $dir
done
