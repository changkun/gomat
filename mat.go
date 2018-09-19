// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// Errors
var (
	ErrNumElem = errors.New("bad number of elements")
	ErrMatSize = errors.New("bad size of matrix")
)

// Matrix is a M x N matrix
type Matrix struct {
	m, n int
	data []float64
}

// Zero matrix
func Zero(m, n int) *Matrix {
	return &Matrix{m: m, n: n, data: make([]float64, m*n)}
}

// Rand creates a size by size random matrix
func Rand(m, n int) *Matrix {
	A := Zero(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, rand.Float64())
		}
	}
	return A
}

// RandP creates a size by size random matrix concurrently
func RandP(m, n int) *Matrix {
	A := Zero(m, n)
	wg := sync.WaitGroup{}
	for i := 0; i < m; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < n; j++ {
				A.Set(i, j, rand.Float64())
			}
		}(i)
	}
	wg.Wait()
	return A
}

// NewP a size by size matrix concurrently
func NewP(m, n int) func(...float64) (*Matrix, error) {
	A := Zero(m, n)
	return func(es ...float64) (*Matrix, error) {
		if len(es) != m*n {
			return nil, ErrNumElem
		}
		// per row
		wg := sync.WaitGroup{}
		for i := 0; i < m; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := 0; j < n; j++ {
					A.Set(i, j, es[i*n+j])
				}
			}(i)
		}
		wg.Wait()
		return A, nil
	}
}

// New a size by size matrix
func New(m, n int) func(...float64) (*Matrix, error) {
	A := Zero(m, n)
	return func(es ...float64) (*Matrix, error) {
		if len(es) != m*n {
			return nil, ErrNumElem
		}
		// per row
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				A.Set(i, j, es[i*n+j])
			}
		}
		return A, nil
	}
}

// Print the matrix
func (A *Matrix) Print() {
	for i := 0; i < A.m; i++ {
		for j := 0; j < A.n; j++ {
			fmt.Printf("%.2f ", A.At(i, j))
		}
		fmt.Printf("\n")
	}
}

// Size of matrix
func (A *Matrix) Size() (int, int) {
	return A.m, A.n
}

// Row of matrix
func (A *Matrix) Row() int {
	return A.m
}

// Col of matrix
func (A *Matrix) Col() int {
	return A.n
}

// At access element (i, j)
func (A *Matrix) At(i, j int) float64 {
	return A.data[i*A.n+j]
}

// Set set element (i, j) with val
func (A *Matrix) Set(i, j int, val float64) {
	A.data[i*A.n+j] = val
}

// Inc adds element (i, j) with wal
func (A *Matrix) Inc(i, j int, val float64) {
	A.data[i*A.n+j] += val
}

// Mult multiple element (i, j) with wal
func (A *Matrix) Mult(i, j int, val float64) {
	A.data[i*A.n+j] *= val
}

// Pow computes power of n of element (i, j)
func (A *Matrix) Pow(i, j int, n float64) {
	A.data[i*A.n+j] = math.Pow(A.data[i*A.n+j], n)
}

// EqualShape check A.Size() == B.Size()
func (A *Matrix) EqualShape(B *Matrix) bool {
	am, an := A.Size()
	bm, bn := B.Size()
	if am != bm || an != bn {
		return false
	}
	return true
}

// Equal A and B?
func (A *Matrix) Equal(B *Matrix) bool {
	if !A.EqualShape(B) {
		return false
	}

	for i := 0; i < A.Row(); i++ {
		for j := 0; j < A.Col(); j++ {
			if A.At(i, j) != B.At(i, j) {
				return false
			}
		}
	}
	return true
}
