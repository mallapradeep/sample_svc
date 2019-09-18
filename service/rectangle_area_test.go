package service

import (
	"testing"

	"github.com/sample_svc/models"
)

func TestArea(t *testing.T) {
	type args struct {
		params models.Rectangle
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "Success, Got the area of Rectangle",
			args: args{params: models.Rectangle{Length: 10, Breadth: 6}},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RectangleArea(tt.args.params); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
