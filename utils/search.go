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

func SkipMismatchedPartNumbers(data1, data2 *[]string, parts1, parts2 map[string]string) {
	if len(parts1) == 0 || len(parts2) == 0 {
		return
	}

	skip := make(map[string]struct{})
	for ref, part1 := range parts1 {
		if part2, ok := parts2[ref]; ok && part1 != "" && part2 != "" && part1 != part2 {
			skip[ref] = struct{}{}
		}
	}
	if len(skip) == 0 {
		return
	}

	filter := func(data []string) []string {
		out := make([]string, 0, len(data))
		for _, elem := range data {
			if _, shouldSkip := skip[elem]; shouldSkip {
				continue
			}
			out = append(out, elem)
		}
		return out
	}

	*data1 = filter(*data1)
	*data2 = filter(*data2)
}
