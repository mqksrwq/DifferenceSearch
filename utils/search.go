package utils

func Search(data1, data2 *[]string) {
	res1, res2 := make([]string, 0, len(*data1)), make([]string, 0, len(*data2))

	lookup1 := make(map[string]struct{}, len(*data1))
	lookup2 := make(map[string]struct{}, len(*data2))

	for _, elem := range *data1 {
		lookup1[elem] = struct{}{}
	}
	for _, elem := range *data2 {
		lookup2[elem] = struct{}{}
	}

	for _, elem := range *data1 {
		if _, ok := lookup2[elem]; !ok {
			res1 = append(res1, elem)
		}
	}
	for _, elem := range *data2 {
		if _, ok := lookup1[elem]; !ok {
			res2 = append(res2, elem)
		}
	}

	*data1 = res1
	*data2 = res2
}
