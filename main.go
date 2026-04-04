package main

import (
	"DifferenceSearch/utils"
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Open("MPRM160D_v7.1.bom")
	file2, _ := os.Open("выгрузка.txt")
	defer file1.Close()
	defer file2.Close()

	data1 := utils.Parse(file1)
	data2 := utils.Parse(file2)

	utils.Search(&data1, &data2)
	fmt.Println(data1)
	fmt.Println(data2)

	if err := utils.SaveDifferences("differences.txt", data1, data2); err != nil {
		fmt.Println(err)
	}
}
