module example.com/grcp-servidor

go 1.15

require (
    google.golang.org/grpc v1.32.0
    example.com/grcp-calcpb v0.0.0
)

replace example.com/grcp-calcpb => ../calcpb