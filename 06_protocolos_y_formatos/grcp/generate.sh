#bin/bash

protoc calcpb/calculator.proto --go_out=module=example.com/grcp-calcpb,plugins=grpc:./calcpb
