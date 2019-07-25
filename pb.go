//go:generate  protoc --go_out=plugins=grpc:. ./math/math.proto
//go:generate  protoc --go_out=plugins=grpc:. qubit.proto
//go:generate  protoc --go_out=plugins=grpc:. q.proto

package qubit
