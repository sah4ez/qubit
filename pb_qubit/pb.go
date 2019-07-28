//go:generate  protoc --go_out=plugins=grpc:. ./math/math.proto
//go:generate  protoc --go_out=plugins=grpc:. q.proto

package pb_qubit
