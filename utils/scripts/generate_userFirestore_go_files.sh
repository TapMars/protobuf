export GO111MODULE=on
go get google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"

#Go to protobuf directory
cd ../../

SERVICE_PATH="userFirestore/v1"

#protoc --proto_path=proto --go_out=gen/go --go_opt=paths=source_relative $USER_PATH/product_search/v1/product_search.proto
#protoc --proto_path=proto --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative $USER_PATH/product_search/v1/product_search.proto

protoc --proto_path=. --go_out=gen/userFirestore-go --go_opt=paths=source_relative $SERVICE_PATH/userFirestore.proto
protoc --proto_path=. --go-grpc_out=gen/userFirestore-go --go-grpc_opt=paths=source_relative $SERVICE_PATH/userFirestore.proto

#protoc --proto_path=proto --go_out=gen/go --go_opt=paths=source_relative $USER_PATH/gateway/v1/gateway.proto
#protoc --proto_path=proto --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative $USER_PATH/gateway/v1/gateway.proto