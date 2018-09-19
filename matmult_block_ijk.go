// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// import "sync"

// // MultIJK matrix multiplication
// func (A *Matrix) MultIJK(blockSize int, B, C *Matrix) (err error) {
// 	var (
// 		kk, jj, i, j, k int
// 		sum             float64
// 		N               = A.N
// 		en              = blockSize * (N / blockSize)
// 	)
// 	if N != B.N || N != C.N {
// 		return ErrMatrixSize
// 	}

// 	for kk = 0; kk < en; kk += blockSize {
// 		for jj = 0; jj < en; jj += blockSize {
// 			for i = 0; i < N; i++ {
// 				for j = jj; j < jj+blockSize; j++ {
// 					sum = 0.0
// 					for k = kk; k < kk+blockSize; k++ {
// 						sum += A.At(i, k) * B.At(k, j)
// 					}
// 					C.Increment(i, j, sum)
// 				}
// 			}
// 		}

// 		for i = 0; i < N; i++ {
// 			for j = en; j < N; j++ {
// 				sum = 0.0
// 				for k = kk; k < kk+blockSize; k++ {
// 					sum += A.At(i, k) * B.At(k, j)
// 				}
// 				C.Increment(i, j, sum)
// 			}
// 		}
// 	}

// 	for jj = 0; jj < en; jj += blockSize {
// 		for i = 0; i < N; i++ {
// 			for j = jj; j < jj+blockSize; j++ {
// 				sum = 0.0
// 				for k = en; k < N; k++ {
// 					sum += A.At(i, k) * B.At(k, j)
// 				}
// 				C.Increment(i, j, sum)
// 			}
// 		}
// 	}
// 	for i = 0; i < N; i++ {
// 		for j = en; j < N; j++ {
// 			sum = 0.0
// 			for k = en; k < N; k++ {
// 				sum += A.At(i, k) * B.At(k, j)
// 			}
// 			C.Increment(i, j, sum)
// 		}
// 	}
// 	return
// }

// // ParalMultIJK matrix multiplication
// // TODO: optimize cache aware
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
