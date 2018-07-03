package gobo

import (
	"fmt"
)

type Matrix [][]float64

func NewMatrix(M uint, N uint) Matrix {
	var matrix Matrix
	for m := 0; m < int(M); m++ {
		string := make([]float64, N)
		matrix = append(matrix, string)
	}
	return matrix
}

func (M Matrix) String() string {
	var s string
	for _, m := range M {
		
		s = fmt.Sprintf("%s\n%v", s, m)
		
	}
	return s
}