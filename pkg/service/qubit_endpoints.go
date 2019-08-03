package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sah4ez/qubit/pkg/math/matrix"
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

func (qs *QubitSet) NumberOfBit() int {
	panic("not implemented")
	response, _ := qs.NumberOfBitEndpoint(context.TODO(), nil)
	return response.(int)

}

func (qs *QubitSet) IsZero(eps ...float64) bool {
	panic("not implemented")
}

func (qs *QubitSet) IsOne(eps ...float64) bool {
	panic("not implemented")
}

func (qs *QubitSet) InnerProduct(q0 IQubit) complex128 {
	panic("not implemented")
}

func (qs *QubitSet) OuterProduct(q0 IQubit) matrix.Matrix {
	panic("not implemented")
}

func (qs *QubitSet) Clone() IQubit {
	panic("not implemented")
}

func (qs *QubitSet) Fidelity(q0 IQubit) float64 {
	panic("not implemented")
}

func (qs *QubitSet) TraceDistance(q0 IQubit) float64 {
	panic("not implemented")
}

func (qs *QubitSet) Equals(q0 IQubit, eps ...float64) bool {
	panic("not implemented")
}

func (qs *QubitSet) TensorProduct(q0 IQubit) IQubit {
	panic("not implemented")
}

func (qs *QubitSet) Apply(m matrix.Matrix) IQubit {
	panic("not implemented")
}

func (qs *QubitSet) Normalize() IQubit {
	panic("not implemented")
}

func (qs *QubitSet) Amplitude() []complex128 {
	panic("not implemented")
}

func (qs *QubitSet) Probability() []float64 {
	panic("not implemented")
}

func (qs *QubitSet) Measure(bit ...int) IQubit {
	panic("not implemented")
}

func (qs *QubitSet) ProbabilityZeroAt(bit int) ([]int, []float64) {
	panic("not implemented")
}

func (qs *QubitSet) ProbabilityOneAt(bit int) ([]int, []float64) {
	panic("not implemented")
}

func (qs *QubitSet) MeasureAt(bit int) IQubit {
	panic("not implemented")
}

func makeNumberOfBit(s IQubit) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return s.NumberOfBit(), nil
	}
}

type EpsRequest struct {
	Eps []float64
}

type QubitRequest struct {
	Qubit IQubit
}

type QubitResponse struct {
	Qubit IQubit
}

type EqualsRequest struct {
	Qubit IQubit
	Eps   []float64
}

type BitsRequest struct {
	Bit []int
}

type ProbabilityResponse struct {
	Index       []int
	Probability []float64
}

func makeIsZero(s IQubit) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EpsRequest)
		return s.IsZero(req.Eps...), nil
	}
}

func makeIsOne(s IQubit) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EpsRequest)
		return s.IsOne(req.Eps...), nil
	}
}

func makeInnerProduct(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeOuterProduct(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeClone(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeFidelity(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeTraceDistance(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeEquals(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeTensorProduct(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeApply(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeNormalize(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeAmplitude(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeProbability(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeMeasure(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeProbabilityZeroAt(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeProbabilityOneAt(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func makeMeasureAt(s IQubit) endpoint.Endpoint {
	panic("not implemented")
}

func MakeEndpoint(s IQubit) (set QubitSet) {
	set.NumberOfBitEndpoint = makeNumberOfBit(s)
	set.IsZeroEndpoint = makeIsZero(s)
	set.IsOneEndpoint = makeIsOne(s)
	set.InnerProductEndpoint = makeInnerProduct(s)
	set.OuterProductEndpoint = makeOuterProduct(s)
	set.FidelityEndpoint = makeFidelity(s)
	set.TraceDistanceEndpoint = makeTraceDistance(s)
	set.EqualsEndpoint = makeEquals(s)
	set.TensorProductEndpoint = makeTensorProduct(s)
	set.ApplyEndpoint = makeApply(s)
	set.NormalizeEndpoint = makeNormalize(s)
	set.AmplitudeEndpoint = makeAmplitude(s)
	set.ProbabilityEndpoint = makeProbability(s)
	set.MeasureEndpoint = makeMeasure(s)
	set.ProbabilityZeroAtEndpoint = makeProbabilityZeroAt(s)
	set.ProbabilityOneAtEndpoint = makeProbabilityOneAt(s)
	set.MeasureAtEndpoint = makeMeasureAt(s)
	return
}
