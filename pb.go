//go:generate  protoc --go_out=plugins=grpc:. q.proto
//go:generate  protoc --go_out=plugins=grpc:. qubit.proto
//go:generate  protoc --go_out=plugins=grpc:. math.proto

package q
