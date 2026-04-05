package main

import (
	"DifferenceSearch/utils"
	"fmt"
	"os"
)

func main() {
	//if len(os.Args) != 3 {
	//	fmt.Printf("Использование: %s <файл1> <файл2>\n", filepath.Base(os.Args[0]))
	//	os.Exit(1)
	//}
	//
	//file1Path := os.Args[1]
	//file2Path := os.Args[2]

	file1Path := "MPRM160D_v7.1.bom"
	file2Path := "выгрузка.txt"

	file1, err := os.Open(file1Path)
	if err != nil {
		fmt.Printf("Не удалось открыть файл '%s': %v\n", file1Path, err)
		os.Exit(1)
	}
	file2, err := os.Open(file2Path)
	if err != nil {
		fmt.Printf("Не удалось открыть файл '%s': %v\n", file2Path, err)
		_ = file1.Close()
		os.Exit(1)
	}
	fmt.Printf("Файлы '%s' и '%s' загружены", file1.Name(), file2.Name())

	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {
			panic(err)
		}
	}(file1)
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			panic(err)
		}
	}(file2)

	data1 := utils.Parse(file1)
	data2 := utils.Parse(file2)

	utils.Search(&data1, &data2)

	if len(data1) != len(data2) {
		fmt.Printf("\nНайдено несоответствий: %d\n", len(data1)+len(data2))
	} else {
		fmt.Println("Файлы идентичны")
	}

	if err := utils.SaveDifferences("result.txt", file1.Name(), file2.Name(), data1, data2); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Результат сохранен в файл result.txt")

}
