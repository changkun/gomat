// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// DotBlockIKJ block matrix multiplication
func (A *Matrix) DotBlockIKJ(blockSize int, B, C *Matrix) (err error) {
	if (A.Col() != B.Row()) || (C.Row() != A.Row()) || (C.Col() != B.Col()) {
		return ErrMatSize
	}
	min := A.Row()
	if A.Col() < min {
		min = A.Col()
	}
	if B.Col() < min {
		min = B.Col()
	}
	var (
		kk, jj, i, j, k int
		r               float64
		en              = blockSize * (min / blockSize)
	)

	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i = 0; i < A.Row(); i++ {
				for k = kk; k < kk+blockSize; k++ {
					r = A.At(i, k)
					for j = jj; j < jj+blockSize; j++ {
						C.Inc(i, j, r*B.At(k, j))
					}
				}
			}
		}
		for i = 0; i < A.Row(); i++ {
			for k = kk; k < kk+blockSize; k++ {
				r = A.At(i, k)
				for j = en; j < B.Col(); j++ {
					C.Inc(i, j, r*B.At(k, j))
				}
			}
		}
	}

	// residule bottom
	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < A.Row(); i++ {
			for k = en; k < A.Col(); k++ {
				r = A.At(i, k)
				for j = jj; j < jj+blockSize; j++ {
					C.Inc(i, j, r*B.At(k, j))
				}
			}
		}
	}

	// residule bottom right
	for i = 0; i < A.Row(); i++ {
		for k = en; k < A.Col(); k++ {
			r = A.At(i, k)
			for j = en; j < B.Col(); j++ {
				C.Inc(i, j, r*B.At(k, j))
			}
		}
	}
	return
}

// // ParalMultIKJ matrix multiplication
// // TODO: optimize cache aware
// func (A *Matrix) ParalMultIKJ(blockSize int, B, C *Matrix) (err error) {
// 	var (
// 		kk, jj, i, k int
// 		N            = A.N
// 		en           = blockSize * (N / blockSize)
// 	)

// 	if N != B.N || N != C.N {
// 		return ErrMatrixSize
// 	}

// 	for kk = 0; kk < en; kk += blockSize {
// 		for jj = 0; jj < en; jj += blockSize {
// 			for i := 0; i < N; i++ {
// 				for k := kk; k < kk+blockSize; k++ {
// 					r := A.At(i, k)
// 					for j := jj; j < jj+blockSize; j++ {
// 						C.Increment(i, j, r*B.At(k, j))
// 					}
// 				}
// 			}
// 		}

// 		for i = 0; i < N; i++ {
// 			for k = kk; k < kk+blockSize; k++ {
// 				r := A.At(i, k)
// 				for j := en; j < N; j++ {
// 					C.Increment(i, j, r*B.At(k, j))
// 				}
// 			}
// 		}
// 	}
// 	for jj = 0; jj < en; jj += blockSize {
// 		for i = 0; i < N; i++ {
// 			for k = en; k < N; k++ {
// 				r := A.At(i, k)
// 				for j := jj; j < jj+blockSize; j++ {
// 					C.Increment(i, j, r*B.At(k, j))
// 				}
// 			}
// 		}
// 	}
// 	for i = 0; i < N; i++ {
// 		for k = en; k < N; k++ {
// 			r := A.At(i, k)
// 			for j := en; j < N; j++ {
// 				C.Increment(i, j, r*B.At(k, j))
// 			}
// 		}
// 	}

// 	return
// }
