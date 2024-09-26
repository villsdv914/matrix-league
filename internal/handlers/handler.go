package handlers

import (
	"encoding/csv"
	"fmt"
	"league_matrix/internal/models"
	"league_matrix/internal/service"
	"league_matrix/internal/utils"
	"net/http"
	"strings"
)

type MatrixHandlerInterface interface {
	FileUpload(w http.ResponseWriter, request *http.Request)
	MatrixStringHandler(w http.ResponseWriter, request *http.Request)
	MatrixInvertHandler(w http.ResponseWriter, request *http.Request)
	MatrixFlattenHandler(w http.ResponseWriter, request *http.Request)
	MatrixSumHandler(w http.ResponseWriter, request *http.Request)
	MatrixMultiplyHandler(w http.ResponseWriter, request *http.Request)
}

type MatrixHandler struct {
	matrixService service.MatrixServiceInterface
}

var Matrix *models.MatrixModel

func NewMatrixHandler(matrixService service.MatrixServiceInterface) *MatrixHandler {
	return &MatrixHandler{
		matrixService: matrixService,
	}
}

func (mh *MatrixHandler) FileUpload(w http.ResponseWriter, request *http.Request) {
	file, _, err := request.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	responseValuetoint, err := utils.ConvertMatrixValuesToInt(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %v", err)))
		return
	}
	Matrix = &models.MatrixModel{Matrix: responseValuetoint}
	fmt.Fprint(w, response)
}

func (mh *MatrixHandler) MatrixStringHandler(w http.ResponseWriter, request *http.Request) {
	if Matrix == nil {
		w.Write([]byte(fmt.Sprintf("error: Please upload matrix csv file")))
		return
	}
	response := mh.matrixService.MatrixConvertToString(Matrix.Matrix)
	fmt.Fprint(w, response)
}

func (mh *MatrixHandler) MatrixInvertHandler(w http.ResponseWriter, request *http.Request) {
	if Matrix == nil {
		w.Write([]byte(fmt.Sprintf("error: Please upload matrix csv file")))
		return
	}
	response := mh.matrixService.MatrixInvert(Matrix.Matrix)
	fmt.Fprint(w, response)
}

func (mh *MatrixHandler) MatrixFlattenHandler(w http.ResponseWriter, request *http.Request) {
	if Matrix == nil {
		w.Write([]byte(fmt.Sprintf("error: Please upload matrix csv file")))
		return
	}
	response := mh.matrixService.MatrixFlatten(Matrix.Matrix)
	fmt.Fprint(w, response)
}

func (mh *MatrixHandler) MatrixSumHandler(w http.ResponseWriter, request *http.Request) {
	if Matrix == nil {
		w.Write([]byte(fmt.Sprintf("error: Please upload matrix csv file")))
		return
	}
	response := mh.matrixService.MatrixSum(Matrix.Matrix)
	fmt.Fprint(w, response)
}

func (mh *MatrixHandler) MatrixMultiplyHandler(w http.ResponseWriter, request *http.Request) {
	if Matrix == nil {
		w.Write([]byte(fmt.Sprintf("error: Please upload matrix csv file")))
		return
	}
	response := mh.matrixService.MatrixMultiply(Matrix.Matrix)
	fmt.Fprint(w, response)
}
