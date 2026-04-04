package utils

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func SaveDifferences(filename string, data1, data2 []string) error {
	content := strings.Builder{}
	leftHeader := "MPRM160D_v7.1.bom"
	rightHeader := "выгрузка.txt"

	if len(data1) == 0 && len(data2) == 0 {
		if err := os.WriteFile(filename, []byte{}, 0644); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
		return nil
	}

	leftWidth := maxColumnWidth(leftHeader, data1)
	rightWidth := maxColumnWidth(rightHeader, data2)
	maxLen := len(data1)
	if len(data2) > maxLen {
		maxLen = len(data2)
	}
	content.WriteString("\t\t\t\t\t\tОТЛИЧИЯ\n")
	numberWidth := len(fmt.Sprintf("%d.", maxLen))
	content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, "", centerText(leftHeader, leftWidth), centerText(rightHeader, rightWidth)))

	for i := 0; i < maxLen; i++ {
		var left, right string
		if i < len(data1) {
			left = data1[i]
		}
		if i < len(data2) {
			right = data2[i]
		}
		content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, fmt.Sprintf("%d.", i+1), centerText(left, leftWidth), centerText(right, rightWidth)))
	}

	if err := os.WriteFile(filename, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return nil
}

func maxColumnWidth(header string, data []string) int {
	width := utf8.RuneCountInString(header)
	for _, elem := range data {
		elemWidth := utf8.RuneCountInString(elem)
		if elemWidth > width {
			width = elemWidth
		}
	}
	return width + 6
}

func centerText(text string, width int) string {
	textWidth := utf8.RuneCountInString(text)
	if textWidth >= width {
		return text
	}

	padding := width - textWidth
	left := padding / 2
	right := padding - left

	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}
