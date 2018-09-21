// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// Add adds matrix B to A
func (A *Matrix) Add(B *Matrix) error {

	if !A.EqualShape(B) {
		return ErrMatSize
	}

	for i := 0; i < A.m; i++ {
		for j := 0; j < A.n; j++ {
			A.Inc(i, j, B.At(i, j))
		}
	}
	return nil
}

// Add A+B
func Add(A, B *Matrix) (*Matrix, error) {
	if !A.EqualShape(B) {
		return nil, ErrMatSize
	}

	C := Zero(A.m, A.n)
	for i := 0; i < A.m; i++ {
		for j := 0; j < A.n; j++ {
			C.Set(i, j, A.At(i, j)+B.At(i, j))
		}
	}
	return C, nil
}
