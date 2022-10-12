package arraysandslices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {

	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)

	// var balance float64
	// for _, t := range transactions {
	// 	if t.From == name {
	// 		balance -= t.Sum
	// 	}
	// 	if t.To == name {
	// 		balance += t.Sum
	// 	}
	// }
	// return balance
}

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// func Sum(numbers []int) int {
// 	sum := 0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails calculates the sums of all but the first
// number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		if len(numbers) == 0 {
// 			sums = append(sums, 0)
// 		} else {
// 			tail := numbers[1:]
// 			sums = append(sums, Sum(tail))
// 		}
// 	}
// 	return sums
// }

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
