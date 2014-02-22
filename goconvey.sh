#!/bin/sh
if [ "$1" = "stop" ]
then
    echo "Stopping GoConvey"
    killall goconvey
else
    echo "Starting GoConvey"
    $GOPATH/bin/goconvey &
    open http://localhost:8080
fi
