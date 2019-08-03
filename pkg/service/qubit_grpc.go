package service

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/sah4ez/qubit/pb_qubit"
	pb "github.com/sah4ez/qubit/pb_qubit"
)

var _ pb.QQubitServer = &GrpcQubitServer{}

type GrpcQubitServer struct {
	GRPCNumberOfBit       kitgrpc.Handler
	GRPCIsZero            kitgrpc.Handler
	GRPCIsOne             kitgrpc.Handler
	GRPCInnerProduct      kitgrpc.Handler
	GRPCOuterProduct      kitgrpc.Handler
	GRPCFidelity          kitgrpc.Handler
	GRPCTraceDistance     kitgrpc.Handler
	GRPCEquals            kitgrpc.Handler
	GRPCTensorProduct     kitgrpc.Handler
	GRPCApply             kitgrpc.Handler
	GRPCNormalize         kitgrpc.Handler
	GRPCAmplitude         kitgrpc.Handler
	GRPCProbability       kitgrpc.Handler
	GRPCMeasure           kitgrpc.Handler
	GRPCProbabilityZeroAt kitgrpc.Handler
	GRPCProbabilityOneAt  kitgrpc.Handler
	GRPCMeasureAt         kitgrpc.Handler
}

func (q *GrpcQubitServer) NumberOfBit(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.NumberOfBitResponse, error) {
	_, rep, err := q.GRPCNumberOfBit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.NumberOfBitResponse), nil
}

func (q *GrpcQubitServer) IsZero(ctx context.Context, req *pb_qubit.EpsRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.GRPCIsZero.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) IsOne(ctx context.Context, req *pb_qubit.EpsRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.GRPCIsOne.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) InnerProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.Complex128Response, error) {
	_, rep, err := q.GRPCInnerProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.Complex128Response), nil
}

func (q *GrpcQubitServer) OuterProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.MatrixResponse, error) {
	_, rep, err := q.GRPCOuterProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.MatrixResponse), nil
}

func (q *GrpcQubitServer) Fidelity(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.DoubleResponse, error) {
	_, rep, err := q.GRPCFidelity.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.DoubleResponse), nil
}

func (q *GrpcQubitServer) TraceDistance(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.DoubleResponse, error) {
	_, rep, err := q.GRPCTraceDistance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.DoubleResponse), nil
}

func (q *GrpcQubitServer) Equals(ctx context.Context, req *pb_qubit.EqualsRequest) (*pb_qubit.BoolResponse, error) {
	_, rep, err := q.GRPCEquals.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.BoolResponse), nil
}

func (q *GrpcQubitServer) TensorProduct(ctx context.Context, req *pb_qubit.QRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.GRPCTensorProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Apply(ctx context.Context, req *pb_qubit.MatrixRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.GRPCApply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Normalize(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.QResponse, error) {
	_, rep, err := q.GRPCNormalize.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) Amplitude(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.AmplitudeResponse, error) {
	_, rep, err := q.GRPCAmplitude.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.AmplitudeResponse), nil
}

func (q *GrpcQubitServer) Probability(ctx context.Context, req *pb_qubit.Empty) (*pb_qubit.ProbabilityResponse, error) {
	_, rep, err := q.GRPCProbability.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityResponse), nil
}

func (q *GrpcQubitServer) Measure(ctx context.Context, req *pb_qubit.MeasureRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.GRPCMeasure.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func (q *GrpcQubitServer) ProbabilityZeroAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.ProbabilityAtResponse, error) {
	_, rep, err := q.GRPCProbabilityZeroAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityAtResponse), nil
}

func (q *GrpcQubitServer) ProbabilityOneAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.ProbabilityAtResponse, error) {
	_, rep, err := q.GRPCProbabilityOneAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.ProbabilityAtResponse), nil
}

func (q *GrpcQubitServer) MeasureAt(ctx context.Context, req *pb_qubit.BitRequest) (*pb_qubit.QResponse, error) {
	_, rep, err := q.GRPCMeasureAt.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb_qubit.QResponse), nil
}

func NewGrpcQubitServer(ep QubitSet) *GrpcQubitServer {
	options := []kitgrpc.ServerOption{}

	return &GrpcQubitServer{
		GRPCNumberOfBit: kitgrpc.NewServer(
			ep.NumberOfBitEndpoint,
			decodeNumberOfBitRequest,
			encodeNumberOfBitResponse,
			options...,
		),
		GRPCIsZero: kitgrpc.NewServer(
			ep.IsZeroEndpoint,
			decodeIsZeroRequest,
			encodeIsZeroResponse,
			options...,
		),
		GRPCIsOne: kitgrpc.NewServer(
			ep.IsOneEndpoint,
			decodeIsOneRequest,
			encodeIsOneResponse,
			options...,
		),
		GRPCInnerProduct: kitgrpc.NewServer(
			ep.InnerProductEndpoint,
			decodeInnerProductRequest,
			encodeInnerProductResponse,
			options...,
		),
		GRPCOuterProduct: kitgrpc.NewServer(
			ep.OuterProductEndpoint,
			decodeOuterProductRequest,
			encodeOuterProductResponse,
			options...,
		),
		GRPCFidelity: kitgrpc.NewServer(
			ep.FidelityEndpoint,
			decodeFidelityRequest,
			encodeFidelityResponse,
			options...,
		),
		GRPCTraceDistance: kitgrpc.NewServer(
			ep.TraceDistanceEndpoint,
			decodeTraceDistanceRequest,
			encodeTraceDistanceResponse,
			options...,
		),
		GRPCEquals: kitgrpc.NewServer(
			ep.EqualsEndpoint,
			decodeEqualsRequest,
			encodeEqualsResponse,
			options...,
		),
		GRPCTensorProduct: kitgrpc.NewServer(
			ep.TensorProductEndpoint,
			decodeTensorProductRequest,
			encodeTensorProductResponse,
			options...,
		),
		GRPCApply: kitgrpc.NewServer(
			ep.ApplyEndpoint,
			decodeApplyRequest,
			encodeApplyResponse,
			options...,
		),
		GRPCNormalize: kitgrpc.NewServer(
			ep.NormalizeEndpoint,
			decodeNormalizeRequest,
			encodeNormalizeResponse,
			options...,
		),
		GRPCAmplitude: kitgrpc.NewServer(
			ep.AmplitudeEndpoint,
			decodeAmplitudeRequest,
			encodeAmplitudeResponse,
			options...,
		),
		GRPCProbability: kitgrpc.NewServer(
			ep.ProbabilityEndpoint,
			decodeProbabilityRequest,
			encodeProbabilityResponse,
			options...,
		),
		GRPCMeasure: kitgrpc.NewServer(
			ep.MeasureEndpoint,
			decodeMeasureRequest,
			encodeMeasureResponse,
			options...,
		),
		GRPCProbabilityZeroAt: kitgrpc.NewServer(
			ep.ProbabilityZeroAtEndpoint,
			decodeProbabilityZeroAtRequest,
			encodeProbabilityZeroAtResponse,
			options...,
		),
		GRPCProbabilityOneAt: kitgrpc.NewServer(
			ep.ProbabilityOneAtEndpoint,
			decodeProbabilityOneAtRequest,
			encodeProbabilityOneAtResponse,
			options...,
		),
		GRPCMeasureAt: kitgrpc.NewServer(
			ep.MeasureAtEndpoint,
			decodeMeasureAtRequest,
			encodeMeasureAtResponse,
			options...,
		),
	}
}

func decodeNumberOfBitRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeNumberOfBitResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeIsZeroRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeIsZeroResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeIsOneRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeIsOneResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeInnerProductRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeInnerProductResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeOuterProductRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeOuterProductResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeFidelityRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeFidelityResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeTraceDistanceRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeTraceDistanceResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeEqualsRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeEqualsResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeTensorProductRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeTensorProductResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeApplyRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeApplyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeNormalizeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeNormalizeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeAmplitudeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeAmplitudeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeProbabilityRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeProbabilityResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeMeasureRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeMeasureResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeProbabilityZeroAtRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeProbabilityZeroAtResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeProbabilityOneAtRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeProbabilityOneAtResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeMeasureAtRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return nil, nil
}

func encodeMeasureAtResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}
