package sumarrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{2, 4, 6, 8, 5}
	want := Sum(nums)
	expected := 25

	if want != expected {
		t.Errorf("wanted %d, but got %d, %v", want, expected, nums)
	}
}

func ExampleSum() {
	nums := []int{2, 4, 6, 8, 5}
	res := Sum(nums)
	fmt.Println(res)
	// Output: 25
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func ExampleSumAll() {
	res := SumAll([]int{1, 2, 3}, []int{4, 5})
	fmt.Printf("%#v", res)
	// Output: []int{6, 9}

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
	t.Run("sums of some slices tail", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5})
		want := []int{5, 5}
		checkSums(t, got, want)
	})

	t.Run("sums of empty slices tail", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	res := SumAllTails([]int{1, 2, 3}, []int{4, 5})
	fmt.Printf("%#v", res)
	// Output: []int{5, 5}

}
