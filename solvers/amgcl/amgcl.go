package amgcl

/*
#cgo CXXFLAGS: -I./include -std=c++11
#cgo LDFLAGS: -L${SRCDIR} -lamgcl_wrapper -lstdc++
#include "amgcl_wrapper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"test.com/mat/matrix"
)

// Solver wraps the AMGCL solver
type Solver struct {
	solver C.AMGCLSolver // Changed from *C.AMGCLSolver to C.AMGCLSolver
}

func NewSolver(A *matrix.CSRMatrix) (*Solver, error) {
	if A == nil {
		return nil, fmt.Errorf("matrix cannot be nil")
	}

	rowPtr := make([]C.int, len(A.RowPtr))         // Changed size_t to int
	colIndices := make([]C.int, len(A.ColIndices)) // Changed size_t to int

	for i := range A.RowPtr {
		rowPtr[i] = C.int(A.RowPtr[i])
	}
	for i := range A.ColIndices {
		colIndices[i] = C.int(A.ColIndices[i])
	}

	s := &Solver{
		solver: C.create_solver(
			C.int(A.Rows),
			(*C.int)(unsafe.Pointer(&rowPtr[0])),
			(*C.int)(unsafe.Pointer(&colIndices[0])),
			(*C.double)(unsafe.Pointer(&A.Values[0])),
		),
	}
	return s, nil
}

func (s *Solver) Solve(rhs []float64) ([]float64, error) {
	if s.solver == nil {
		return nil, fmt.Errorf("solver has been freed")
	}
	x := make([]float64, len(rhs))
	C.solve_system(
		s.solver, // Pass the C.AMGCLSolver directly
		(*C.double)(unsafe.Pointer(&rhs[0])),
		(*C.double)(unsafe.Pointer(&x[0])),
	)
	return x, nil
}

// Optional: Add a Free method to clean up
func (s *Solver) Free() {
	if s.solver != nil {
		C.destroy_solver(s.solver)
		s.solver = nil
	}
}
