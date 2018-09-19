// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"fmt"
	"testing"
)

// type NaiveMult func(B, C *Matrix) error

func TestMatrix_MultNaive(t *testing.T) {
	A, err := New(2, 3)(
		1, 2, 3,
		2, 3, 1,
	)
	if err != nil {
		t.Errorf("New(2, 3) error, expect nil")
	}
	B, err := New(3, 1)(
		3,
		2,
		1,
	)
	if err != nil {
		t.Errorf("New(3, 1) error, expect nil")
	}

	t.Run("DotNaive()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaive(B, C)
		if err != nil {
			t.Errorf("DotNaive(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaive() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveP()", func(t *testing.T) {
		C, err := New(2, 1)(
			0,
			0,
		)
		T, err := New(2, 1)(
			10,
			13,
		)
		if err != nil {
			t.Errorf("New(2, 1) error, expect nil")
		}
		err = A.DotNaiveP(B, C)
		if err != nil {
			t.Errorf("DotNaiveP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaive() not euqal, expect euqal, got:")
			C.Print()
		}
	})
}

// // ----------------------- benchmarks ----------------------------

func BenchmarkNaive(b *testing.B) {
	for n := 0; n < 100; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		b.Run(fmt.Sprintf("DotNaive() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaive(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveP() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveP(B, C)
			}
		})
	}
}

// var (
// 	bSize = []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	A     = &Matrix{
// 		N: 8,
// 		data: [][]float64{
// 			[]float64{1, 2, 3, 4, 5, 6, 7, 8},
// 			[]float64{9, 1, 2, 3, 4, 5, 6, 7},
// 			[]float64{8, 9, 1, 2, 3, 4, 5, 6},
// 			[]float64{7, 8, 9, 1, 2, 3, 4, 5},
// 			[]float64{6, 7, 8, 9, 1, 2, 3, 4},
// 			[]float64{5, 6, 7, 8, 9, 1, 2, 3},
// 			[]float64{4, 5, 6, 7, 8, 9, 1, 2},
// 			[]float64{3, 4, 5, 6, 7, 8, 9, 0},
// 		},
// 	}
// 	B = &Matrix{
// 		N: 8,
// 		data: [][]float64{
// 			[]float64{9, 8, 7, 6, 5, 4, 3, 2},
// 			[]float64{1, 9, 8, 7, 6, 5, 4, 3},
// 			[]float64{2, 1, 9, 8, 7, 6, 5, 4},
// 			[]float64{3, 2, 1, 9, 8, 7, 6, 5},
// 			[]float64{4, 3, 2, 1, 9, 8, 7, 6},
// 			[]float64{5, 4, 3, 2, 1, 9, 8, 7},
// 			[]float64{6, 5, 4, 3, 2, 1, 9, 8},
// 			[]float64{7, 6, 5, 4, 3, 2, 1, 0},
// 		},
// 	}
// 	C = &Matrix{
// 		N: 8,
// 		data: [][]float64{
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 		},
// 	}
// )

// func BenchmarkMatrix_MultNaive(b *testing.B) {
// 	tests := []struct {
// 		name string
// 		f    NaiveMult
// 	}{
// 		{
// 			name: "A.MultNaive",
// 			f:    A.MultNaive,
// 		},
// 		{
// 			name: "A.ParalMultNaive",
// 			f:    A.ParalMultNaive,
// 		},
// 	}
// 	for _, tt := range tests {
// 		b.Run(tt.name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				if err := tt.f(B, C); err != nil {
// 					b.Errorf("%s() error %v, want nil", tt.name, err)
// 				}
// 			}
// 		})
// 	}
// }
