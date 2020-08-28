package structs

import "testing"

func TestArea(t *testing.T) {
	testsArea := []struct {
		name   string
		form   Form
		expect float64
	}{
		{name: "Rectangle", form: Rectangle{12, 6}, expect: 72.0},
		{name: "Circle", form: Circle{10}, expect: 314.1592653589793},
		{name: "Triangle", form: Triangle{12, 6}, expect: 36.00},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()
			if result != tt.expect {
				t.Errorf("%#v result %.2f expect %.2f", tt.form, result, tt.expect)
			}
		})
	}
}
