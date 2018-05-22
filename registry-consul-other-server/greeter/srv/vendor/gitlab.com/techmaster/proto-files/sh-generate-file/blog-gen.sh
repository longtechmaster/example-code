#!/bin/sh
protoc --proto_path=../service-blog/ --doc_out=../documents/. --doc_opt=html,service-blog.html \
-I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-blog/. \
--micro_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:../service-blog/.  \
blog.proto base.proto blog_author.proto blog_category.proto blog_comment.proto blog_post.proto blog_tag.proto
