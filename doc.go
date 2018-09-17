// Package gommult is a golang package for NxN matrix multiplication with blocking optimization
//
// Prior knowledge
//
// Assume only 2 levels in the hierarchy, fast(registers/cache) and slow(main memory).
// All data initially in slow memory
//
//   - m:       number of memory elements (words) moved between fast and slow memory
//   - t_m:     time per slow memory operation
//   - f:       number of arithemetic operations
//   - t_f:     time per arithmetic operation (t_f << t_m)
//   - q = f/m: computational intensity (key to algorithm efficiency) average number
//              of flops per slow memory access
//
// Minimum possible time = f * t_f when all data in fast memory.
// Actual time = f * t_f + m * t_m = f * t_f * [1 + (t_m / t_f) * (1 / q)]
// Machine balance a = t_m / t_f (key to machine efficiency)
// Larger q means time closer to minimum f*t_f
//   - q >= t_m / t_f needed to get at least half of peak speed
//
// Blocked (Tiled) Matrix Multiply
//
// Consider A,B,C to be N-by-N matrices of b-by-b sub-blocks
//   - b = n / N is called the *block size*
//
package gommult
