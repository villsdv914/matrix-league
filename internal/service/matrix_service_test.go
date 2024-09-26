package service

import "testing"

func TestMatrixService_MatrixInvert(t *testing.T) {
	type args struct {
		matrixSlc [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{matrixSlc: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: "1,4,7\n2,5,8\n3,6,9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MatrixService{}
			if got := m.MatrixInvert(tt.args.matrixSlc); got != tt.want {
				t.Errorf("MatrixInvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixService_MatrixFlatten(t *testing.T) {
	type args struct {
		matrixSlc [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{matrixSlc: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: "1,2,3,4,5,6,7,8,9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MatrixService{}
			if got := m.MatrixFlatten(tt.args.matrixSlc); got != tt.want {
				t.Errorf("MatrixFlatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
