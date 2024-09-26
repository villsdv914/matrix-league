package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func ConvertMatrixValuesString(matrixVals [][]int) [][]string {
	var strMatrix [][]string
	for _, matrixRow := range matrixVals {
		var tempMatrix []string
		for _, row := range matrixRow {
			tempMatrix = append(tempMatrix, strconv.Itoa(row))
		}
		strMatrix = append(strMatrix, tempMatrix)
	}
	return strMatrix
}

func ConvertMatrixValuesToInt(matrixVals [][]string) ([][]int, error) {
	var intMatrix [][]int
	for _, matrixRow := range matrixVals {
		var tempMatrix []int
		for _, row := range matrixRow {
			val, err := strconv.Atoi(row)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Error converting value to int: %s", row))
			}
			tempMatrix = append(tempMatrix, val)
		}
		intMatrix = append(intMatrix, tempMatrix)
	}
	return intMatrix, nil
}
