package math

import "testing"

type AddData struct {
	x, y int
	result int
}

type DivideData struct {
	x, y int
	result float64
}


func TestAdd(t *testing.T){

	// result := Add(1,5)

	// if result != 6 {
	// 	t.Errorf("Add(1,5) failed. Expected %d, got %d\n", 6, result)
	// } else {
	// 	t.Logf("Add(1,5) passed. Expected %d, got %d\n", 6, result)
	// }

	testData := []AddData {
		{1, 2, 3},
		{3, 5, 8},
		{7, 4, 11},
	}

	for _, datum := range testData {
		result := Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Add(%d, %d) failed. Expected %d got %d\n", datum.x, datum.y, datum.result, result)
		}
	}
}

func TestDivide(t *testing.T){

	// result := Divide(5,0)

	// if result != 0 {
	// 	t.Errorf("Divide(5,0) failed. Expected %f, got %f\n", 0.0, result)
	// } else {
	// 	t.Logf("Divide(5,0) passed. Expected %f, got %f\n", 0.0, result)
	// }

	testData := []DivideData {
		{6, 6, 1.0},
		{50, -5, -10},
		{28, 3, 9.0},
	}

	for _, datum := range testData {
		result := Divide(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Divide(%d, %d) failed. Expected %f got %f\n", datum.x, datum.y, datum.result, result)
		}
	}
}

