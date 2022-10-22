#!/bin/sh

webdir=$(PWD)/../hansip-web
webserverdir=$(PWD)/../hansip-web-server

echo "Building web sources..."
cd $webdir && npm run build

echo "Copying web assets..."
cd $webdir \
    && rm -rf $webserverdir/assets/web \
    && cp -rf $webdir/dist $webserverdir/assets/web
cd $webserverdir && go run build
echo "Build done.\n"