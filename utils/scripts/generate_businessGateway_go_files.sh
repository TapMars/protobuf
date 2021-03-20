export GO111MODULE=on
go get google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"

#Go to protobuf directory
cd ../../

SERVICE_PATH="businessGateway/v1"

protoc --proto_path=. --go_out=gen/businessGateway-go --go_opt=paths=source_relative $SERVICE_PATH/businessGateway.proto
protoc --proto_path=. --go-grpc_out=gen/businessGateway-go --go-grpc_opt=paths=source_relative $SERVICE_PATH/businessGateway.proto