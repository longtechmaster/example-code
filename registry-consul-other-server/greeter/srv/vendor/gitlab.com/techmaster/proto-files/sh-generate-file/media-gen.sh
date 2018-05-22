#!/bin/sh
protoc --proto_path=../service-media/ --doc_out=../documents/. --doc_opt=html,service-media.html \
-I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-media/. \
--micro_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-media/.  \
media.proto base.proto media_file.proto media_image.proto