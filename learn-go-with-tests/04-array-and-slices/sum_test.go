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
	want := SumAll([]int{1, 2, 3}, []int{4, 5})
	expect := []int{6, 9}
	if !reflect.DeepEqual(expect, want) {
		t.Errorf("wanted %v, but got %v", expect, want)
	}
}

func ExampleSumAll() {
	res := SumAll([]int{1, 2, 3}, []int{4, 5})
	fmt.Printf("%#v", res)
	// Output: []int{6, 9}

}
