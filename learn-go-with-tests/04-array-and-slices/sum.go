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

// Calculate totals of the tails of each slice and return slice of totals
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {
		if len(nums) > 0 {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
