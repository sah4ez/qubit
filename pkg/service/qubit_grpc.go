package service

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/sah4ez/qubit/pb_qubit"
)

type GrpcQubitServer struct {
	NumberOfBit       kitgrpc.Handler
	IsZero            kitgrpc.Handler
	IsOne             kitgrpc.Handler
	InnerProduct      kitgrpc.Handler
	OuterProduct      kitgrpc.Handler
	Fidelity          kitgrpc.Handler
	TraceDistance     kitgrpc.Handler
	Equals            kitgrpc.Handler
	TensorProduct     kitgrpc.Handler
	Apply             kitgrpc.Handler
	Normalize         kitgrpc.Handler
	Amplitude         kitgrpc.Handler
	Probability       kitgrpc.Handler
	Measure           kitgrpc.Handler
	ProbabilityZeroAt kitgrpc.Handler
	ProbabilityOneAt  kitgrpc.Handler
	MeasureAt         kitgrpc.Handler
}

func (q *GrpcQubitServer) NumberOfBit(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.NumberOfBitResponse, error) {
	_, rep, err := q.NumberOfBit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.NumberOfBitResponse), nil
}

func (q *GrpcQubitServer) IsZero(ctx context.Context, req *pb_qubit.EpsRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.IsZero.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) IsOne(ctx context.Context, req *pb_qubit.EpsRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.IsOne.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) InnerProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.Complex128Response, error) {
	_, rep, err := q.InnerProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.Complex128Response), nil
}

func (q *GrpcQubitServer) OuterProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.MatrixResponse, error) {
	_, rep, err := q.OuterProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.MatrixResponse), nil
}

func (q *GrpcQubitServer) OuterProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.MatrixResponse, error) {
	_, rep, err := q.OuterProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.MatrixResponse), nil
}

func (q *GrpcQubitServer) Fidelity(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.DoubleResponse, error) {
	_, rep, err := q.Fidelity.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.DoubleResponse), nil
}

func (q *GrpcQubitServer) TraceDistance(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.DoubleResponse, error) {
	_, rep, err := q.TraceDistance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.DoubleResponse), nil
}

func (q *GrpcQubitServer) Equals(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.Equals.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) TensorProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.TensorProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Apply(ctx context.Context, req *pb_qubit.MatrixRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.Apply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Normalize(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.QResponse, error) {
	_, rep, err := q.Normalize.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Amplitude(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.AmplitudeResponse, error) {
	_, rep, err := q.Amplitude.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.AmplitudeResponse), nil
}

func (q *GrpcQubitServer) Probability(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.ProbabilityResponse, error) {
	_, rep, err := q.Probability.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityResponse), nil
}

func (q *GrpcQubitServer) Measure(ctx context.Context, req *pb_qubit.MeasureRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.Measure.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) ProbabilityZeroAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.ProbabilityAtResponse, error) {
	_, rep, err := q.ProbabilityZeroAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityAtResponse), nil
}

func (q *GrpcQubitServer) ProbabilityOneAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.ProbabilityAtResponse, error) {
	_, rep, err := q.ProbabilityOneAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityAtResponse), nil
}

func (q *GrpcQubitServer) MeasureAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.MeasureAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}
