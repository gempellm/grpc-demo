.PHONY: gen
gen:
	cd api/demo_service && protoc --go_out=../../pkg/demo_service --go_opt=paths=source_relative --go-grpc_out=../../pkg/demo_service --go-grpc_opt=paths=source_relative ./demo_service.proto

.PHONE: run
run:
	go run internal/demo_server/main.go && go run internal/demo_client/main.go