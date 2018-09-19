// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"sync"
)

// DotNaive matrix multiplication O(n^3)
func (A *Matrix) DotNaive(B, C *Matrix) (err error) {
	var (
		i, j, k int
		sum     float64
	)
	if A.Col() != B.Row() || C.Row() != A.Row() || C.Col() != B.Col() {
		return ErrMatSize
	}

	for i = 0; i < A.Row(); i++ {
		for j = 0; j < B.Col(); j++ {
			sum = 0.0
			for k = 0; k < A.Col(); k++ {
				sum += A.At(i, k) * B.At(k, j)
			}
			C.Set(i, j, sum)
		}
	}
	return
}

// DotNaiveP matrix multiplication O(n^3)
func (A *Matrix) DotNaiveP(B, C *Matrix) (err error) {
	if A.Col() != B.Row() || C.Row() != A.Row() || C.Col() != B.Col() {
		return ErrMatSize
	}

	wg := sync.WaitGroup{}
	for i := 0; i < A.Row(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < B.Col(); j++ {
				sum := 0.0
				for k := 0; k < A.Col(); k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Set(i, j, sum)
			}
		}(i)
	}
	wg.Wait()
	return
}
