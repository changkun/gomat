// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// DotNaive matrix multiplication O(n^3)
// Use JIK version here, see ./benchmark/README.md
func (A *Matrix) DotNaive(B, C *Matrix) (err error) {
	return A.DotNaiveJIK(B, C)
}

// DotNaiveP matrix multiplication O(n^3)
// Use JIKP version here, see ./benchmark/README.md
func (A *Matrix) DotNaiveP(B, C *Matrix) (err error) {
	return A.DotNaiveJIKP(B, C)
}
