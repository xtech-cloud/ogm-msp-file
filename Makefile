.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/file/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/file/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/file/bucket.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/file/object.proto
