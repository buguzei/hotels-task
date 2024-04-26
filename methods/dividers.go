package methods

func Dividers(nums []int) []int {
	res := make([]int, 0, len(nums))
	dividersAmount := make(map[int]int)

	for _, num := range nums {
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				dividersAmount[i]++
			}
		}
	}

	for num, amount := range dividersAmount {
		if amount == len(nums) {
			res = append(res, num)
		}
	}

	return res
}
