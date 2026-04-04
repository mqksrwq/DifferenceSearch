package main

import (
	"DifferenceSearch/utils"
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Open("MPRM160D_v7.1.bom")
	file2, _ := os.Open("выгрузка.txt")
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

	if err := utils.SaveDifferences("result.txt", data1, data2); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Результат сохранен в файл result.txt")

}
