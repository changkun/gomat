// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// DotBlockKIJ block matrix multiplication
func (A *Matrix) DotBlockKIJ(blockSize int, B, C *Matrix) (err error) {
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
			for k = kk; k < kk+blockSize; k++ {
				for i = 0; i < A.Row(); i++ {
					r = A.At(i, k)
					for j = jj; j < jj+blockSize; j++ {
						C.Inc(i, j, r*B.At(k, j))
					}
				}
			}
		}
		for k = kk; k < kk+blockSize; k++ {
			for i = 0; i < A.Row(); i++ {
				r = A.At(i, k)
				for j = en; j < B.Col(); j++ {
					C.Inc(i, j, r*B.At(k, j))
				}
			}
		}
	}

	// residule bottom
	for jj = 0; jj < en; jj += blockSize {
		for k = en; k < A.Col(); k++ {
			for i = 0; i < A.Row(); i++ {
				r = A.At(i, k)
				for j = jj; j < jj+blockSize; j++ {
					C.Inc(i, j, r*B.At(k, j))
				}
			}
		}
	}

	// residule bottom right
	for k = en; k < A.Col(); k++ {
		for i = 0; i < A.Row(); i++ {
			r = A.At(i, k)
			for j = en; j < B.Col(); j++ {
				C.Inc(i, j, r*B.At(k, j))
			}
		}
	}
	return
}
