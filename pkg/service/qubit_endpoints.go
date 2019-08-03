package service

import (
	"github.com/go-kit/kit/endpoint"
)

type QubitSet struct {
	NumberOfBitEndpoint       endpoint.Endpoint
	IsZeroEndpoint            endpoint.Endpoint
	IsOneEndpoint             endpoint.Endpoint
	InnerProductEndpoint      endpoint.Endpoint
	OuterProductEndpoint      endpoint.Endpoint
	FidelityEndpoint          endpoint.Endpoint
	TraceDistanceEndpoint     endpoint.Endpoint
	EqualsEndpoint            endpoint.Endpoint
	TensorProductEndpoint     endpoint.Endpoint
	ApplyEndpoint             endpoint.Endpoint
	NormalizeEndpoint         endpoint.Endpoint
	AmplitudeEndpoint         endpoint.Endpoint
	ProbabilityEndpoint       endpoint.Endpoint
	MeasureEndpoint           endpoint.Endpoint
	ProbabilityZeroAtEndpoint endpoint.Endpoint
	ProbabilityOneAtEndpoint  endpoint.Endpoint
	MeasureAtEndpoint         endpoint.Endpoint
}
