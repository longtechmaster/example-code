# Proto files

Lưu các file Proto dành cho dự án.

Mỗi service sẽ có 1 thư mục để lưu file proto dành cho service đó. 
Service nào sử dụng file proto thì tự generate ra code tương ứng với ngôn ngữ sử dụng.

## Cài đặt protoc

[http://google.github.io/proto-lens/installing-protoc.html](http://google.github.io/proto-lens/installing-protoc.html) (Bản đang sử dụng là 3.5.1)

- Để generate ra code Golang thì cần thêm `protoc-gen-go`: 

```
go get github.com/golang/protobuf/protoc-gen-go
```
- Để generate ra code Micro thì cần thêm `protoc-gen-micro`: 

```
go get github.com/micro/protoc-gen-micro
```

 - Để có thêm anotation cho message thì cần thêm `gogo-protobuf`:
  
```
 go get github.com/gogo/protobuf/protoc-gen-gogofast
 go get github.com/gogo/protobuf/protoc-gen-gogofaster
 go get github.com/gogo/protobuf/protoc-gen-gogoslick
```


## Generate file pb.go và micro.go

# Sử dụng gogo protobuf
## Dùng lựa chọn --gofast_out đơn giản nhất không cho phép extension
```shell
protoc --proto_path=service-blog/proto --gofast_out=service-blog/gogoout/. --micro_out=service-blog/gogoout/. base.proto blog_author.proto blog_category.proto blog_comment.proto blog_post.proto blog_tag.proto blog.proto
```
## Generate file với sh-file

```
cd proto-files
./auth-gen.sh
./blog-gen.sh
./media-gen.sh
...

```




