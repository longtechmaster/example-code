#!/bin/sh
protoc --proto_path=../service-auth/ --doc_out=../documents/. --doc_opt=html,service-auth.html \
-I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-auth/. \
--micro_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-auth/.  \
user.proto base.proto user_admin.proto user_role.proto user_member.proto service_other.proto
