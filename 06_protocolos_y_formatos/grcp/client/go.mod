module example.com/grcp-cliente

go 1.15

require (
	example.com/grcp-calcpb v0.0.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace example.com/grcp-calcpb => ../calcpb
