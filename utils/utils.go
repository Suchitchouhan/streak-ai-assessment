package utils

func FindPair(numbers []int, target int) [][]int {
	pair := [][]int{}
	tmp_map := make(map[int]int)
	for index, num := range numbers {
		comp := target - num
		if j, ok := tmp_map[comp]; ok {
			pair = append(pair, []int{j, num})
		}
		tmp_map[num] = index
	}
	return pair
}
