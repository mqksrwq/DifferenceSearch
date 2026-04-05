package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Parse(file *os.File) []string {
	fileString, err := fileToString(file)
	var data []string

	if err != nil {
		fmt.Println(err)
	}
	switch filepath.Ext(file.Name()) {
	case ".bom":
		data = serializeBom(fileString)
		unpacking(&data)
	default:
		data = serializeTxt(fileString)
		unpacking(&data)
	}
	return data
}

func fileToString(file *os.File) (string, error) {
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return "", errors.New("ошибка чтения файла")
	}
	return string(data), nil
}

func serializeBom(file string) []string {
	var res []string
	rows := strings.Split(file, "\n")
	for _, row := range rows[2 : len(rows)-1] {
		value := strings.TrimSpace(strings.Trim(strings.Split(row, "|")[2], `""`))
		if value == "" {
			continue
		}
		res = append(res, value)
	}
	return res
}

func serializeTxt(file string) []string {
	var res []string

	rows := strings.Split(file, "\n")
	for _, row := range rows {
		before, _, _ := strings.Cut(row, "\t")
		value := strings.TrimSpace(before)
		if value == "" {
			continue
		}
		res = append(res, value)
	}
	return res
}

func unpacking(data *[]string) {
	result := make([]string, 0, len(*data))

	for _, elem := range *data {
		elem = strings.TrimSpace(elem)
		if elem == "" {
			continue
		}

		if strings.Contains(elem, ",") {

			for _, elem := range strings.Split(elem, ",") {
				var chStart, numberStr string
				elem = strings.TrimSpace(elem)

				for _, ch := range elem {
					if ch >= '0' && ch <= '9' {
						numberStr += string(ch)
					} else {
						chStart += string(ch)
					}
				}

				s := chStart + numberStr
				result = append(result, s)
			}
		} else if strings.Contains(elem, "...") {
			before, after, _ := strings.Cut(elem, "...")

			var start, end string
			var chStart string

			for _, ch := range before {
				if ch >= '0' && ch <= '9' {
					start += string(ch)
				} else {
					chStart += string(ch)
				}
			}
			for _, ch := range after {
				if ch >= '0' && ch <= '9' {
					end += string(ch)
				}
			}

			s1, _ := strconv.Atoi(start)
			s2, _ := strconv.Atoi(end)

			for i := s1; i <= s2; i++ {
				st := chStart + strconv.Itoa(i)
				result = append(result, st)
			}
		} else {
			result = append(result, elem)
		}
	}
	*data = result
}
