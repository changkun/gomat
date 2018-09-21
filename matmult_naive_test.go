// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

import (
	"fmt"
	"testing"
)

func TestMatrix_DotNaive(t *testing.T) {
	A, err := New(2, 3)(
		1, 2, 3,
		2, 3, 1,
	)
	if err != nil {
		t.Errorf("New(2, 3) error, expect nil")
	}
	B, err := New(3, 1)(
		3,
		2,
		1,
	)
	if err != nil {
		t.Errorf("New(3, 1) error, expect nil")
	}
	T, err := New(2, 1)(
		10,
		13,
	)
	if err != nil {
		t.Errorf("New(2, 1) error, expect nil")
	}
	t.Run("DotNaive()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaive(B, C)
		if err != nil {
			t.Errorf("DotNaive(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaive() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveP(B, C)
		if err != nil {
			t.Errorf("DotNaiveP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIJK()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveIJK(B, C)
		if err != nil {
			t.Errorf("DotNaiveIJK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIJK() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIKJ()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveIKJ(B, C)
		if err != nil {
			t.Errorf("DotNaiveIKJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIKJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJIK()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveJIK(B, C)
		if err != nil {
			t.Errorf("DotNaiveJIK(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJIK() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJKI()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveJKI(B, C)
		if err != nil {
			t.Errorf("DotNaiveJKI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJKI() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKIJ()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveKIJ(B, C)
		if err != nil {
			t.Errorf("DotNaiveKIJ(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKIJ() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKJI()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveKJI(B, C)
		if err != nil {
			t.Errorf("DotNaiveKJI(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKJI() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIJKP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveIJKP(B, C)
		if err != nil {
			t.Errorf("DotNaiveIJKP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIJKP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveIKJP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveIKJP(B, C)
		if err != nil {
			t.Errorf("DotNaiveIKJP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveIKJP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJIKP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveJIKP(B, C)
		if err != nil {
			t.Errorf("DotNaiveJIKP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJIKP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveJKIP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveJKIP(B, C)
		if err != nil {
			t.Errorf("DotNaiveJKIP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveJKIP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKIJP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveKIJP(B, C)
		if err != nil {
			t.Errorf("DotNaiveKIJP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKIJP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
	t.Run("DotNaiveKJIP()", func(t *testing.T) {
		C := Zero(2, 1)
		err := A.DotNaiveKJIP(B, C)
		if err != nil {
			t.Errorf("DotNaiveKJIP(B, C) error, expect nil")
		}
		if !T.Equal(C) {
			t.Errorf("DotNaiveKJIP() not euqal, expect euqal, got:")
			C.Print()
		}
	})
}

// ----------------------- benchmarks ----------------------------

func BenchmarkNaive(b *testing.B) {
	for n := 80; n < 100; n++ {
		A := Rand(n, n)
		B := Rand(n, n)
		C := Zero(n, n)

		// vanilla
		b.Run(fmt.Sprintf("DotNaiveIJK() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIJK(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveIKJ()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIKJ(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJIK()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJIK(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJKI()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJKI(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKIJ()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKIJ(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKJI()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKJI(B, C)
			}
		})

		// concurrency
		b.Run(fmt.Sprintf("DotNaiveIJKP() %dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIJKP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveIKJP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveIKJP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJIKP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJIKP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveJKIP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveJKIP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKIJP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKIJP(B, C)
			}
		})
		b.Run(fmt.Sprintf("DotNaiveKJIP()-%dx%d", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				A.DotNaiveKJIP(B, C)
			}
		})
	}
}
