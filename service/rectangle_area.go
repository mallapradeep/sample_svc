package service

import "github.com/sample_svc/models"

// RectangleArea is used to calculate the area of Rectangle
func RectangleArea(params models.Rectangle) int64 {
	return params.Length * params.Breadth
}
