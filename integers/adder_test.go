package integers_test

import (
	"github.com/nithyanatarajan/go-with-tests/integers"
	. "github.com/nithyanatarajan/go-with-tests/test"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("should add two numbers", func(t *testing.T) {
		want := 5
		got := integers.Add(2, 3)
		AssertEqualInt(t, got, want)
	})
}

func TestSum(t *testing.T) {
	t.Run("should return 0 when no numbers are given", func(t *testing.T) {
		want := 0
		got := integers.Sum()
		AssertEqualInt(t, got, want)
	})

	t.Run("should return same number when only one number is given", func(t *testing.T) {
		want := 10
		got := integers.Sum(10)
		AssertEqualInt(t, got, want)
	})

	t.Run("should add ten numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 55
		got := integers.Sum(numbers...)
		AssertEqualInt(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should return empty array when no arrays are given", func(t *testing.T) {
		var want []int
		got := integers.SumAll()
		AssertEqualIntArray(t, got, want)
	})

	t.Run("should return sum of given array", func(t *testing.T) {
		want := []int{3}
		got := integers.SumAll([]int{1, 1, 1})
		AssertEqualIntArray(t, got, want)
	})

	t.Run("should return sum of all given arrays", func(t *testing.T) {
		want := []int{3, 9}
		got := integers.SumAll([]int{1, 2}, []int{0, 9})
		AssertEqualIntArray(t, got, want)
	})
}
