package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("works with a size of 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})

	t.Run("works with a slice of 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})

	t.Run("works with a slice of 0", func(t *testing.T) {
		numbers := []int{}

		got := Sum(numbers)
		want := 0

		if want != got {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	t.Run("works with two arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test with normal array", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("test with empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{22})
		want := []int{0, 0}
		checkSum(t, got, want)
	})
}
