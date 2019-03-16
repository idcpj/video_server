#! /bin/bash

## BUild web UI
cd  $GOPATH/src/video_server/web
go install

cp $GOPATH/bin/web $GOPATH/bin/video_server_web_ui/web
cp -R $GOPATH/src/video_server/templates $GOPATH/bin/video_server_web_ui/

