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

	t.Run("DotBlock()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlock(B, C)
		if err != nil {
			t.Errorf("DotBlock(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlock() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotBlockP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockP(B, C)
		if err != nil {
			t.Errorf("DotBlockP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockIJK()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockIJK(2, B, C)
		if err != nil {
			t.Errorf("DotBlockIJK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIJK() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockIKJ()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockIKJ(2, B, C)
		if err != nil {
			t.Errorf("DotBlockIKJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIKJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJIK()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockJIK(2, B, C)
		if err != nil {
			t.Errorf("DotBlockJIK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJIK() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJKI()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockJKI(2, B, C)
		if err != nil {
			t.Errorf("DotBlockJKI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJKI() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKIJ()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockKIJ(2, B, C)
		if err != nil {
			t.Errorf("DotBlockKIJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKIJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKJI()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockKJI(2, B, C)
		if err != nil {
			t.Errorf("DotBlockKJI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKJI() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockIJKP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockIJKP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockIJKP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIJKP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockIKJP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockIKJP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockIKJP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockIKJP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJIKP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockJIKP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockJIKP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJIKP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockJKIP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockJKIP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockJKIP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockJKIP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKIJP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockKIJP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockKIJP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKIJP() not euqal, expect euqal, got:")
			C.Print()
		}
	})

	t.Run("DotBlockKJIP()", func(t *testing.T) {
		C := Zero(5, 3)
		err := A.DotBlockKJIP(2, B, C)
		if err != nil {
			t.Errorf("DotBlockKJIP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotBlockKJIP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
}

// ----------------------- benchmarks ----------------------------

func BenchmarkBlock(b *testing.B) {
	for n := 80; n < 130; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		// L2 cache line: 256K
		// In core i7 the line sizes in L1 , L2 and L3 are the same: that is 64 Bytes.
		// see: https://stackoverflow.com/questions/14707803/line-size-of-l1-and-l2-caches
		// blockSize = sqrt(#CacheLines/3)
		nCacheLines := 256 * 1024 / 64
		blockSize := int(math.Sqrt(float64(nCacheLines) / 3))

		// vanilla
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
		b.Run(fmt.Sprintf("DotBlockJIK()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockJIK(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockJKI()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockJKI(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockKIJ()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockKIJ(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockKJI()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockKJI(blockSize, B, C)
			}
		})

		// concurrency
		b.Run(fmt.Sprintf("DotBlockIJKP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockIJKP(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockIKJP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockIKJP(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockJIKP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockJIKP(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockJKIP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockJKIP(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockKIJP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockKIJP(blockSize, B, C)
			}
		})
		b.Run(fmt.Sprintf("DotBlockKJIP()-block-size-%d-%dx%d", blockSize, n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotBlockKJIP(blockSize, B, C)
			}
		})
	}
}
