package utils

import (
	"fmt"
	"io"
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
	default:
		data = serializeTxt(fileString)
	}
	unpacking(&data)
	return data
}

func fileToString(file *os.File) (string, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return "", fmt.Errorf("ошибка позиционирования файла: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return string(data), nil
}

func serializeBom(file string) []string {
	var res []string
	rows := strings.Split(file, "\n")
	if len(rows) <= 2 {
		return res
	}

	for _, row := range rows[2:] {
		parts := strings.Split(row, "|")
		if rowHasRepZnak(parts) {
			continue
		}
		if len(parts) < 5 || strings.ToUpper(strings.TrimSpace(strings.Trim(parts[4], `""`))) == "NOT" {
			continue
		}

		value := strings.ToUpper(strings.TrimSpace(strings.Trim(parts[2], `""`)))
		if value == "" {
			continue
		}
		res = append(res, value)
	}
	return res
}

func rowHasRepZnak(parts []string) bool {
	for _, cell := range parts {
		value := strings.ToUpper(strings.TrimSpace(strings.Trim(cell, `""`)))
		if value == "REP_ZNAK" {
			return true
		}
	}
	return false
}

func serializeTxt(file string) []string {
	var res []string

	rows := strings.Split(file, "\n")
	for _, row := range rows {
		before, _, _ := strings.Cut(row, "\t")
		value := strings.ToUpper(strings.TrimSpace(before))
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
		for _, token := range strings.Split(elem, ",") {
			token = strings.ToUpper(strings.TrimSpace(token))
			if token == "" {
				continue
			}

			if strings.Contains(token, "...") {
				before, after, _ := strings.Cut(token, "...")
				prefix, start := splitPrefixAndNumber(before)
				_, end := splitPrefixAndNumber(after)

				s1, err1 := strconv.Atoi(start)
				s2, err2 := strconv.Atoi(end)
				if err1 != nil || err2 != nil || s1 > s2 {
					result = append(result, token)
					continue
				}

				for i := s1; i <= s2; i++ {
					result = append(result, prefix+strconv.Itoa(i))
				}
				continue
			}

			prefix, number := splitPrefixAndNumber(token)
			if number == "" {
				result = append(result, token)
				continue
			}
			result = append(result, prefix+number)
		}
	}
	*data = result
}

func splitPrefixAndNumber(s string) (string, string) {
	var prefixBuilder strings.Builder
	var numberBuilder strings.Builder

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			numberBuilder.WriteRune(ch)
			continue
		}
		prefixBuilder.WriteRune(ch)
	}

	return prefixBuilder.String(), numberBuilder.String()
}
