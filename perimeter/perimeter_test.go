package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	validateResult := func(t *testing.T, got, expected float64) {
		t.Helper()
		if got != expected {
			t.Errorf("got %.2f, expected %.2f", got, expected)
		}
	}

	t.Run("counts perimeter for a valid perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		validateResult(t, got, want)
	})

}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("%#v:, got %.2f, expected %.2f", shape, got, expected)
		}
	}

	t.Run("counts area for a rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, rectangle, want)
	})

	t.Run("counts area for a circle area", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{5.0, 2.0}, want: 5},
	}

	t.Run("data driven tests", func(t *testing.T) {
		for _, test := range areaTests {
			t.Run(test.name, func(t *testing.T) {
				checkArea(t, test.shape, test.want)
			})
		}
	})

}
