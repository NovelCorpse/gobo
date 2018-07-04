package gobo

import (
	"time"
	"errors"
	"fmt"
	"math/rand"
)

type Matrix [][]float64

func NewMatrix(M uint, N uint) (Matrix, error) {
	if M == 0 || N == 0 {
		return nil, errors.New("Incorrect input, M and N must be more then zero")
	}
	var matrix Matrix
	for m := 0; m < int(M); m++ {
		mstring := make([]float64, N)
		matrix = append(matrix, mstring)
	}
	return matrix, nil
}

func (M Matrix) String() string {
	var s string
	for _, m := range M {
		s = fmt.Sprintf("%s\n%9.2f", s, m)
	}
	return s
}

func (M Matrix) GetSize() (uint, uint) {
	m := uint(len(M))
	n := uint(len(M[0]))
	return m, n
}

func (M Matrix) Transpose() Matrix {
	m, n := M.GetSize()
	newMatrix, _ := NewMatrix(n, m)
	for i := range M {
		for j := range M[i] {
			newMatrix[j][i] = M[i][j]
		}
	}
	return newMatrix
}

func (M Matrix) Multiplie(l float64) Matrix {
	newMatrix, _ := NewMatrix(M.GetSize())
	for i := range M {
		for j := range M[i] {
			newMatrix[i][j] = M[i][j] * l
		}
	}
	return newMatrix
}

func (M Matrix) Randomize() {
	
	for i := range M {
		for j := range M[i] {
			M[i][j] = rand.Float64()
		}
	}
}

func init(){
	rand.Seed(int64(time.Now().Second()))
}