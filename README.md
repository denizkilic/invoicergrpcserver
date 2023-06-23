## INVOICER GRPC SERVER

Download protobuf compiler on https://grpc.io/docs/protoc-installation/

```
brew install protobuf
protoc --version

go install -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
export PATH=$PATH:$(go env GOPATH)/bin
which protoc-gen-go


go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH=$PATH:$(go env GOPATH)/bin

then run this command;
 protoc --plugin=protoc-gen-go=$(which protoc-gen-go) \       
    --go_out=invoicer \
    --go_opt=paths=source_relative \
    --go-grpc_out=invoicer \
    --go-grpc_opt=paths=source_relative \
    invoicer.proto

this command will generate files under invoicer folder.

        go run main.go

I test the server with gRPCurl 
        brew install grpcurl
        

```

