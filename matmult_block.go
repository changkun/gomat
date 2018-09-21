// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// DotBlock matrix multiplication
// Use JIK block 36 version here, see ./benchmark/README.md
func (A *Matrix) DotBlock(B, C *Matrix) (err error) {
	return A.DotBlockJIK(36, B, C)
}

// DotBlockP matrix multiplication
// Use JIKP block 36 version here, see ./benchmark/README.md
func (A *Matrix) DotBlockP(B, C *Matrix) (err error) {
	return A.DotBlockJIKP(36, B, C)
}
