package service

import (
	"math/big"
	"strings"
)
import "league_matrix/internal/utils"

type MatrixServiceInterface interface {
	MatrixConvertToString(matrixSlc [][]int) string
	MatrixInvert(matrixSlc [][]int) string
	MatrixFlatten(matrixSlc [][]int) string
	MatrixSum(matrixSlc [][]int) int
	MatrixMultiply(matrixSlc [][]int) *big.Int
}

type MatrixService struct{}

func NewMatrixService() MatrixServiceInterface {
	return &MatrixService{}
}

func (m *MatrixService) MatrixConvertToString(matrixSlc [][]int) string {
	strMatrix := utils.ConvertMatrixValuesString(matrixSlc)
	var resultMatrix string
	for idx, matrix := range strMatrix {
		resultMatrix += strings.Join(matrix, ",")
		if idx != len(matrixSlc)-1 {
			resultMatrix += "\n"
		}
	}
	return resultMatrix
}

func (m *MatrixService) MatrixInvert(matrixSlc [][]int) string {
	invertMatrix := make([][]int, len(matrixSlc))
	for i := 0; i < len(matrixSlc); i++ {
		invertMatrix[i] = make([]int, len(matrixSlc[i]))
		for j := 0; j < len(matrixSlc[i]); j++ {
			invertMatrix[i][j] = matrixSlc[j][i]
		}
	}
	strMatrix := utils.ConvertMatrixValuesString(invertMatrix)
	var resultMatrix string
	for idx, matrix := range strMatrix {
		resultMatrix += strings.Join(matrix, ",")
		if idx != len(invertMatrix)-1 {
			resultMatrix += "\n"
		}

	}
	return resultMatrix
}

func (m *MatrixService) MatrixFlatten(matrixSlc [][]int) string {
	strMatrix := utils.ConvertMatrixValuesString(matrixSlc)
	var resultMatrix string
	for idx, matrix := range strMatrix {
		resultMatrix += strings.Join(matrix, ",")
		if idx != len(matrixSlc)-1 {
			resultMatrix += ","
		}
	}
	return resultMatrix
}

func (m *MatrixService) MatrixSum(matrixSlc [][]int) int {
	var resultMatrixSum int
	for _, matrix := range matrixSlc {
		sum := 0
		for _, matNum := range matrix {
			sum += matNum
		}
		resultMatrixSum += sum
	}
	return resultMatrixSum
}

// multiplicaton will hold big number , for that we have to use math/big package
func (m *MatrixService) MatrixMultiply(matrixSlc [][]int) *big.Int {
	resultMatrixMultiply := big.NewInt(1)
	for _, matrix := range matrixSlc {
		for _, matNum := range matrix {
			resultMatrixMultiply.Mul(resultMatrixMultiply, big.NewInt(int64(matNum)))
		}
	}
	return resultMatrixMultiply
}
