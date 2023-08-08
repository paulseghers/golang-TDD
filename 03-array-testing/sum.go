package main

func Sum(numbers []int) int {
	sum := 0
	for _, k := range numbers {
		sum += k
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int
	for _, l := range numbersToSum {
		res = append(res, Sum(l))
	}
	return res
}

func SumAllTails(numbersToSum ...[]int) []int {
	var res []int
	for _, l := range numbersToSum {
		if len(l) == 0 {
			res = append(res, 0)
		} else {
			tail := l[1:]
			res = append(res, Sum(tail))
		}
	}
	return res
}
