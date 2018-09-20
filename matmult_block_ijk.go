// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// import "sync"

// DotBlockIJK block matrix multiplication
func (A *Matrix) DotBlockIJK(blockSize int, B, C *Matrix) (err error) {
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
		sum             float64
		en              = blockSize * (min / blockSize)
	)

	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i = 0; i < A.Row(); i++ {
				for j = jj; j < jj+blockSize; j++ {
					sum = 0.0
					for k = kk; k < kk+blockSize; k++ {
						sum += A.At(i, k) * B.At(k, j)
					}
					C.Inc(i, j, sum)
				}
			}
		}

		// residue right
		for i = 0; i < A.Row(); i++ {
			for j = en; j < B.Col(); j++ {
				sum = 0.0
				for k = kk; k < kk+blockSize; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Inc(i, j, sum)
			}
		}
	}

	// residue bottom
	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < A.Row(); i++ {
			for j = jj; j < jj+blockSize; j++ {
				sum = 0.0
				for k = en; k < A.Col(); k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Inc(i, j, sum)
			}
		}
	}

	// residule bottom right
	for i = 0; i < A.Row(); i++ {
		for j = en; j < B.Col(); j++ {
			sum = 0.0
			for k = en; k < A.Col(); k++ {
				sum += A.At(i, k) * B.At(k, j)
			}
			C.Inc(i, j, sum)
		}
	}
	return
}

// ParalMultIJK matrix multiplication
// TODO: optimize cache aware
// func (A *Matrix) ParalMultIJK(blockSize int, B, C *Matrix) (err error) {
// 	var (
// 		kk, jj, i, j int
// 		N            = A.N
// 		en           = blockSize * (N / blockSize)
// 	)
// 	if N != B.N || N != C.N {
// 		return ErrMatrixSize
// 	}

// 	wg := sync.WaitGroup{}
// 	for kk = 0; kk < en; kk += blockSize {
// 		for jj = 0; jj < en; jj += blockSize {
// 			for i = 0; i < N; i++ {
// 				for j = jj; j < jj+blockSize; j++ {
// 					wg.Add(1)
// 					go func(i, j int) {
// 						defer wg.Done()
// 						sum := 0.0
// 						for k := kk; k < kk+blockSize; k++ {
// 							sum += A.At(i, k) * B.At(k, j)
// 						}
// 						C.Increment(i, j, sum)
// 					}(i, j)
// 				}
// 			}
// 			wg.Wait()
// 		}

// 		for i = 0; i < N; i++ {
// 			for j = en; j < N; j++ {
// 				wg.Add(1)
// 				go func(i, j int) {
// 					defer wg.Done()
// 					sum := 0.0
// 					for k := kk; k < kk+blockSize; k++ {
// 						sum += A.At(i, k) * B.At(k, j)
// 					}
// 					C.Increment(i, j, sum)
// 				}(i, j)
// 			}
// 		}
// 		wg.Wait()
// 	}
// 	for jj = 0; jj < en; jj += blockSize {
// 		for i = 0; i < N; i++ {
// 			for j = jj; j < jj+blockSize; j++ {
// 				wg.Add(1)
// 				go func(i, j int) {
// 					defer wg.Done()
// 					sum := 0.0
// 					for k := en; k < N; k++ {
// 						sum += A.At(i, k) * B.At(k, j)
// 					}
// 					C.Increment(i, j, sum)
// 				}(i, j)
// 			}
// 		}
// 		wg.Wait()
// 	}
// 	for i = 0; i < N; i++ {
// 		for j = en; j < N; j++ {
// 			wg.Add(1)
// 			go func(i, j int) {
// 				defer wg.Done()
// 				sum := 0.0
// 				for k := en; k < N; k++ {
// 					sum += A.At(i, k) * B.At(k, j)
// 				}
// 				C.Increment(i, j, sum)
// 			}(i, j)
// 		}
// 	}
// 	wg.Wait()
// 	return
// }
