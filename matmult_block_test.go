// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"fmt"
	"math"
	"testing"
)

func TestMatrix_MultBlock(t *testing.T) {
	A, err := New(5, 8)(
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 8, 7, 6, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 8, 7, 6, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 6, 7, 8,
	)
	if err != nil {
		t.Errorf("New(5, 8) error, expect nil")
	}
	B, err := New(8, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		2, 3, 4,
		5, 6, 7,
		8, 9, 8,
		7, 6, 5,
		4, 3, 2,
	)
	if err != nil {
		t.Errorf("New(8, 3) error, expect nil")
	}

	t.Run("DotBlockIJK()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockIJK(3, B, C)
		if err != nil {
			t.Errorf("DotBlockIJK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIJK() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockIKJ()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockIKJ(3, B, C)
		if err != nil {
			t.Errorf("DotBlockIKJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIKJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJIK()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockJIK(3, B, C)
		if err != nil {
			t.Errorf("DotBlockJIK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJIK() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJKI()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockJKI(3, B, C)
		if err != nil {
			t.Errorf("DotBlockJKI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJKI() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKIJ()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockKIJ(3, B, C)
		if err != nil {
			t.Errorf("DotBlockKIJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKIJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKJI()", func(t *testing.T) {
		C, err := New(5, 3)(
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		)
		T, err := New(5, 3)(
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
			248, 234, 212,
			192, 186, 168,
		)
		if err != nil {
			t.Errorf("New(5, 3) error, expect nil")
		}
		err = A.DotBlockKJI(3, B, C)
		if err != nil {
			t.Errorf("DotBlockKJI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKJI() not euqal, expect euqal, got:")
			C.Print()
		}
	})
}

// ----------------------- benchmarks ----------------------------

func BenchmarkBlock(b *testing.B) {
	for n := 0; n < 100; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		// L2 cache line: 256K
		// In core i7 the line sizes in L1 , L2 and L3 are the same: that is 64 Bytes.
		// see: https://stackoverflow.com/questions/14707803/line-size-of-l1-and-l2-caches
		// blockSize = sqrt(#CacheLines)
		nCacheLines := 256 * 1024 / 64
		blockSize := int(math.Sqrt(float64(nCacheLines)))
		b.Run(fmt.Sprintf("DotBlockIJK()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockIJK(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockIKJ()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockIKJ(blockSize, B, C)
			}
		})
	}
}
