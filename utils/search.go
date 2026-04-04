package utils

import "slices"

func Search(data1, data2 *[]string) {
	res1, res2 := make([]string, 0, len(*data1)), make([]string, 0, len(*data2))
	for _, elem := range *data1 {
		if !slices.Contains(*data2, elem) {
			res1 = append(res1, elem)
		}
	}
	for _, elem := range *data2 {
		if !slices.Contains(*data1, elem) {
			res2 = append(res2, elem)
		}
	}
	*data1 = res1
	*data2 = res2
}
