package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 5}

		got := Sum(numbers)
		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{2, 4})
	want := []int{4, 6}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 4})
		want := []int{0, 6}

		checkSums(t, got, want)
	})
}
