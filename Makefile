generate_grpc_code:
	protoc --plugin=protoc-gen-go=$(which protoc-gen-go) \
    --go_out=invoicer \
    --go_opt=paths=source_relative \
    --go-grpc_out=invoicer \
    --go-grpc_opt=paths=source_relative \
    invoicer.proto