package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode/utf8"
)

func SaveDifferences(filename, leftHeader, rightHeader string, data1, data2 []string, attentionSide int, attentionData []string) error {
	content := strings.Builder{}

	if len(data1) == 0 && len(data2) == 0 && len(attentionData) == 0 {
		if err := os.WriteFile(filename, []byte{}, 0644); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
		return nil
	}

	leftData := data1
	rightData := data2
	if attentionSide == 1 && len(attentionData) > 0 {
		leftData = slices.Concat(data1, attentionData)
	}
	if attentionSide == 2 && len(attentionData) > 0 {
		rightData = slices.Concat(data2, attentionData)
	}

	leftWidth := maxColumnWidth(leftHeader, leftData)
	rightWidth := maxColumnWidth(rightHeader, rightData)

	diffLen := len(data1)
	if len(data2) > diffLen {
		diffLen = len(data2)
	}
	attentionLen := len(attentionData)
	maxNumber := diffLen
	if attentionLen > maxNumber {
		maxNumber = attentionLen
	}
	numberWidth := len(fmt.Sprintf("%d.", max(1, maxNumber)))

	if diffLen > 0 {
		content.WriteString("\t\t\t\t\t\tОТЛИЧИЯ\n")
		content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, "", centerText(leftHeader, leftWidth), centerText(rightHeader, rightWidth)))

		for i := 0; i < diffLen; i++ {
			var left, right string
			if i < len(data1) {
				left = data1[i]
			}
			if i < len(data2) {
				right = data2[i]
			}
			content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, fmt.Sprintf("%d.", i+1), centerText(left, leftWidth), centerText(right, rightWidth)))
		}
	}

	if attentionLen > 0 {
		if content.Len() > 0 {
			content.WriteString("\n")
		}

		content.WriteString(centerText("Обратить внимание", numberWidth+1+leftWidth+3+rightWidth))
		content.WriteString("\n")
		content.WriteString(centerText("Есть part number, но отсутствует номер сборки", numberWidth+1+leftWidth+3+rightWidth))
		content.WriteString("\n")
		content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, "", centerText(leftHeader, leftWidth), centerText(rightHeader, rightWidth)))

		for i, ref := range attentionData {
			var left, right string
			if attentionSide == 1 {
				left = ref
			} else {
				right = ref
			}
			content.WriteString(fmt.Sprintf("%-*s %s | %s\n", numberWidth, fmt.Sprintf("%d.", i+1), centerText(left, leftWidth), centerText(right, rightWidth)))
		}
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
