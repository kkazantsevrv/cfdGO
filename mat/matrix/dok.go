package matrix

import (
	"fmt"
)

// Entry represents a matrix entry with its position and value
type Entry struct {
	Row   int
	Col   int
	Value float64
}

// DOKMatrix represents a sparse matrix in Dictionary of Keys format
type DOKMatrix struct {
	entries map[int]map[int]float64 // Row -> Col -> Value mapping
	Rows    int
	Cols    int
}

// NewDOKMatrix creates a new DOK matrix with given dimensions
func NewDOKMatrix(rows, cols int) (*DOKMatrix, error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("invalid dimensions: rows=%d, cols=%d", rows, cols)
	}

	return &DOKMatrix{
		entries: make(map[int]map[int]float64),
		Rows:    rows,
		Cols:    cols,
	}, nil
}

// Set sets the value at position (i,j)
func (m *DOKMatrix) Set(i, j int, value float64) error {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
		return fmt.Errorf("index out of bounds")
	}

	if row, exists := m.entries[i]; exists {
		if value != 0 {
			row[j] = value
		} else {
			delete(row, j)
			if len(row) == 0 {
				delete(m.entries, i)
			}
		}
	} else if value != 0 {
		m.entries[i] = map[int]float64{j: value}
	}

	return nil
}

// Get returns the value at position (i,j)
func (m *DOKMatrix) Get(i, j int) float64 {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
		return 0.0
	}

	if row, exists := m.entries[i]; exists {
		return row[j]
	}
	return 0.0
}

// ToCSR converts the DOK matrix to CSR format
func (m *DOKMatrix) ToCSR() (*CSRMatrix, error) {
	// Count non-zero elements
	nnz := 0
	for _, row := range m.entries {
		nnz += len(row)
	}

	// Allocate arrays
	values := make([]float64, 0, nnz)
	colIndices := make([]int, 0, nnz)
	rowPtr := make([]int, m.Rows+1)

	// Fill arrays
	currentPos := 0
	for i := 0; i < m.Rows; i++ {
		rowPtr[i] = currentPos

		if row, exists := m.entries[i]; exists {
			// Get and sort column indices for this row
			cols := make([]int, 0, len(row))
			for j := range row {
				cols = append(cols, j)
			}
			sortInts(cols)

			// Add values in column order
			for _, j := range cols {
				values = append(values, row[j])
				colIndices = append(colIndices, j)
				currentPos++
			}
		}
	}
	rowPtr[m.Rows] = currentPos

	return NewCSRMatrix(values, rowPtr, colIndices, m.Rows, m.Cols)
}

// FromCSR converts a CSR matrix to DOK format
func FromCSRToDOK(csr *CSRMatrix) (*DOKMatrix, error) {
	dok, err := NewDOKMatrix(csr.Rows, csr.Cols)
	if err != nil {
		return nil, err
	}

	for i := 0; i < csr.Rows; i++ {
		for j := csr.RowPtr[i]; j < csr.RowPtr[i+1]; j++ {
			err := dok.Set(i, csr.ColIndices[j], csr.Values[j])
			if err != nil {
				return nil, err
			}
		}
	}

	return dok, nil
}

// NonZeros returns the number of non-zero elements in the matrix
func (m *DOKMatrix) NonZeros() int {
	count := 0
	for _, row := range m.entries {
		count += len(row)
	}
	return count
}

// sortInts sorts a slice of integers in ascending order
func sortInts(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
