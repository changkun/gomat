// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gomat

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"
// )

// type BlockMult func(bsize int, B, C *Matrix) error

// func TestMatrix_MultBlock(t *testing.T) {
// 	type args struct {
// 		blockSize int
// 		B         *Matrix
// 		C         *Matrix
// 	}
// 	origin := &Matrix{
// 		N: 8,
// 		data: [][]float64{
// 			[]float64{1, 2, 3, 4, 5, 6, 7, 8},
// 			[]float64{9, 1, 2, 3, 4, 5, 6, 7},
// 			[]float64{8, 9, 1, 2, 3, 4, 5, 6},
// 			[]float64{7, 8, 9, 1, 2, 3, 4, 5},
// 			[]float64{6, 7, 8, 9, 1, 2, 3, 4},
// 			[]float64{5, 6, 7, 8, 9, 1, 2, 3},
// 			[]float64{4, 5, 6, 7, 8, 9, 1, 2},
// 			[]float64{3, 4, 5, 6, 7, 8, 9, 0},
// 		},
// 	}
// 	tests := []struct {
// 		name    string
// 		f       BlockMult
// 		args    args
// 		wantErr bool
// 		truth   *Matrix
// 	}{
// 		{
// 			name: "8x8 ParalMultIJK",
// 			f:    origin.ParalMultIJK,
// 			args: args{
// 				blockSize: 3,
// 				B: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
// 						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
// 						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
// 						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
// 						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
// 						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
// 						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
// 						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
// 					},
// 				},
// 				C: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 			truth: &Matrix{
// 				N: 8,
// 				data: [][]float64{
// 					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
// 					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
// 					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
// 					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
// 					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
// 					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
// 					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
// 					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
// 				},
// 			},
// 		},
// 		{
// 			name: "8x8 ParalMultIKJ",
// 			f:    origin.ParalMultIKJ,
// 			args: args{
// 				blockSize: 2,
// 				B: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
// 						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
// 						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
// 						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
// 						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
// 						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
// 						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
// 						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
// 					},
// 				},
// 				C: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 			truth: &Matrix{
// 				N: 8,
// 				data: [][]float64{
// 					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
// 					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
// 					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
// 					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
// 					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
// 					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
// 					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
// 					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
// 				},
// 			},
// 		},
// 		{
// 			name: "8x8 MultIJK",
// 			f:    origin.MultIJK,
// 			args: args{
// 				blockSize: 2,
// 				B: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
// 						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
// 						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
// 						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
// 						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
// 						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
// 						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
// 						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
// 					},
// 				},
// 				C: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 			truth: &Matrix{
// 				N: 8,
// 				data: [][]float64{
// 					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
// 					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
// 					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
// 					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
// 					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
// 					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
// 					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
// 					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
// 				},
// 			},
// 		},
// 		{
// 			name: "8x8 MultIKJ",
// 			f:    origin.MultIKJ,
// 			args: args{
// 				blockSize: 2,
// 				B: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
// 						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
// 						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
// 						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
// 						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
// 						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
// 						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
// 						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
// 					},
// 				},
// 				C: &Matrix{
// 					N: 8,
// 					data: [][]float64{
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 			truth: &Matrix{
// 				N: 8,
// 				data: [][]float64{
// 					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
// 					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
// 					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
// 					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
// 					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
// 					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
// 					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
// 					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.f(tt.args.blockSize, tt.args.B, tt.args.C); (err != nil) != tt.wantErr {
// 				t.Errorf("Matrix.MultIJK() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !reflect.DeepEqual(tt.args.C, tt.truth) {
// 				t.Errorf("Matrix.MultIJK() = %v, want %v", tt.args.C, tt.truth)
// 			}
// 		})
// 	}
// }

// func BenchmarkMatrix_MultBlock(b *testing.B) {
// 	tests := []struct {
// 		name string
// 		f    BlockMult
// 	}{
// 		{
// 			name: "A.MultIJK",
// 			f:    A.MultIJK,
// 		},
// 		{
// 			name: "A.MultIKJ",
// 			f:    A.MultIKJ,
// 		},
// 		{
// 			name: "A.ParalMultIJK",
// 			f:    A.ParalMultIJK,
// 		},
// 		{
// 			name: "A.ParalMultIKJ",
// 			f:    A.ParalMultIKJ,
// 		},
// 	}
// 	for _, tt := range tests {
// 		b.Run(tt.name, func(b *testing.B) {
// 			for _, size := range bSize {
// 				b.Run(fmt.Sprintf("blockSize-%d", size), func(b *testing.B) {
// 					for i := 0; i < b.N; i++ {
// 						if err := tt.f(size, B, C); err != nil {
// 							b.Errorf("%s() error %v, want nil", tt.name, err)
// 						}
// 					}
// 				})
// 			}
// 		})
// 	}
// }
