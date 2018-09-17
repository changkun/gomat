package gommult

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size int
		arr  []float64
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			name: "3x3 matrix",
			args: args{
				size: 3,
				arr:  []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: &Matrix{
				N: 3,
				data: [][]float64{
					[]float64{1, 2, 3},
					[]float64{4, 5, 6},
					[]float64{7, 8, 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.size)(tt.args.arr...)
			if err != nil {
				t.Errorf("New() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_At(t *testing.T) {
	type fields struct {
		N    int
		data [][]float64
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "3x3 matrix",
			fields: fields{
				N: 3,
				data: [][]float64{
					[]float64{1, 2, 3},
					[]float64{4, 5, 6},
					[]float64{7, 8, 9},
				},
			},
			args: args{1, 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &Matrix{
				N:    tt.fields.N,
				data: tt.fields.data,
			}
			if got := A.At(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Matrix.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type fields struct {
		N    int
		data [][]float64
	}
	type args struct {
		i   int
		j   int
		val float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		{
			name: "3x3 matrix",
			fields: fields{
				N: 3,
				data: [][]float64{
					[]float64{1, 2, 3},
					[]float64{4, 5, 6},
					[]float64{7, 8, 9},
				},
			},
			args: args{1, 1, 6},
			want: &Matrix{
				N: 3,
				data: [][]float64{
					[]float64{1, 2, 3},
					[]float64{4, 6, 6},
					[]float64{7, 8, 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := &Matrix{
				N:    tt.fields.N,
				data: tt.fields.data,
			}
			A.Set(tt.args.i, tt.args.j, tt.args.val)
			if !reflect.DeepEqual(A, tt.want) {
				t.Errorf("Matrix.Set() = %v, want %v", A, tt.want)
			}
		})
	}
}

type NaiveMult func(B, C *Matrix) error
type BlockMult func(bsize int, B, C *Matrix) error

func TestMatrix_MultNaive(t *testing.T) {
	type args struct {
		B *Matrix
		C *Matrix
	}
	origin := &Matrix{
		N: 2,
		data: [][]float64{
			[]float64{1, 2},
			[]float64{2, 1},
		},
	}
	origin.debug()
	tests := []struct {
		name    string
		args    args
		wantErr bool
		f       NaiveMult
		truth   *Matrix
	}{
		{
			name: "2x2 MultNaive",
			args: args{
				B: &Matrix{
					N: 2,
					data: [][]float64{
						[]float64{2, 3},
						[]float64{3, 4},
					},
				},
				C: &Matrix{
					N: 2,
					data: [][]float64{
						[]float64{0, 0},
						[]float64{0, 0},
					},
				},
			},
			wantErr: false,
			f:       origin.MultNaive,
			truth: &Matrix{
				N: 2,
				data: [][]float64{
					[]float64{8, 11},
					[]float64{7, 10},
				},
			},
		},
		{
			name: "2x2 ParalMultNaive",
			args: args{
				B: &Matrix{
					N: 2,
					data: [][]float64{
						[]float64{2, 3},
						[]float64{3, 4},
					},
				},
				C: &Matrix{
					N: 2,
					data: [][]float64{
						[]float64{0, 0},
						[]float64{0, 0},
					},
				},
			},
			wantErr: false,
			f:       origin.ParalMultNaive,
			truth: &Matrix{
				N: 2,
				data: [][]float64{
					[]float64{8, 11},
					[]float64{7, 10},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f(tt.args.B, tt.args.C); (err != nil) != tt.wantErr {
				t.Errorf("%s() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.C, tt.truth) {
				t.Errorf("%s() = %v, want %v", tt.name, tt.args.C, tt.truth)
			}
		})
	}
}

func TestMatrix_MultBlock(t *testing.T) {
	type args struct {
		blockSize int
		B         *Matrix
		C         *Matrix
	}
	origin := &Matrix{
		N: 8,
		data: [][]float64{
			[]float64{1, 2, 3, 4, 5, 6, 7, 8},
			[]float64{9, 1, 2, 3, 4, 5, 6, 7},
			[]float64{8, 9, 1, 2, 3, 4, 5, 6},
			[]float64{7, 8, 9, 1, 2, 3, 4, 5},
			[]float64{6, 7, 8, 9, 1, 2, 3, 4},
			[]float64{5, 6, 7, 8, 9, 1, 2, 3},
			[]float64{4, 5, 6, 7, 8, 9, 1, 2},
			[]float64{3, 4, 5, 6, 7, 8, 9, 0},
		},
	}
	tests := []struct {
		name    string
		f       BlockMult
		args    args
		wantErr bool
		truth   *Matrix
	}{
		{
			name: "8x8 ParalMultIJK",
			f:    origin.ParalMultIJK,
			args: args{
				blockSize: 3,
				B: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
					},
				},
				C: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			wantErr: false,
			truth: &Matrix{
				N: 8,
				data: [][]float64{
					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
				},
			},
		},
		{
			name: "8x8 ParalMultIKJ",
			f:    origin.ParalMultIKJ,
			args: args{
				blockSize: 2,
				B: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
					},
				},
				C: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			wantErr: false,
			truth: &Matrix{
				N: 8,
				data: [][]float64{
					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
				},
			},
		},
		{
			name: "8x8 MultIJK",
			f:    origin.MultIJK,
			args: args{
				blockSize: 2,
				B: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
					},
				},
				C: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			wantErr: false,
			truth: &Matrix{
				N: 8,
				data: [][]float64{
					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
				},
			},
		},
		{
			name: "8x8 MultIKJ",
			f:    origin.MultIKJ,
			args: args{
				blockSize: 2,
				B: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{9, 8, 7, 6, 5, 4, 3, 2},
						[]float64{1, 9, 8, 7, 6, 5, 4, 3},
						[]float64{2, 1, 9, 8, 7, 6, 5, 4},
						[]float64{3, 2, 1, 9, 8, 7, 6, 5},
						[]float64{4, 3, 2, 1, 9, 8, 7, 6},
						[]float64{5, 4, 3, 2, 1, 9, 8, 7},
						[]float64{6, 5, 4, 3, 2, 1, 9, 8},
						[]float64{7, 6, 5, 4, 3, 2, 1, 0},
					},
				},
				C: &Matrix{
					N: 8,
					data: [][]float64{
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
						[]float64{0, 0, 0, 0, 0, 0, 0, 0},
					},
				},
			},
			wantErr: false,
			truth: &Matrix{
				N: 8,
				data: [][]float64{
					[]float64{177, 159, 150, 150, 159, 177, 204, 168},
					[]float64{221, 193, 174, 164, 163, 171, 188, 151},
					[]float64{193, 236, 207, 187, 176, 174, 181, 143},
					[]float64{174, 207, 249, 219, 198, 186, 183, 144},
					[]float64{164, 187, 219, 260, 229, 207, 194, 154},
					[]float64{163, 176, 198, 229, 269, 237, 214, 173},
					[]float64{171, 174, 186, 207, 237, 276, 243, 201},
					[]float64{181, 175, 178, 190, 211, 241, 280, 238},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f(tt.args.blockSize, tt.args.B, tt.args.C); (err != nil) != tt.wantErr {
				t.Errorf("Matrix.MultIJK() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.C, tt.truth) {
				t.Errorf("Matrix.MultIJK() = %v, want %v", tt.args.C, tt.truth)
			}
		})
	}
}

// ----------------------- benchmarks ----------------------------

var (
	bSize = []int{1, 2, 3, 4, 5, 6, 7, 8}
	A     = &Matrix{
		N: 8,
		data: [][]float64{
			[]float64{1, 2, 3, 4, 5, 6, 7, 8},
			[]float64{9, 1, 2, 3, 4, 5, 6, 7},
			[]float64{8, 9, 1, 2, 3, 4, 5, 6},
			[]float64{7, 8, 9, 1, 2, 3, 4, 5},
			[]float64{6, 7, 8, 9, 1, 2, 3, 4},
			[]float64{5, 6, 7, 8, 9, 1, 2, 3},
			[]float64{4, 5, 6, 7, 8, 9, 1, 2},
			[]float64{3, 4, 5, 6, 7, 8, 9, 0},
		},
	}
	B = &Matrix{
		N: 8,
		data: [][]float64{
			[]float64{9, 8, 7, 6, 5, 4, 3, 2},
			[]float64{1, 9, 8, 7, 6, 5, 4, 3},
			[]float64{2, 1, 9, 8, 7, 6, 5, 4},
			[]float64{3, 2, 1, 9, 8, 7, 6, 5},
			[]float64{4, 3, 2, 1, 9, 8, 7, 6},
			[]float64{5, 4, 3, 2, 1, 9, 8, 7},
			[]float64{6, 5, 4, 3, 2, 1, 9, 8},
			[]float64{7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	C = &Matrix{
		N: 8,
		data: [][]float64{
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
)

func BenchmarkMatrix_MultNaive(b *testing.B) {
	tests := []struct {
		name string
		f    NaiveMult
	}{
		{
			name: "A.MultNaive",
			f:    A.MultNaive,
		},
		{
			name: "A.ParalMultNaive",
			f:    A.ParalMultNaive,
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := tt.f(B, C); err != nil {
					b.Errorf("%s() error %v, want nil", tt.name, err)
				}
			}
		})
	}
}
func BenchmarkMatrix_MultBlock(b *testing.B) {
	tests := []struct {
		name string
		f    BlockMult
	}{
		{
			name: "A.MultIJK",
			f:    A.MultIJK,
		},
		{
			name: "A.MultIKJ",
			f:    A.MultIKJ,
		},
		{
			name: "A.ParalMultIJK",
			f:    A.ParalMultIJK,
		},
		{
			name: "A.ParalMultIKJ",
			f:    A.ParalMultIKJ,
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for _, size := range bSize {
				b.Run(fmt.Sprintf("blockSize-%d", size), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						if err := tt.f(size, B, C); err != nil {
							b.Errorf("%s() error %v, want nil", tt.name, err)
						}
					}
				})
			}
		})
	}
}
