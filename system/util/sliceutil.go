package util

func RemoveDuplicates(d []string) []string {

	check := make(map[string]int)

	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}
