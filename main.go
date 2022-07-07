package main

import (
	"MyProjects/education/list"
	"fmt"
)

func main() {
	l := list.NewList()
	l.AddData(4)
	l.AddData(32)
	l.AddData(42)
	l.AddData(1)
	l.AddData(5)
	l.AddData(4)
	l.AddData(9)
	fmt.Printf("Длинна списка %d\n", l.Len())
	fmt.Println("list data:")
	l.PrintAll()
	fmt.Println()

	n, err := l.GetDataByIndex(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	fmt.Println()

	_, err = l.DeleteDataByIndex(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.PrintAll()
	fmt.Println()
	l.AddData(111)
	l.PrintAll()
	_, err = l.DeleteDataByIndex(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	l.PrintAll()
}
