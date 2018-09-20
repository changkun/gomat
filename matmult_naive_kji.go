// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import "sync"

// DotNaiveKJI matrix multiplication O(n^3)
func (A *Matrix) DotNaiveKJI(B, C *Matrix) (err error) {
	var (
		i, j, k int
		r       float64
	)
	if A.Col() != B.Row() || C.Row() != A.Row() || C.Col() != B.Col() {
		return ErrMatSize
	}

	for k = 0; k < A.Col(); k++ {
		for j = 0; j < B.Col(); j++ {
			r = B.At(k, j)
			for i = 0; i < A.Row(); i++ {
				C.Inc(i, j, r*A.At(i, k))
			}
		}
	}
	return
}

// DotNaiveKJIP matrix multiplication O(n^3)
func (A *Matrix) DotNaiveKJIP(B, C *Matrix) (err error) {
	if A.Col() != B.Row() || C.Row() != A.Row() || C.Col() != B.Col() {
		return ErrMatSize
	}

	wg := sync.WaitGroup{}
	for k := 0; k < A.Col(); k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for j := 0; j < B.Col(); j++ {
				r := B.At(k, j)
				for i := 0; i < A.Row(); i++ {
					C.Inc(i, j, r*A.At(i, k))
				}
			}
		}(k)
	}
	wg.Wait()
	return
}
