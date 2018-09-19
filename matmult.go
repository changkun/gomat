package gomat

import (
	"errors"
	"fmt"
	"sync"
)

// Errors
var (
	ErrNumElements = errors.New("Error number of matrix elements")
	ErrMatrixSize  = errors.New("Error size of matrix")
)

// Matrix is a 2d array
type Matrix struct {
	N    int
	data [][]float64
}

// New a size by size matrix
func New(size int) func(...float64) (*Matrix, error) {
	wg := sync.WaitGroup{}
	d := make([][]float64, size)
	for i := range d {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			d[i] = make([]float64, size)
		}(i)
	}
	wg.Wait()
	m := &Matrix{
		N:    size,
		data: d,
	}
	return func(es ...float64) (*Matrix, error) {
		if len(es) != size*size {
			return nil, ErrNumElements
		}
		for i := range es {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				m.data[i/size][i%size] = es[i]
			}(i)
		}
		wg.Wait()
		return m, nil
	}
}

func (A *Matrix) debug() {
	for i := 0; i < A.N; i++ {
		for j := 0; j < A.N; j++ {
			fmt.Println(&A.data[i][j], " ")
		}
	}
}

// At access (i, j) element
func (A *Matrix) At(i, j int) float64 {
	return A.data[i][j]
}

// Set set (i, j) element with val
func (A *Matrix) Set(i, j int, val float64) {
	A.data[i][j] = val
}

// Add adds (i, j) element with wal
func (A *Matrix) Add(i, j int, val float64) {
	A.data[i][j] += val
}

// MultNaive matrix multiplication O(n^3)
func (A *Matrix) MultNaive(B, C *Matrix) (err error) {
	var (
		i, j, k int
		sum     float64
		N       = A.N
	)

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			sum = 0.0
			for k = 0; k < N; k++ {
				sum += A.At(i, k) * B.At(k, j)
			}
			C.Set(i, j, sum)
		}
	}
	return
}

// ParalMultNaive matrix multiplication O(n^3)
// TODO: optimize cache aware
func (A *Matrix) ParalMultNaive(B, C *Matrix) (err error) {
	var (
		i, j int
		N    = A.N
	)

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	wg := sync.WaitGroup{}
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0.0
				for k := 0; k < N; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Set(i, j, sum)
			}(i, j)
		}
	}
	wg.Wait()
	return
}

// MultIJK matrix multiplication
func (A *Matrix) MultIJK(blockSize int, B, C *Matrix) (err error) {
	var (
		kk, jj, i, j, k int
		sum             float64
		N               = A.N
		en              = blockSize * (N / blockSize)
	)
	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i = 0; i < N; i++ {
				for j = jj; j < jj+blockSize; j++ {
					sum = 0.0
					for k = kk; k < kk+blockSize; k++ {
						sum += A.At(i, k) * B.At(k, j)
					}
					C.Add(i, j, sum)
				}
			}
		}

		for i = 0; i < N; i++ {
			for j = en; j < N; j++ {
				sum = 0.0
				for k = kk; k < kk+blockSize; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Add(i, j, sum)
			}
		}
	}

	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < N; i++ {
			for j = jj; j < jj+blockSize; j++ {
				sum = 0.0
				for k = en; k < N; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Add(i, j, sum)
			}
		}
	}
	for i = 0; i < N; i++ {
		for j = en; j < N; j++ {
			sum = 0.0
			for k = en; k < N; k++ {
				sum += A.At(i, k) * B.At(k, j)
			}
			C.Add(i, j, sum)
		}
	}
	return
}

// ParalMultIJK matrix multiplication
// TODO: optimize cache aware
func (A *Matrix) ParalMultIJK(blockSize int, B, C *Matrix) (err error) {
	var (
		kk, jj, i, j int
		N            = A.N
		en           = blockSize * (N / blockSize)
	)
	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	wg := sync.WaitGroup{}
	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i = 0; i < N; i++ {
				for j = jj; j < jj+blockSize; j++ {
					wg.Add(1)
					go func(i, j int) {
						defer wg.Done()
						sum := 0.0
						for k := kk; k < kk+blockSize; k++ {
							sum += A.At(i, k) * B.At(k, j)
						}
						C.Add(i, j, sum)
					}(i, j)
				}
			}
			wg.Wait()
		}

		for i = 0; i < N; i++ {
			for j = en; j < N; j++ {
				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					sum := 0.0
					for k := kk; k < kk+blockSize; k++ {
						sum += A.At(i, k) * B.At(k, j)
					}
					C.Add(i, j, sum)
				}(i, j)
			}
		}
		wg.Wait()
	}
	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < N; i++ {
			for j = jj; j < jj+blockSize; j++ {
				wg.Add(1)
				go func(i, j int) {
					defer wg.Done()
					sum := 0.0
					for k := en; k < N; k++ {
						sum += A.At(i, k) * B.At(k, j)
					}
					C.Add(i, j, sum)
				}(i, j)
			}
		}
		wg.Wait()
	}
	for i = 0; i < N; i++ {
		for j = en; j < N; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0.0
				for k := en; k < N; k++ {
					sum += A.At(i, k) * B.At(k, j)
				}
				C.Add(i, j, sum)
			}(i, j)
		}
	}
	wg.Wait()
	return
}

// MultIKJ matrix multiplication
func (A *Matrix) MultIKJ(blockSize int, B, C *Matrix) (err error) {
	var (
		kk, jj, i, j, k int
		r               float64
		N               = A.N
		en              = blockSize * (N / blockSize)
	)

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i = 0; i < N; i++ {
				for k = kk; k < kk+blockSize; k++ {
					r = A.At(i, k)
					for j = jj; j < jj+blockSize; j++ {
						C.Add(i, j, r*B.At(k, j))
					}
				}
			}
		}
		for i = 0; i < N; i++ {
			for k = kk; k < kk+blockSize; k++ {
				r = A.At(i, k)
				for j = en; j < N; j++ {
					C.Add(i, j, r*B.At(k, j))
				}
			}
		}
	}
	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < N; i++ {
			for k = en; k < N; k++ {
				r = A.At(i, k)
				for j = jj; j < jj+blockSize; j++ {
					C.Add(i, j, r*B.At(k, j))
				}
			}
		}
	}
	for i = 0; i < N; i++ {
		for k = en; k < N; k++ {
			r = A.At(i, k)
			for j = en; j < N; j++ {
				C.Add(i, j, r*B.At(k, j))
			}
		}
	}

	return
}

// ParalMultIKJ matrix multiplication
// TODO: optimize cache aware
func (A *Matrix) ParalMultIKJ(blockSize int, B, C *Matrix) (err error) {
	var (
		kk, jj, i, k int
		N            = A.N
		en           = blockSize * (N / blockSize)
	)

	if N != B.N || N != C.N {
		return ErrMatrixSize
	}

	for kk = 0; kk < en; kk += blockSize {
		for jj = 0; jj < en; jj += blockSize {
			for i := 0; i < N; i++ {
				for k := kk; k < kk+blockSize; k++ {
					r := A.At(i, k)
					for j := jj; j < jj+blockSize; j++ {
						C.Add(i, j, r*B.At(k, j))
					}
				}
			}
		}

		for i = 0; i < N; i++ {
			for k = kk; k < kk+blockSize; k++ {
				r := A.At(i, k)
				for j := en; j < N; j++ {
					C.Add(i, j, r*B.At(k, j))
				}
			}
		}
	}
	for jj = 0; jj < en; jj += blockSize {
		for i = 0; i < N; i++ {
			for k = en; k < N; k++ {
				r := A.At(i, k)
				for j := jj; j < jj+blockSize; j++ {
					C.Add(i, j, r*B.At(k, j))
				}
			}
		}
	}
	for i = 0; i < N; i++ {
		for k = en; k < N; k++ {
			r := A.At(i, k)
			for j := en; j < N; j++ {
				C.Add(i, j, r*B.At(k, j))
			}
		}
	}

	return
}
