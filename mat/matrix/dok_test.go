package matrix

import (
	"math"
	"testing"
)

func TestNewDOKMatrix(t *testing.T) {
	tests := []struct {
		name    string
		rows    int
		cols    int
		wantErr bool
	}{
		{
			name:    "valid dimensions",
			rows:    3,
			cols:    4,
			wantErr: false,
		},
		{
			name:    "zero rows",
			rows:    0,
			cols:    4,
			wantErr: true,
		},
		{
			name:    "negative columns",
			rows:    3,
			cols:    -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDOKMatrix(tt.rows, tt.cols)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDOKMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDOKSetGet(t *testing.T) {
	m, _ := NewDOKMatrix(3, 3)

	// Test setting and getting values
	testCases := []struct {
		i, j     int
		value    float64
		wantErr  bool
		expected float64
	}{
		{0, 0, 1.0, false, 1.0},
		{1, 1, 2.0, false, 2.0},
		{2, 2, 3.0, false, 3.0},
		{1, 2, 4.0, false, 4.0},
		{3, 3, 5.0, true, 0.0},  // Out of bounds
		{-1, 0, 6.0, true, 0.0}, // Out of bounds
	}

	for _, tc := range testCases {
		err := m.Set(tc.i, tc.j, tc.value)
		if (err != nil) != tc.wantErr {
			t.Errorf("Set(%d,%d,%f) error = %v, wantErr %v", tc.i, tc.j, tc.value, err, tc.wantErr)
		}

		got := m.Get(tc.i, tc.j)
		if math.Abs(got-tc.expected) > 1e-15 {
			t.Errorf("Get(%d,%d) = %f, want %f", tc.i, tc.j, got, tc.expected)
		}
	}
}

func TestDOKToCSRConversion(t *testing.T) {
	// Create a DOK matrix with known values
	dok, _ := NewDOKMatrix(3, 3)
	dok.Set(0, 0, 1.0)
	dok.Set(0, 2, 2.0)
	dok.Set(1, 1, 3.0)
	dok.Set(2, 0, 4.0)

	// Convert to CSR
	csr, err := dok.ToCSR()
	if err != nil {
		t.Fatalf("Failed to convert to CSR: %v", err)
	}

	// Verify dimensions
	if csr.Rows != 3 || csr.Cols != 3 {
		t.Errorf("Wrong dimensions after conversion")
	}

	// Verify values through Get method
	testCases := []struct {
		i, j     int
		expected float64
	}{
		{0, 0, 1.0},
		{0, 2, 2.0},
		{1, 1, 3.0},
		{2, 0, 4.0},
		{0, 1, 0.0}, // Zero element
		{1, 0, 0.0}, // Zero element
	}

	for _, tc := range testCases {
		got := csr.Get(tc.i, tc.j)
		if math.Abs(got-tc.expected) > 1e-15 {
			t.Errorf("Get(%d,%d) = %f, want %f", tc.i, tc.j, got, tc.expected)
		}
	}
}

func TestCSRToDOKConversion(t *testing.T) {
	// Create a CSR matrix
	csr, _ := NewCSRMatrix(
		[]float64{1.0, 2.0, 3.0, 4.0},
		[]int{0, 2, 3, 4},
		[]int{0, 2, 1, 0},
		3, 3,
	)

	// Convert to DOK
	dok, err := FromCSRToDOK(csr)
	if err != nil {
		t.Fatalf("Failed to convert to DOK: %v", err)
	}

	// Verify dimensions
	if dok.Rows != 3 || dok.Cols != 3 {
		t.Errorf("Wrong dimensions after conversion")
	}

	// Verify values
	testCases := []struct {
		i, j     int
		expected float64
	}{
		{0, 0, 1.0},
		{0, 2, 2.0},
		{1, 1, 3.0},
		{2, 0, 4.0},
		{0, 1, 0.0}, // Zero element
		{1, 0, 0.0}, // Zero element
	}

	for _, tc := range testCases {
		got := dok.Get(tc.i, tc.j)
		if math.Abs(got-tc.expected) > 1e-15 {
			t.Errorf("Get(%d,%d) = %f, want %f", tc.i, tc.j, got, tc.expected)
		}
	}
}

func TestDOKNonZeros(t *testing.T) {
	m, _ := NewDOKMatrix(3, 3)

	// Empty matrix should have zero non-zeros
	if m.NonZeros() != 0 {
		t.Errorf("Empty matrix has %d non-zeros, want 0", m.NonZeros())
	}

	// Add some non-zero elements
	m.Set(0, 0, 1.0)
	m.Set(1, 1, 2.0)
	m.Set(2, 2, 3.0)

	if m.NonZeros() != 3 {
		t.Errorf("Matrix has %d non-zeros, want 3", m.NonZeros())
	}

	// Set an element to zero (should decrease non-zero count)
	m.Set(1, 1, 0.0)

	if m.NonZeros() != 2 {
		t.Errorf("Matrix has %d non-zeros after zero set, want 2", m.NonZeros())
	}
}
