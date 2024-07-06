#### install golang packages
```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0
```

#### setting up golang path for protobuf compiler
```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```