export GO111MODULE=on
go get google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"

#Go to protobuf directory
cd ../../

SERVICE_PATH="models/v1"

protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/address.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/day.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/happy_hour.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/item_type.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/operational_hours.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/success_response.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/void_request.proto
protoc --proto_path=. --go_out=gen/models-go --go_opt=paths=source_relative $SERVICE_PATH/firestore_interface.proto