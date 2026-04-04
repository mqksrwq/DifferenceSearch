package utils

import (
	"fmt"
	"os"
	"strings"
)

func SaveDifferences(filename string, data1, data2 []string) error {
	content := strings.Builder{}
	content.WriteString("Уникальные значения из первого файла:\n")
	content.WriteString(strings.Join(data1, "\n"))
	content.WriteString("\n\nУникальные значения из второго файла:\n")
	content.WriteString(strings.Join(data2, "\n"))
	content.WriteString("\n")

	if err := os.WriteFile(filename, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("error writing differences file: %w", err)
	}

	return nil
}
