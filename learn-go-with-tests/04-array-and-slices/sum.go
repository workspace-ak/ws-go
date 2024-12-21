package sumarrays

// Add elements of slice and return single number
func Sum(a []int) int {
	res := 0
	for _, num := range a {
		res += num
	}
	return res
}

// Add elements of varied slice and return slice of sum of slices
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
