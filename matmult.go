// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// Dot matrix multiplication
func (A *Matrix) Dot(B, C *Matrix) (err error) {
	return A.DotBlock(B, C)
}

// DotP matrix multiplication
func (A *Matrix) DotP(B, C *Matrix) (err error) {
	return A.DotBlockP(B, C)
}

// Dot matrix multiplication
func Dot(A, B *Matrix) (*Matrix, error) {
	C := Zero(A.Row(), B.Col())
	if err := A.Dot(B, C); err != nil {
		return nil, err
	}
	return C, nil
}

// DotP matrix multiplication
func DotP(A, B *Matrix) (*Matrix, error) {
	C := Zero(A.Row(), B.Col())
	if err := A.DotP(B, C); err != nil {
		return nil, err
	}
	return C, nil
}
