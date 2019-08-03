package service

import (
	"github.com/sah4ez/qubit/pkg/math/matrix"
)

// func New(z ...complex128) *Qubit {
// v := v.Vector{}
// for _, zi := range z {
// v = append(v, zi)
// }
// q := &Qubit{v}
// q.Normalize()
// return q
// }
//
// func Zero(bit ...int) *Qubit {
// return &Qubit{v.TensorProductN(v.Vector{1, 0}, bit...)}
// }
//
// func One(bit ...int) *Qubit {
// return &Qubit{v.TensorProductN(v.Vector{0, 1}, bit...)}
// }

type IQubit interface {
	NumberOfBit() int
	IsZero(eps ...float64) bool
	IsOne(eps ...float64) bool
	InnerProduct(q0 IQubit) complex128
	OuterProduct(q0 IQubit) matrix.Matrix
	Clone() IQubit
	Fidelity(q0 IQubit) float64
	TraceDistance(q0 IQubit) float64
	Equals(q0 IQubit, eps ...float64) bool
	TensorProduct(q0 IQubit) IQubit
	Apply(m matrix.Matrix) IQubit
	Normalize() IQubit
	Amplitude() []complex128
	Probability() []float64
	Measure(bit ...int) IQubit
	ProbabilityZeroAt(bit int) ([]int, []float64)
	ProbabilityOneAt(bit int) ([]int, []float64)
	MeasureAt(bit int) IQubit
}
